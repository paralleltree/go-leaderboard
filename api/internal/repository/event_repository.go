package repository

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/paralleltree/go-leaderboard/internal/contract/driver"
	"github.com/paralleltree/go-leaderboard/internal/contract/repository"
	"github.com/paralleltree/go-leaderboard/internal/model"
)

const (
	eventNameKey    = "name"
	eventStartAtKey = "start_at"
	eventEndAtKey   = "end_at"
	eventListKey    = "events_ordered_by_end_at"
)

type eventRepository struct {
	idGenerator     driver.UniqueIdGenerator
	hashDriver      driver.HashDriver
	sortedSetDriver driver.SortedSetDriver
}

func NewEventRepository(
	idGenerator driver.UniqueIdGenerator,
	hashDriver driver.HashDriver,
	sortedSetDriver driver.SortedSetDriver,
) repository.EventRepository {
	return &eventRepository{
		idGenerator:     idGenerator,
		hashDriver:      hashDriver,
		sortedSetDriver: sortedSetDriver,
	}
}

func (r *eventRepository) RegisterEvent(ctx context.Context, event model.Event) (string, error) {
	fields := map[string]string{
		eventNameKey:    event.Name,
		eventStartAtKey: strconv.FormatInt(event.StartAt.Unix(), 10),
		eventEndAtKey:   strconv.FormatInt(event.EndAt.Unix(), 10),
	}
	id, err := r.idGenerator.GenerateNewId()
	if err != nil {
		return "", fmt.Errorf("generate new id: %w", err)
	}
	if err := r.hashDriver.Set(ctx, buildEventKey(id), fields); err != nil {
		return "", fmt.Errorf("set event data: %w", err)
	}
	if err := r.sortedSetDriver.SetScore(ctx, eventListKey, id, float64(event.EndAt.Unix())); err != nil {
		return "", fmt.Errorf("set event to list: %w", err)
	}
	return id, nil
}

func (r *eventRepository) GetEvent(ctx context.Context, id string) (model.Event, bool, error) {
	fields, ok, err := r.hashDriver.Get(ctx, buildEventKey(id))
	if err != nil {
		return model.Event{}, false, fmt.Errorf("get event data: %w", err)
	}
	if !ok {
		return model.Event{}, false, nil
	}

	parseTime := func(field string) (time.Time, error) {
		unixStr, ok := fields[field]
		if !ok {
			return time.Time{}, fmt.Errorf("field %s was not found", field)
		}
		unix, err := strconv.ParseInt(unixStr, 10, 64)
		if err != nil {
			return time.Time{}, fmt.Errorf("parse unix time: %w", err)
		}
		return time.Unix(unix, 0).UTC(), nil
	}

	startAt, err := parseTime(eventStartAtKey)
	if err != nil {
		return model.Event{}, true, fmt.Errorf("parse StartAt: %w", err)
	}
	endAt, err := parseTime(eventEndAtKey)
	if err != nil {
		return model.Event{}, true, fmt.Errorf("parse EndAt: %w", err)
	}
	return model.Event{
		Name:    fields[eventNameKey],
		StartAt: startAt,
		EndAt:   endAt,
	}, true, nil
}

func (r *eventRepository) GetEvents(ctx context.Context, page, count int64) ([]model.Record[model.Event], error) {
	if page < 1 {
		return nil, fmt.Errorf("page must be greater than 0")
	}
	if count < 1 {
		return nil, fmt.Errorf("count must be greater than 0")
	}
	start := (page - 1) * count
	end := page*count - 1

	ids, ok, err := r.sortedSetDriver.GetRankedList(ctx, eventListKey, start, end)
	if err != nil {
		return nil, fmt.Errorf("get event list: %w", err)
	}
	if !ok {
		return nil, nil
	}

	res := make([]model.Record[model.Event], 0, len(ids))
	for _, id := range ids {
		event, ok, err := r.GetEvent(ctx, id.Member)
		if err != nil {
			return nil, fmt.Errorf("get event: %w", err)
		}
		if !ok {
			continue
		}
		res = append(res, model.Record[model.Event]{Id: id.Member, Item: event})
	}
	return res, nil
}

func buildEventKey(id string) string {
	return fmt.Sprintf("events:%s", id)
}

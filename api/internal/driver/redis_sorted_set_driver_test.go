package driver_test

import (
	"context"
	"testing"

	driverContract "github.com/paralleltree/go-leaderboard/internal/contract/driver"
	driverImpl "github.com/paralleltree/go-leaderboard/internal/driver"
)

func TestRedisSortedSetDriver_SetScoreAndGetScore(t *testing.T) {
	endpoint, teardown := buildMockedRedis(t)
	defer teardown()
	sortedSetDriver := driverImpl.NewRedisSortedSetDriver(endpoint)

	key := "testsetkey"
	members := map[string]float64{
		"member1": 0.1,
		"member2": 10,
	}

	ctx := context.Background()

	for member, score := range members {
		if err := sortedSetDriver.SetScore(ctx, key, member, score); err != nil {
			t.Fatalf("unexpected error while setting score: %v", err)
		}
	}

	for member, wantScore := range members {
		wantOk := true
		gotScore, gotOk, err := sortedSetDriver.GetScore(ctx, key, member)
		if err != nil {
			t.Fatalf("unexpected error while getting score: %v", err)
		}
		if wantOk != gotOk {
			t.Fatalf("unexpected result(ok) with value %v: expected %v, but got %v", member, wantOk, gotOk)
		}
		if wantScore != gotScore {
			t.Fatalf("unexpected score with value %v: expected %v, but got %v", member, wantScore, gotScore)
		}
	}
}

func TestRedisSortedSetDriver_GetScoreNotExistingMember_ReturnsNothing(t *testing.T) {
	endpoint, teardown := buildMockedRedis(t)
	defer teardown()
	sortedSetDriver := driverImpl.NewRedisSortedSetDriver(endpoint)

	key := "testsetkey"
	tempMember := "temp"
	tempScore := float64(0)
	wantMember := "notexists"
	wantOk := false
	ctx := context.Background()

	if err := sortedSetDriver.SetScore(ctx, key, tempMember, tempScore); err != nil {
		t.Fatalf("unexpected error while setting score: %v", err)
	}

	_, gotOk, err := sortedSetDriver.GetScore(ctx, key, wantMember)
	if err != nil {
		t.Fatalf("unexpected error while getting score: %v", err)
	}

	if wantOk != gotOk {
		t.Fatalf("unexpected result(ok): expected %v, but got %v", wantOk, gotOk)
	}
}

func TestRedisSortedSetDriver_GetRankByDescending(t *testing.T) {
	endpoint, teardown := buildMockedRedis(t)
	defer teardown()
	sortedSetDriver := driverImpl.NewRedisSortedSetDriver(endpoint)

	key := "testsetkey"
	members := map[string]float64{
		"member1": 10,
		"member2": 20,
	}
	wantRanks := map[string]int64{
		"member1": 1,
		"member2": 0,
	}
	ctx := context.Background()

	for member, score := range members {
		if err := sortedSetDriver.SetScore(ctx, key, member, score); err != nil {
			t.Fatalf("unexpected error while setting score: %v", err)
		}
	}

	for member, wantRank := range wantRanks {
		wantOk := true
		gotRank, gotOk, err := sortedSetDriver.GetRankByDescending(ctx, key, member)
		if err != nil {
			t.Fatalf("unexpected error while getting rank: %v", err)
		}
		if wantOk != gotOk {
			t.Fatalf("unexpected result(ok): expected %v, but got %v", wantOk, gotOk)
		}
		if wantRank != gotRank {
			t.Fatalf("unexpected result: expected rank: %v, but got %v", wantRank, gotRank)
		}
	}
}

func TestRedisSortedSetDriver_GetRankByDescending_WhenMemberNotExists_ReturnsNothing(t *testing.T) {
	endpoint, teardown := buildMockedRedis(t)
	defer teardown()
	sortedSetDriver := driverImpl.NewRedisSortedSetDriver(endpoint)

	key := "testsetkey"
	members := map[string]float64{
		"member1": 10,
	}
	wantMember := "notexists"
	wantOk := false
	ctx := context.Background()

	for member, score := range members {
		if err := sortedSetDriver.SetScore(ctx, key, member, score); err != nil {
			t.Fatalf("unexpected error while setting score: %v", err)
		}
	}

	_, gotOk, err := sortedSetDriver.GetRankByDescending(ctx, key, wantMember)
	if err != nil {
		t.Fatalf("unexpected error while getting rank: %v", err)
	}
	if wantOk != gotOk {
		t.Fatalf("unexpected result(ok): expected %v, but got %v", wantOk, gotOk)
	}
}

func TestRedisSortedDriver_GetRankedList(t *testing.T) {
	endpoint, teardown := buildMockedRedis(t)
	defer teardown()
	sortedSetDriver := driverImpl.NewRedisSortedSetDriver(endpoint)

	key := "testsetkey"
	members := map[string]float64{
		"member1": 10,
		"member2": 20,
	}
	wantOk := true
	wantList := []driverContract.SortedSetItem{
		{Member: "member2", Score: 20},
		{Member: "member1", Score: 10},
	}
	ctx := context.Background()

	for member, score := range members {
		if err := sortedSetDriver.SetScore(ctx, key, member, score); err != nil {
			t.Fatalf("unexpected error while setting score: %v", err)
		}
	}

	res, gotOk, err := sortedSetDriver.GetRankedList(ctx, key, 0, int64(len(members)))
	if err != nil {
		t.Fatalf("unexpected error while getting rank: %v", err)
	}
	if wantOk != gotOk {
		t.Fatalf("unexpected result(ok): expected %v, but got %v", wantOk, gotOk)
	}

	if len(wantList) != len(res) {
		t.Fatalf("unexpected result: expected count: %v, but got %v", len(wantList), len(res))
	}
	for i, wantItem := range wantList {
		gotItem := res[i]
		if wantItem != gotItem {
			t.Fatalf("unexpected result: expected %v, but got %v", wantItem, gotItem)
		}
	}
}

func TestRedisSortedDriver_GetRankedList_WhenKeyDoesNotExist_ReturnsNothing(t *testing.T) {
	endpoint, teardown := buildMockedRedis(t)
	defer teardown()
	sortedSetDriver := driverImpl.NewRedisSortedSetDriver(endpoint)

	key := "testsetkey"
	wantOk := false
	ctx := context.Background()

	_, gotOk, err := sortedSetDriver.GetRankedList(ctx, key, 0, -1)

	if err != nil {
		t.Fatalf("unexpected error while getting rank: %v", err)
	}
	if wantOk != gotOk {
		t.Fatalf("unexpected result(ok): expected %v, but got %v", wantOk, gotOk)
	}
}

package driver_test

import (
	"context"
	"testing"

	"github.com/paralleltree/go-leaderboard/internal/driver"
)

func TestRedisHashDriver_SetAndGet_ReturnsStoredFields(t *testing.T) {
	endpoint, teardown := buildMockedRedis(t)
	defer teardown()
	hashDriver := driver.NewRedisHashDriver(endpoint)

	key := "testhashkey"
	fields := map[string]string{
		"field1": "value1",
		"field2": "value2",
	}
	wantOk := true

	ctx := context.Background()

	if err := hashDriver.Set(ctx, key, fields); err != nil {
		t.Fatalf("unexpected error while setting value: %v", err)
	}

	gotFields, gotOk, err := hashDriver.Get(ctx, key)
	if err != nil {
		t.Fatalf("unexpected error while getting value: %v", err)
	}

	if wantOk != gotOk {
		t.Fatalf("unexpected result(ok): expected %v, but got %v", wantOk, gotOk)
	}

	assertMapEquals(fields, gotFields)
}

func TestRedisHashDriver_GetNotExistingKey_ReturnsNothing(t *testing.T) {
	endpoint, teardown := buildMockedRedis(t)
	defer teardown()
	hashDriver := driver.NewRedisHashDriver(endpoint)

	key := "testhashkey"
	wantOk := false

	ctx := context.Background()

	_, gotOk, err := hashDriver.Get(ctx, key)

	if err != nil {
		t.Fatalf("unexpected error while getting value: %v", err)
	}

	if wantOk != gotOk {
		t.Fatalf("unexpected result(ok): expected %v, but got %v", wantOk, gotOk)
	}
}

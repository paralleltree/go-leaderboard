package driver_test

import (
	"testing"

	"github.com/alicebob/miniredis/v2"
)

func buildMockedRedis(t *testing.T) (string, func()) {
	server, err := miniredis.Run()
	if err != nil {
		t.Fatalf("unexpected error while starting test redis server: %v", err)
	}
	return server.Addr(), func() { server.Close() }
}

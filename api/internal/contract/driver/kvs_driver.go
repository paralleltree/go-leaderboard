package driver

import "context"

type HashDriver interface {
	Get(ctx context.Context, key string) (map[string]string, bool, error)
	Set(ctx context.Context, key string, fields map[string]string) error
}

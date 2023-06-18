package repository

import (
	"context"

	"github.com/bonaysoft/amiicli/internal/entity"
)

type Device interface {
	Restore(ctx context.Context, amiibo *entity.Amiibo, password string) error
	Simulate(ctx context.Context, amiibo *entity.Amiibo) error
}

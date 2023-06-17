package repository

import (
	"context"

	"github.com/bonaysoft/amiicli/internal/entity"
)

type Device interface {
	Clone(ctx context.Context, amiibo *entity.Amiibo) error
	Simulate(ctx context.Context, amiibo *entity.Amiibo) error
}

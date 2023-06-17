package repository

import (
	"context"

	"github.com/bonaysoft/amiicli/internal/entity"
)

type AmiiboSelectOptions struct {
	Name string
}

type AmiiboSelectOption interface {
	apply(*AmiiboSelectOptions)
}

type nameOption string

func (o nameOption) apply(opt *AmiiboSelectOptions) {
	opt.Name = string(o)
}

func AmiiboSelectWithName(name string) AmiiboSelectOption {
	return nameOption(name)
}

type Amiibo interface {
	List(ctx context.Context) ([]entity.Amiibo, error)
	Select(ctx context.Context, mode entity.Mode, opts ...AmiiboSelectOption) (*entity.Amiibo, error)
}

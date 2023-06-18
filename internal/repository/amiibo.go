package repository

import (
	"context"

	"github.com/bonaysoft/amiicli/internal/entity"
)

type AmiiboListOptions struct {
	Name       string
	SearchPath string
}

func NewAmiiboListOptions(opts ...ListOption) *AmiiboListOptions {
	options := new(AmiiboListOptions)
	for _, opt := range opts {
		opt.apply(options)
	}

	return options
}

type ListOption interface {
	apply(*AmiiboListOptions)
}

type nameOption string

func (o nameOption) apply(opt *AmiiboListOptions) {
	opt.Name = string(o)
}

func AmiiboListWithName(name string) ListOption {
	return nameOption(name)
}

type searchDirOption string

func (o searchDirOption) apply(opt *AmiiboListOptions) {
	opt.SearchPath = string(o)
}

func AmiiboListWithSearchDir(dir string) ListOption {
	return searchDirOption(dir)
}

type AmiiboSearcher interface {
	List(ctx context.Context, opts ...ListOption) ([]entity.Amiibo, error)
}

type AmiiboConstructor func() AmiiboSearcher

var supportAmiiboSearchers = map[string]AmiiboConstructor{
	"remote": NewAmiiboRemote,
	"local":  NewAmiiboLocal,
}

func NewAmiiboSearcher(name string) AmiiboSearcher {
	return supportAmiiboSearchers[name]()
}

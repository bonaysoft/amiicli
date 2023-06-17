package repository

import (
	"context"
	"fmt"
	"io/fs"
	"math/rand"
	"os"
	"path/filepath"
	"strings"

	"github.com/bonaysoft/amiicli/internal/entity"
	"github.com/samber/lo"
)

var _ Amiibo = NewAmiiboLocal()

type AmiiboLocal struct {
	dirPath string
}

func NewAmiiboLocal() *AmiiboLocal {
	return &AmiiboLocal{dirPath: "/Users/saltbo/Develop/bogit/amiibogo/build"}
}

func (a *AmiiboLocal) List(ctx context.Context) ([]entity.Amiibo, error) {
	return a.loadAll()
}

func (a *AmiiboLocal) Select(ctx context.Context, mode entity.Mode, opts ...AmiiboSelectOption) (*entity.Amiibo, error) {
	amiibos, err := a.loadAll()
	if err != nil {
		return nil, err
	}

	var cfg AmiiboSelectOptions
	for _, opt := range opts {
		opt.apply(&cfg)
	}

	switch mode {
	case entity.ModeRandom:
		return &amiibos[rand.Intn(len(amiibos))], nil
	case entity.ModeSpecify:
		amiibo, ok := lo.Find(amiibos, func(item entity.Amiibo) bool { return item.Name == cfg.Name })
		if !ok {
			return nil, fmt.Errorf("not found the specified amiibo: %s", cfg.Name)
		}
		return &amiibo, nil
	default:
		return nil, fmt.Errorf("unsupported mode: %s", mode)
	}
}

func (a *AmiiboLocal) loadAll() ([]entity.Amiibo, error) {
	amiibos := make([]entity.Amiibo, 0)
	walker := func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		name := strings.TrimSuffix(d.Name(), filepath.Ext(d.Name()))
		if name == "" {
			return nil
		}

		amiibos = append(amiibos, entity.Amiibo{Name: name, Path: filepath.Join(a.dirPath, path)})
		return nil
	}

	return amiibos, fs.WalkDir(os.DirFS(a.dirPath), ".", walker)
}

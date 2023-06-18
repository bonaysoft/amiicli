package repository

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/bonaysoft/amiicli/internal/entity"
)

var _ AmiiboSearcher = (*AmiiboLocal)(nil)

type AmiiboLocal struct {
}

func NewAmiiboLocal() AmiiboSearcher {
	return &AmiiboLocal{}
}

func (a *AmiiboLocal) List(ctx context.Context, opts ...ListOption) ([]entity.Amiibo, error) {
	options := NewAmiiboListOptions(opts...)
	amiibos := make([]entity.Amiibo, 0)
	walker := func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		name := strings.TrimSuffix(d.Name(), filepath.Ext(d.Name()))
		if name == "" {
			return nil
		}

		if !strings.Contains(name, options.Name) {
			return nil
		}

		amiibos = append(amiibos, entity.Amiibo{Name: name, Path: filepath.Join(options.SearchPath, path)})
		return nil
	}

	return amiibos, fs.WalkDir(os.DirFS(options.SearchPath), ".", walker)
}

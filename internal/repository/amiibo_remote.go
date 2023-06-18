package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/bonaysoft/amiicli/internal/entity"
	"github.com/samber/lo"
)

const (
	remoteAmiiboDatabase = "https://raw.githubusercontent.com/N3evin/AmiiboAPI/master/database/amiibo.json"
)

var _ AmiiboSearcher = (*AmiiboRemote)(nil)

type AmiiboRemote struct {
}

func NewAmiiboRemote() AmiiboSearcher {
	return &AmiiboRemote{}
}

func (a *AmiiboRemote) List(ctx context.Context, opts ...ListOption) ([]entity.Amiibo, error) {
	options := NewAmiiboListOptions(opts...)
	resp, err := http.Get(remoteAmiiboDatabase)
	if err != nil {
		return nil, fmt.Errorf("error getting remote amiibo database: %v", err)
	}

	var adb entity.AmiiboDatabase
	if err := json.NewDecoder(resp.Body).Decode(&adb); err != nil {
		return nil, fmt.Errorf("error decoding remote amiibo database: %s", err)
	}

	return lo.Filter(lo.MapToSlice(adb.Amiibos, func(key string, value entity.Amiibo) entity.Amiibo {
		value.ID = key
		return value
	}), func(item entity.Amiibo, index int) bool {
		if options.Name == "" {
			return true
		}

		return strings.Contains(item.Name, options.Name)
	}), nil
}

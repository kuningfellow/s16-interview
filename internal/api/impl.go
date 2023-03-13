package api

import (
	"context"
	"errors"
	"fmt"

	"github.com/kuningfellow/s16-interview/internal/omdb"
)

type apiImpl struct {
	omdbClient omdb.OMDB
}

func (a *apiImpl) Search(ctx context.Context, req SearchRequest) (*SearchResponse, error) {
	resp, err := a.omdbClient.Search(ctx, req)
	return resp, processOMDBError(err)
}

func (a *apiImpl) Detail(ctx context.Context, req DetailRequest) (*DetailResponse, error) {
	resp, err := a.omdbClient.Detail(ctx, req)
	return resp, processOMDBError(err)
}

var _ API = (*apiImpl)(nil) // interface satisfaction check

func NewAPIImpl(omdbClient omdb.OMDB) *apiImpl {
	return &apiImpl{
		omdbClient: omdbClient,
	}
}

func processOMDBError(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, omdb.ErrorNotFound) {
		return fmt.Errorf("%w: %v", ErrorMovieNotFound, err)
	} else if errors.Is(err, omdb.ErrorServerError) {
		return fmt.Errorf("%w: %v", ErrorExternal, err)
	}
	return fmt.Errorf("%w: %v", ErrorInternal, err)
}

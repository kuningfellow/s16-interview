package api

import (
	"context"
	"errors"
	"sync"

	"github.com/kuningfellow/s16-interview/internal/omdb"
)

var (
	ErrorNotInitialized = errors.New("not initialized")

	ErrorMovieNotFound = errors.New("movie found")

	// ErrorExternal is errors on other server
	ErrorExternal = errors.New("external error")
	// ErrorInternal is errors on our server
	ErrorInternal = errors.New("internal error")
)

// For now we only have OMDB backend so we just reuse the struct definitions.
type (
	SearchRequest  = omdb.SearchRequest
	SearchResponse = omdb.SearchResponse

	DetailRequest  = omdb.DetailRequest
	DetailResponse = omdb.DetailResponse
)

// API is our program's API facade.
type API interface {
	Search(context.Context, SearchRequest) (*SearchResponse, error)
	Detail(context.Context, DetailRequest) (*DetailResponse, error)
}

var (
	globalAPIInstance     API
	globalAPIInstanceLock sync.RWMutex
)

func GetGlobalAPIInstance() (API, error) {
	globalAPIInstanceLock.RLock()
	defer globalAPIInstanceLock.RUnlock()
	if globalAPIInstance == nil {
		return nil, ErrorNotInitialized
	}
	return globalAPIInstance, nil
}

func SetGlobalAPIInstance(instance API) {
	globalAPIInstanceLock.Lock()
	defer globalAPIInstanceLock.Unlock()
	globalAPIInstance = instance
}

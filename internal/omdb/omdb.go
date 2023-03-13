package omdb

import (
	"context"
	"errors"
)

var (
	ErrorUnauthorized  = errors.New("unauthorized")
	ErrorNotFound      = errors.New("not found")
	ErrorServerError   = errors.New("server error")
	ErrorEmptyResponse = errors.New("empty response")
	ErrorInternal      = errors.New("internal error")
)

// OMDB is an OMDB client
type OMDB interface {
	Search(context.Context, SearchRequest) (*SearchResponse, error)
	Detail(context.Context, DetailRequest) (*DetailResponse, error)
}

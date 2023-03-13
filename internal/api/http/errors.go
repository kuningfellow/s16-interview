package api_http

import (
	"errors"
	"net/http"

	"github.com/kuningfellow/s16-interview/internal/api"
)

func ToHttpError(err error) (code int, message string) {
	if errors.Is(err, api.ErrorNotInitialized) {
		return http.StatusServiceUnavailable, "server not ready"
	}
	if errors.Is(err, api.ErrorMovieNotFound) {
		return http.StatusNotFound, "movie not found"
	}
	if errors.Is(err, api.ErrorInternal) {
		return http.StatusInternalServerError, "internal server error"
	}
	if errors.Is(err, api.ErrorExternal) {
		return http.StatusInternalServerError, "external server error"
	}

	return http.StatusOK, ""
}

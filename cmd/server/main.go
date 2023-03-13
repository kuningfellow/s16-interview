package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"

	api "github.com/kuningfellow/s16-interview/internal/api"
	api_gin "github.com/kuningfellow/s16-interview/internal/api/http/gin"
	omdb_http "github.com/kuningfellow/s16-interview/internal/omdb/http"
)

var (
	port    = flag.Int("port", 8080, "port to serve")
	omdbURL = flag.String("omdb_url", "", "URL of OMDB server")
	omdbKey = flag.String("omdb_key", "", "API key of OMDB server")
)

func run(stopChan <-chan os.Signal) error {
	r := gin.Default()

	omdbClient, err := omdb_http.NewHTTPOMDB(*omdbURL, *omdbKey)
	if err != nil {
		return err
	}

	apiInstance := api.NewAPIImpl(omdbClient)

	api_gin.AddAPI(apiInstance, r)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", *port),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	<-stopChan
	ctx, done := context.WithTimeout(context.Background(), 2*time.Second)
	defer done()
	err = srv.Shutdown(ctx)
	return err
}

func main() {
	flag.Parse()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill)

	err := run(quit)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

package api_gin

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kuningfellow/s16-interview/internal/api"
	api_http "github.com/kuningfellow/s16-interview/internal/api/http"
)

type ginAPI struct {
	instance api.API
}

func AddAPI(instance api.API, g *gin.Engine) error {
	a := &ginAPI{
		instance: instance,
	}
	var err error
	if err = a.RegisterSearch(g); err != nil {
		return fmt.Errorf("error register Search API: %w", err)
	}
	if err = a.RegisterDetail(g); err != nil {
		return fmt.Errorf("error register Detail API: %w", err)
	}
	return nil
}

func (a *ginAPI) RegisterSearch(g *gin.Engine) error {
	g.GET("/search", func(c *gin.Context) {
		query := c.Query("query")

		resp, err := a.instance.Search(c.Request.Context(), api.SearchRequest{
			Query: query,
		})
		if err != nil {
			log.Default().Print(err)
			code, msg := api_http.ToHttpError(err)
			c.JSON(code, msg)
			return
		}

		c.JSON(http.StatusOK, resp)
	})
	return nil
}

func (a *ginAPI) RegisterDetail(g *gin.Engine) error {
	g.GET("/detail/:id", func(c *gin.Context) {
		id := c.Param("id")

		resp, err := a.instance.Detail(c.Request.Context(), api.DetailRequest{
			ID: id,
		})
		if err != nil {
			log.Default().Print(err)
			code, msg := api_http.ToHttpError(err)
			c.JSON(code, msg)
			return
		}

		c.JSON(http.StatusOK, resp)
	})
	return nil
}

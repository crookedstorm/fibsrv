// Package api generates a basic web server with gin middleware and router to
// fibsrv
package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/crookedstorm/fibsrv/pkg/fibber"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

var fibs fibber.MainSequence

// InitRouter sets up the router for the https service. It adds in the the gin
// middleware for a logger and recovery as well as an external prometheus module.
func InitRouter() *gin.Engine {
	fibs.InitCache()
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	p := ginprometheus.NewPrometheus("fibsrv")
	// Reduce the cardinality of the path part of the URL for Promtheus
	p.ReqCntURLLabelMappingFn = func(c *gin.Context) string {
		url := c.Request.URL.String()
		for _, p := range c.Params {
			if p.Key == "num" {
				url = strings.Replace(url, p.Value, ":num", 1)
				break
			}
		}
		return url
	}
	p.Use(router)

	// Under 94, we get extremely fast lookups. Over the 94th number, there
	// is still some benefit to having precalculated the start of the sequence
	router.GET("/api/fibonacci/:num", func(c *gin.Context) {
		n := c.Param("num")
		fNum, _ := strconv.Atoi(n)
		if fNum > 94 {
			ans, _ := fibs.BigFib(fNum)
			c.String(http.StatusOK, ans.String())
		} else {
			ans, _ := fibs.LookupFib(fNum)
			c.String(http.StatusOK, fmt.Sprintf("%d", ans))
		}
	})

	return router
}

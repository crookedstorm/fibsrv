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

func InitRouter() *gin.Engine {
	fibs.InitCache()
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	p := ginprometheus.NewPrometheus("fibsrv")
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

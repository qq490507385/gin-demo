package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const v1Api = "/v1"

func load(g *gin.Engine) *gin.Engine {
	// Middlewares
	g.Use(gin.Recovery())

	// 404
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	adminRouter(v1Api, g)

	return g
}

// InitRoutes 注册路由
func InitRoutes(g *gin.Engine) *gin.Engine {
	return load(g)
}

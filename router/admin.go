package router

import (
	"doc-system/app/controllers/teams"
	"doc-system/app/controllers/users"

	"github.com/gin-gonic/gin"
)

func adminRouter(v string, g *gin.Engine) {
	u := g.Group(v)
	{
		u.POST("/login", users.Login)
		u.POST("/user/create", users.Create)

		u.GET("/teams/list", teams.List)
	}
}

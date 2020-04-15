package users

import (
	"doc-system/app/exceptions"

	"github.com/gin-gonic/gin"
)

// Login 登录
func Login(c *gin.Context) {
	var login LoginRequest
	if err := c.Bind(&login); err != nil {
		exceptions.SendResponse(c, exceptions.ErrBind, nil)
		return
	}
	
}

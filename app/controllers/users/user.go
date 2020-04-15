package users

import (
	exc "doc-system/app/exceptions"
	"doc-system/app/models"
	request "doc-system/app/requests"
	"doc-system/util"

	"github.com/gin-gonic/gin"
)

// Create 新增用户
func Create(c *gin.Context) {
	var r request.UserCreateRequest
	if err := c.Bind(&r); err != nil {
		exc.SendResponse(c, exc.ErrBind, nil)
		return
	}
	password, _ := util.EncryptPassword(r.Password)

	u := models.UserModel{
		Name:     r.Name,
		Sex:      r.Sex,
		Username: r.Username,
		Password: password,
	}

	// Validte
	if msg, err := r.ValidateCreate(); err != nil {
		exc.SendValidResponse(c, msg)
		return
	}

	if err := u.Create(); err != nil {
		exc.SendResponse(c, exc.ErrDatabase, nil)
		return
	}

	exc.SendResponse(c, exc.OK, nil)
}

package teams

import (
	"doc-system/app/exceptions"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	exceptions.SendResponse(c, exceptions.ErrValidation, nil)
	// exceptions.SendResponse(c, exceptions.OK, nil)
}

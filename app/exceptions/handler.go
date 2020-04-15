package exceptions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Exception 异常
type Exception struct {
	Code    int
	Message string
}

// Response 返回类型
type Response struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ValidResponse 表单返回类型
type ValidResponse struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

func (err Exception) Error() string {
	return err.Message
}

// RenderError 解析异常
func renderError(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}
	switch typed := err.(type) {
	case *Exception:
		return typed.Code, typed.Message
	default:
	}
	return ServerError.Code, err.Error()
}

// SendResponse 请求返回解析
func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := renderError(err)
	var status string
	if code == 0 {
		status = "ok"
	} else {
		status = "error"
	}
	c.JSON(http.StatusOK, Response{
		Status:  status,
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// SendValidResponse 表单验证消息返回
func SendValidResponse(c *gin.Context, errs interface{}) {
	code, message := renderError(ErrValidation)

	c.JSON(http.StatusOK, ValidResponse{
		Status:  "error",
		Code:    code,
		Message: message,
		Errors:  errs,
	})
}

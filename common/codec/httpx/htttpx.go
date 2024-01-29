package httpx

import (
	"net/http"

	"github.com/binbin6363/icuc/common/err"
	"github.com/gin-gonic/gin"
)

// Response 统一的 http 返回格式
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func newResponse(data interface{}, e error) *Response {
	return &Response{
		Code:    err.Code(e),
		Message: err.Msg(e),
		Data:    data,
	}
}

// SendResponse 回复json数据包
func SendResponse(c *gin.Context, data interface{}, err error) {
	if c == nil {
		return
	}

	rsp := newResponse(data, err)
	c.JSON(http.StatusOK, rsp)
}

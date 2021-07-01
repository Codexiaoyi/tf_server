package response

import (
	"net/http"
	"tfserver/internal/errmsg"

	"github.com/gin-gonic/gin"
)

//服务端响应
func Response(c *gin.Context, status int) {
	var context = gin.H{
		"status":  status,
		"message": errmsg.GetErrMsg(status),
	}
	c.JSON(
		http.StatusOK,
		context,
	)
}

//服务端带信息响应
func ResponseWithData(c *gin.Context, status int, data map[string]interface{}) {
	var context = gin.H{
		"status":  status,
		"message": errmsg.GetErrMsg(status),
	}
	for key, value := range data {
		context[key] = value
	}
	c.JSON(
		http.StatusOK,
		context,
	)
}

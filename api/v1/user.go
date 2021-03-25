package v1

import (
	"net/http"
	"tfserver/model"
	"tfserver/repository"
	"tfserver/util/errmsg"

	"github.com/gin-gonic/gin"
)

// 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)

	result := repository.AddUser(&data)

	if result {
		c.JSON(
			http.StatusOK, gin.H{
				"status": errmsg.SUCCESS,
			},
		)
	} else {
		c.JSON(
			http.StatusOK, gin.H{
				"status": errmsg.ERROR,
			},
		)
	}
}

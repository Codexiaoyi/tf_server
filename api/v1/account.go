package v1

import (
	"tfserver/model"
	"tfserver/repository"
	"tfserver/util/errmsg"
	"tfserver/util/jwt"
	"tfserver/util/response"

	"github.com/gin-gonic/gin"
)

//注册接口
func Register(c *gin.Context) {
	var account model.Account
	_ = c.ShouldBindJSON(&account)

	status := errmsg.ERROR
	isExist := repository.CheckAccount(account.Email)
	if !isExist {
		//账号不存在，添加新账号
		repository.AddAccount(&account)
		//添加默认账户信息
		repository.AddUser(&model.User{
			Email: account.Email,
		})
		status = errmsg.SUCCESS
	} else {
		status = errmsg.ERROR_ACCOUNT_EXIST
	}

	response.Response(c, status)
}

//登录接口
func Login(c *gin.Context) {
	var account model.Account
	_ = c.ShouldBindJSON(&account)

	status := errmsg.ERROR
	isExist := repository.CheckAccount(account.Email)
	if !isExist {
		//账号不存在
		status = errmsg.ERROR_ACCOUNT_NOT_EXIST
		response.Response(c, status)
		return
	}

	password, err := repository.QueryPasswordByEmail(account.Email)
	if err != nil || password != account.Password {
		//密码错误
		status = errmsg.ERROR_PASSWORD_ERROR
	}

	if password == account.Password {
		//成功登录颁发token
		data := make(map[string]interface{})
		token, err := jwt.GetToken(account.Email)
		if err == nil {
			data["token"] = token
			status = errmsg.SUCCESS
			response.ResponseWithData(c, status, data)
			return
		}
	}

	response.Response(c, status)
}

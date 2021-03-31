package v1

import (
	"tfserver/model"
	"tfserver/repository"
	"tfserver/util/errmsg"
	"tfserver/util/jwt"
	"tfserver/util/response"

	"github.com/gin-gonic/gin"
	"gopkg.in/hlandau/passlib.v1"
)

//注册接口
func Register(c *gin.Context) {
	var account model.Account
	_ = c.ShouldBindJSON(&account)

	status := errmsg.ERROR
	isExist := repository.CheckAccount(account.Email)
	if !isExist {
		hash, err := passlib.Hash(account.Password)
		if err == nil {
			//hash成功
			account.Password = hash
		}
		//账号不存在，添加新账号
		repoErr := repository.AddAccount(&account)
		if repoErr == nil {
			status = errmsg.SUCCESS
		}
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

	newAccount, err := repository.QueryAccountByEmail(account.Email)
	_, hashErr := passlib.Verify(account.Password, newAccount.Password)
	if err != nil || hashErr != nil {
		//密码错误
		status = errmsg.ERROR_PASSWORD_ERROR
	} else {
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

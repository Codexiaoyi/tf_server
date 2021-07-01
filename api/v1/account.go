package v1

import (
	"tfserver/internal/errmsg"
	"tfserver/internal/model"
	"tfserver/internal/repository"
	"tfserver/internal/response"
	"tfserver/pkg/jwt"

	"github.com/gin-gonic/gin"
	"gopkg.in/hlandau/passlib.v1"
)

//注册接口
func Register(c *gin.Context) {
	var account model.Account
	_ = c.ShouldBindJSON(&account)

	isExist := repository.CheckAccount(account.Email)
	if isExist {
		response.Response(c, errmsg.ERROR_ACCOUNT_EXIST)
		return
	}

	hash, err := passlib.Hash(account.Password)
	if err != nil || hash == "" {
		response.Response(c, errmsg.ERROR)
		return
	}

	//hash成功
	account.Password = hash
	//添加新账号
	repoErr := repository.AddAccount(&account)
	if repoErr != nil {
		response.Response(c, errmsg.ERROR)
		return
	}

	//成功注册颁发token
	data := make(map[string]interface{})
	token, err := jwt.GetToken(account.Email)
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}
	data["token"] = token
	response.ResponseWithData(c, errmsg.SUCCESS, data)
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

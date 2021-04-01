package v1

import (
	"tfserver/application/command"
	"tfserver/application/query"
	"tfserver/model"
	"tfserver/repository"
	"tfserver/util/errmsg"
	"tfserver/util/response"

	"github.com/Codexiaoyi/go-mapper"
	"github.com/gin-gonic/gin"
)

//创建新团队
func CreateNewTeam(c *gin.Context) {
	var command command.CreateTeam
	_ = c.ShouldBindJSON(&command)
	status := errmsg.ERROR

	var team model.Team
	err := mapper.StructMapByFieldName(&command, &team)
	if err == nil {
		err = repository.CreateNewTeam(&team)
		if err == nil {
			status = errmsg.SUCCESS
		}
	}

	response.Response(c, status)
}

//更新团队信息
func UpdateTeamInfo(c *gin.Context) {
	var command command.UpdateTeamInfo
	_ = c.ShouldBindJSON(&command)
	status := errmsg.ERROR

	if team, err := repository.QueryTeamInfoByTeamId(command.TeamId); err == nil && team.ID > 0 {
		//团队存在
		err = mapper.StructMapByFieldName(&command, &team)
		if err == nil {
			err = repository.UpdateTeamInfo(&team)
			if err == nil {
				status = errmsg.SUCCESS
			}
		}
	} else {
		status = errmsg.ERROR_TEAM_NOT_EXIST
	}

	response.Response(c, status)
}

//获取团队信息
func GetTeamInfo(c *gin.Context) {
	var query query.GetTeamInfo
	_ = c.ShouldBindJSON(&query)

	team, err := repository.QueryTeamInfoByTeamId(query.TeamId)
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}

	if team.ID <= 0 {
		response.Response(c, errmsg.ERROR_TEAM_NOT_EXIST)
		return
	}

	data := make(map[string]interface{})
	data["team"] = team
	response.ResponseWithData(c, errmsg.SUCCESS, data)
}

//获取团队所有成员
func GetTeamMembers(c *gin.Context) {
	var query query.GetTeamMembers
	_ = c.ShouldBindJSON(&query)

	isExist := repository.CheckTeam(query.TeamId)
	if !isExist {
		response.Response(c, errmsg.ERROR_TEAM_NOT_EXIST)
		return
	}

	members, err := repository.QueryMembersByTeamId(query.TeamId)
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}

	data := make(map[string]interface{})
	data["members"] = members
	response.ResponseWithData(c, errmsg.SUCCESS, data)
}

//获取用户所有团队
func GetUserTeams(c *gin.Context) {
	var query query.GetUserTeams
	_ = c.ShouldBindJSON(&query)

	teams, err := repository.QueryTeamsByEmail(query.Email)
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}

	data := make(map[string]interface{})
	data["teams"] = teams
	response.ResponseWithData(c, errmsg.SUCCESS, data)
}

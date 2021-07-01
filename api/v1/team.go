package v1

import (
	"tfserver/internal/application/command"
	"tfserver/internal/application/query"
	"tfserver/internal/errmsg"
	"tfserver/internal/model"
	"tfserver/internal/repository"
	"tfserver/internal/response"

	"github.com/Codexiaoyi/go-mapper"
	"github.com/gin-gonic/gin"
)

//创建新团队
func CreateNewTeam(c *gin.Context) {
	var command command.CreateTeam
	_ = c.ShouldBindJSON(&command)
	status := errmsg.ERROR
	email := c.GetString("email")

	var team model.Team
	err := mapper.StructMapByFieldName(&command, &team)
	if err == nil {
		team.Email = email
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
	email := c.GetString("email")

	teams, err := repository.QueryTeamsByEmail(email)
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}

	data := make(map[string]interface{})
	data["teams"] = teams
	response.ResponseWithData(c, errmsg.SUCCESS, data)
}

//拉成员进入团队
func AddMember(c *gin.Context) {
	var command command.AddMember
	_ = c.ShouldBindJSON(&command)

	isExist := repository.CheckAccount(command.Email)
	if !isExist {
		response.Response(c, errmsg.ERROR_USER_NOT_EXIST)
		return
	}

	member, err := repository.GetMemberByTeamIdAndEmail(command.Email, command.TeamId)
	if err == nil && member.ID > 0 {
		//已经在团队内
		response.Response(c, errmsg.ERROR_MEMBER_ALREADY_IN_TEAM)
		return
	}

	newMember := model.Member{
		TeamId:   command.TeamId,
		Email:    command.Email,
		IsLeader: false,
	}

	err = repository.AddMember(&newMember)
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}

	response.Response(c, errmsg.SUCCESS)
}

//成员退出团队
func MemberLeave(c *gin.Context) {
	var command command.MemberLeave
	_ = c.ShouldBindJSON(&command)
	email := c.GetString("email")

	member, err := repository.GetMemberByTeamIdAndEmail(email, command.TeamId)
	if err != nil || member.ID <= 0 {
		//团队查无此人
		response.Response(c, errmsg.ERROR_MEMBER_NOT_IN_TEAM)
		return
	}

	if member.IsLeader {
		//队长不能直接退出，走转让后退出或者解散
		response.Response(c, errmsg.ERROR_MEMBER_IS_LEADER)
		return
	}

	err = repository.RemoveMember(&member)
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}

	response.Response(c, errmsg.SUCCESS)
}

//队长踢除成员
func RemoveMember(c *gin.Context) {
	var command command.RemoveMember
	_ = c.ShouldBindJSON(&command)
	email := c.GetString("email")
	if email == command.Email {
		//不能对自己操作
		response.Response(c, errmsg.ERROR_MEMBER_IS_ME)
		return
	}

	member, err := repository.GetMemberByTeamIdAndEmail(email, command.TeamId)
	if err != nil || !member.IsLeader {
		//不是队长
		response.Response(c, errmsg.ERROR_MEMBER_IS_NOT_LEADER)
		return
	}

	member, err = repository.GetMemberByTeamIdAndEmail(command.Email, command.TeamId)
	if err != nil || member.ID <= 0 {
		//团队查无此人
		response.Response(c, errmsg.ERROR_MEMBER_NOT_IN_TEAM)
		return
	}

	err = repository.RemoveMember(&member)
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}

	response.Response(c, errmsg.SUCCESS)
}

//队长转让
func TransformLeader(c *gin.Context) {
	var command command.TransformLeader
	_ = c.ShouldBindJSON(&command)
	email := c.GetString("email")
	oldLeader, err := repository.GetMemberByTeamIdAndEmail(email, command.TeamId)
	if err != nil || !oldLeader.IsLeader {
		//不是队长
		response.Response(c, errmsg.ERROR_MEMBER_IS_NOT_LEADER)
		return
	}

	newLeader, err := repository.GetMemberByTeamIdAndEmail(command.Email, command.TeamId)
	if err != nil || newLeader.ID <= 0 {
		//团队查无此人
		response.Response(c, errmsg.ERROR_MEMBER_NOT_IN_TEAM)
		return
	}

	err = repository.TransformLeader(&oldLeader, &newLeader)
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}

	response.Response(c, errmsg.SUCCESS)
}

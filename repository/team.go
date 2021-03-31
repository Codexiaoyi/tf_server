package repository

import (
	"tfserver/model"

	"gorm.io/gorm"
)

//查询用户拥有的团队
func QueryTeamsByEmail(email string) ([]model.Member, error) {
	members := make([]model.Member, 0)
	err := db.Where("email = ?", email).Find(&members).Error
	return members, err
}

//获取团队所有成员
func QueryMembersByTeamId(teamId int) ([]model.Member, error) {
	members := make([]model.Member, 0)
	err := db.Where("team_id = ?", teamId).Find(&members).Error
	return members, err
}

//获取团队信息
func QueryTeamInfoByTeamId(teamId int) (model.Team, error) {
	var team model.Team
	err := db.Limit(1).Where("ID = ?", teamId).Find(&team).Error
	return team, err
}

//创建团队
func CreateNewTeam(team *model.Team) error {
	return BeginTransaction(db, func(tx *gorm.DB) error {
		var err error
		err = tx.Create(&team).Error
		if err != nil {
			return err
		}
		//添加创建人为默认成员
		err = tx.Create(&model.Member{
			TeamId:   int(team.ID),
			Email:    team.Create,
			IsLeader: true,
		}).Error
		return err
	})
}

//添加成员

//更新团队信息

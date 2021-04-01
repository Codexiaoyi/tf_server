package repository

import (
	"tfserver/model"

	"gorm.io/gorm"
)

//查询团队是否存在
func CheckTeam(teamId int) bool {
	var team model.Team
	db.Select("id").Where("ID = ?", teamId).First(&team)
	return team.ID > 0
}

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
			Email:    team.Email,
			IsLeader: true,
		}).Error
		return err
	})
}

//更新团队信息
func UpdateTeamInfo(team *model.Team) error {
	return BeginTransaction(db, func(tx *gorm.DB) error {
		return tx.Model(&team).Where("ID = ?", team.ID).Updates(map[string]interface{}{
			"team_name":    team.TeamName,
			"avatar":       team.Avatar,
			"introduction": team.Introduction,
		}).Error
	})
}

//获取某团队的某成员
func GetMemberByTeamIdAndEmail(email string, teamId int) (model.Member, error) {
	var member model.Member
	err := db.Limit(1).Where("team_id = ? and email = ?", teamId, email).Find(&member).Error
	return member, err
}

//添加成员进入团队
func AddMember(member *model.Member) error {
	return BeginTransaction(db, func(tx *gorm.DB) error {
		return tx.Create(&member).Error
	})
}

//将成员移除团队
func RemoveMember(member *model.Member) error {
	return BeginTransaction(db, func(tx *gorm.DB) error {
		return tx.Delete(&member).Error
	})
}

//队长转让
func TransformLeader(oldLeader, newLeader *model.Member) error {
	return BeginTransaction(db, func(tx *gorm.DB) error {
		var err error
		err = tx.Model(&oldLeader).Where("ID = ?", oldLeader.ID).Updates(map[string]interface{}{
			"is_leader": false,
		}).Error
		if err != nil {
			return err
		}
		err = tx.Model(&oldLeader).Where("ID = ?", newLeader.ID).Updates(map[string]interface{}{
			"is_leader": true,
		}).Error
		return err
	})
}

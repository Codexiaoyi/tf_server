package command

type CreateTeam struct {
	TeamName     string
	Introduction string
}

type UpdateTeamInfo struct {
	TeamId       int
	TeamName     string
	Introduction string
}

type AddMember struct {
	TeamId int
	Email  string
}

type MemberLeave struct {
	TeamId int
}

type RemoveMember struct {
	TeamId int
	Email  string
}

type TransformLeader struct {
	TeamId int
	Email  string
}

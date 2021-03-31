package query

type GetUserTeams struct {
	Email string
}

type GetTeamMembers struct {
	TeamId int
}

type GetTeamInfo struct {
	TeamId int
}

package command

type CreateTeam struct {
	TeamName     string
	Email        string
	Introduction string
}

type UpdateTeamInfo struct {
	TeamId       int
	TeamName     string
	Introduction string
}

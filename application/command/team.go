package command

type CreateTeam struct {
	TeamName     string
	Email        string
	Introduction string
}

type UpdateTeamInfo struct {
	TeamName     string
	Introduction string
}

package flyweight

import "time"

type Team struct {
	ID             uint64
	Name           int
	Shield         []byte
	Players        []Player
	HistorycalData []HistorycalData
}

type Player struct {
	Name         string
	Surname      string
	PreviousTeam uint64
	Photo        []byte
}

type HistorycalData struct {
	Year          uint8
	LeagueResults []Match
}

type Match struct {
	Date          time.Time
	VisitorID     uint64
	LocalID       uint64
	LocalStore    byte
	VisitorScore  byte
	LocalShoots   uint16
	VisitorShoots uint16
}

type teamFlyweightFactory struct {
	createdTeams map[int]*Team
}

const (
	TEAM_A = iota
	TEAM_B
)

func (t *teamFlyweightFactory) GetTeam(teamID int) *Team {
	if t.createdTeams[teamID] != nil {
		return t.createdTeams[teamID]
	}

	team := getTeamFactory(teamID)
	t.createdTeams[teamID] = &team
	return t.createdTeams[teamID]
}

func (t *teamFlyweightFactory) GetNumberOfObjects() int {
	return len(t.createdTeams)
}

func getTeamFactory(team int) Team {
	switch team {
	case TEAM_B:
		return Team{
			ID:   2,
			Name: TEAM_B,
		}
	default:
		return Team{
			ID:   1,
			Name: TEAM_A,
		}
	}
}

func NewTeamFactory() teamFlyweightFactory {
	return teamFlyweightFactory{
		createdTeams: make(map[int]*Team),
	}
}

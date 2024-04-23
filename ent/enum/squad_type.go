package enum

import (
	"slices"
)

type SquadType string

const (
	MAIN            SquadType = "main"
	MAIN_GOALKEEPER SquadType = "main_goalkeeper"
	RESERVE         SquadType = "reserve"
	STATE           SquadType = "state"
)

func getSquadTypes() []SquadType {
	return []SquadType{MAIN, MAIN_GOALKEEPER, RESERVE, STATE}
}

func (SquadType) Values() (kinds []string) {
	for _, value := range getSquadTypes() {
		kinds = append(kinds, string(value))
	}
	return
}

func (s SquadType) Value() string {
	return string(s)
}

func (s SquadType) IsValid() bool {
	return slices.Contains(getSquadTypes(), s)
}

func (s SquadType) IsMain() bool {
	return s == MAIN
}

func (s SquadType) IsMainGoalkeeper() bool {
	return s == MAIN_GOALKEEPER
}

func (s SquadType) IsReserve() bool {
	return s == RESERVE
}

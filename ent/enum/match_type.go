package enum

import "slices"

type MatchType string

const (
	FRIENDLY MatchType = "friendly"
	LEAGUE   MatchType = "league"
	GROUP    MatchType = "group"
	ROUND_64 MatchType = "round-64"
	ROUND_32 MatchType = "round-32"
	ROUND_16 MatchType = "round-16"
	QUARTER  MatchType = "quarter"
	SEMI     MatchType = "semi"
	FINAL    MatchType = "final"
)

func getMatchType() []MatchType {
	return []MatchType{FRIENDLY, LEAGUE, GROUP, ROUND_64, ROUND_32, ROUND_16, QUARTER, SEMI, FINAL}
}

func (MatchType) Values() (kinds []string) {
	for _, value := range getMatchType() {
		kinds = append(kinds, string(value))
	}
	return
}

func (m MatchType) Value() string {
	return string(m)
}

func (m MatchType) IsValid() bool {
	return slices.Contains(getMatchType(), m)
}

func (m MatchType) IsOnlyMainTime() bool {
	return m == FRIENDLY || m == LEAGUE || m == GROUP
}

func (m MatchType) HasExtraTime() bool {
	return m == ROUND_64 || m == ROUND_32 || m == ROUND_16 || m == QUARTER || m == SEMI || m == FINAL
}

package enum

import (
	"math/rand"
	"slices"
)

type MainPosition string

const (
	GKP MainPosition = "goalkeeper"
	DEF MainPosition = "defender"
	MID MainPosition = "midfielder"
	ATT MainPosition = "attacker"
	ALL MainPosition = "all"
)

type Position string

const (
	GK  Position = "GK"
	CB  Position = "CB"
	LB  Position = "LB"
	RB  Position = "RB"
	DM  Position = "DM"
	CM  Position = "CM"
	LM  Position = "LM"
	RM  Position = "RM"
	CAM Position = "CAM"
	LW  Position = "LW"
	RW  Position = "RW"
	ST  Position = "ST"
)

func getPositions() []Position {
	return []Position{GK, CB, LB, RB, DM, CM, LM, RM, CAM, LW, RW, ST}
}

func getAttackPositions() []Position {
	return []Position{LW, RW, ST}
}

func getMidfieldPositions() []Position {
	return []Position{DM, CM, LM, RM, CAM}
}

func getDefencePositions() []Position {
	return []Position{CB, LB, RB}
}

func (Position) Values() (kinds []string) {
	for _, value := range getPositions() {
		kinds = append(kinds, string(value))
	}
	return
}

func (p Position) Value() string {
	return string(p)
}

func (p Position) Definition() string {
	switch p {
	case GK:
		return "goalkeeper"
	case CB:
		return "center-back"
	case LB:
		return "left-back"
	case RB:
		return "right-back"
	case DM:
		return "defensive-midfielder"
	case CM:
		return "central-midfielder"
	case LM:
		return "left-midfielder"
	case RM:
		return "right-midfielder"
	case CAM:
		return "central-attacking-midfielder"
	case LW:
		return "left-winger"
	case RW:
		return "right-winger"
	case ST:
		return "striker"
	default:
		return ""
	}
}

func (p Position) IsValid() bool {
	return slices.Contains(getPositions(), p)
}

func (p Position) IsGK() bool {
	return p == GK
}

func (p Position) IsDefender() bool {
	return p == CB || p == LB || p == RB
}

func (p Position) IsMidfielder() bool {
	return p == DM || p == CM || p == LM || p == RM || p == CAM
}

func (p Position) IsAttacker() bool {
	return p == LW || p == RW || p == ST
}

func GetRandomPositions(mainPosition MainPosition) []Position {
	positionCount := 1
	if rand.Intn(6) == 2 {
		positionCount = 2
	}

	switch mainPosition {
	case ATT:
		return getCountedPositions(getAttackPositions(), positionCount)
	case MID:
		return getCountedPositions(getMidfieldPositions(), positionCount)
	case DEF:
		return getCountedPositions(getDefencePositions(), positionCount)
	case GKP:
		return []Position{GK}
	default:
		return []Position{getPositions()[rand.Intn(len(getPositions()))]}
	}
}

func getCountedPositions(positions []Position, count int) []Position {
	mapCount := make(map[int]interface{})
	randomPositions := make([]Position, 0)

	for i := 0; i < count; i++ {
		random := rand.Intn(len(positions))
		if _, ok := mapCount[random]; !ok {
			randomPositions = append(randomPositions, positions[random])
			mapCount[random] = nil
		}
	}

	return randomPositions
}

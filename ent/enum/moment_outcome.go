package enum

import (
	"math/rand"
	"slices"
)

type MomentOutcome string

const (
	GOAL    MomentOutcome = "goal"
	MISS    MomentOutcome = "miss"
	OFFSIDE MomentOutcome = "offside"
	SAVE    MomentOutcome = "save"
)

func getMomentOutcome() []MomentOutcome {
	return []MomentOutcome{GOAL, MISS, OFFSIDE, SAVE}
}

func (MomentOutcome) Values() (kinds []string) {
	for _, value := range getMomentOutcome() {
		kinds = append(kinds, string(value))
	}
	return
}

func (m MomentOutcome) Value() string {
	return string(m)
}

func (m MomentOutcome) IsValid() bool {
	return slices.Contains(getMomentOutcome(), m)
}

func (m MomentOutcome) GetContent() string {
	random := rand.Intn(3)
	switch m {
	case GOAL:
		switch random {
		case 0:
			return "⚽️Net bulger! Crowd roars!"
		case 1:
			return "⚽️Upset! Underdog wonder goal!"
		default:
			return "⚽️Teamwork! Clinical finish!"
		}
	case MISS:
		switch random {
		case 0:
			return "🚫Close shave! Inches away!"
		case 1:
			return "🚫Agony! Inches wide!"
		default:
			return "🚫Off the post! Frustration!"
		}
	case OFFSIDE:
		switch random {
		case 0:
			return "🚩Flag up! Offside trap works!"
		case 1:
			return "🚩Offside! Linesman spot on!"
		default:
			return "🚩Offside! VAR check incoming? 🚩VAR❓"
		}
	case SAVE:
		switch random {
		case 0:
			return "🧤Safe hands! Keeper denies!"
		case 1:
			return "🧤Brilliant stop! Keeper shines!"
		default:
			return "🧤Reflex save! Keeper alert!"
		}
	default:
		return ""
	}
}

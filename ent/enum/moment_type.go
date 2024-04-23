package enum

import (
	"math/rand"
	"slices"
)

type MomentType string

const (
	ATTACK      MomentType = "attack"
	FREE_KICK   MomentType = "free-kick"
	CORNER_KICK MomentType = "corner-kick"
	THROW_IN    MomentType = "throw-in"
	PENALTY     MomentType = "penalty"
	FOUL        MomentType = "foul"
)

func getMomentType() []MomentType {
	return []MomentType{ATTACK, FREE_KICK, CORNER_KICK, THROW_IN, PENALTY, FOUL}
}

func (MomentType) Values() (kinds []string) {
	for _, value := range getMomentType() {
		kinds = append(kinds, string(value))
	}
	return
}

func (m MomentType) Value() string {
	return string(m)
}

func (m MomentType) IsValid() bool {
	return slices.Contains(getMomentType(), m)
}

func (m MomentType) HasOutcome() bool {
	return m == ATTACK || m == FREE_KICK || m == CORNER_KICK || m == PENALTY
}

func (m MomentType) IsFoul() bool {
	return m == FOUL
}

func GetRandomMoment() MomentType {
	return getMomentType()[rand.Intn(len(getMomentType()))]
}

func (m MomentType) GetContent() string {
	random := rand.Intn(3)
	switch m {
	case ATTACK:
		switch random {
		case 0:
			return "👟Screamer: A powerful shot from long distance that finds the net."
		case 1:
			return "👟Skillful dribbling, weaving through defenders.\n"
		default:
			return "👟Cross delivered, a striker lurks in the box."
		}
	case FREE_KICK:
		switch random {
		case 0:
			return "🎯Curling effort, over the wall, into the top corner."
		case 1:
			return "🎯Direct shot on goal? Or a tactical pass to build?"
		default:
			return "🎯Wall set up, defenders ready to block."
		}
	case CORNER_KICK:
		switch random {
		case 0:
			return "🚩Keeper commands the area, aerial battle ensues.\n"
		case 1:
			return "🚩Whipped in, seeking a head or foot to finish."
		default:
			return "🚩Out-swinger, a scramble in the box."
		}
	case THROW_IN:
		switch random {
		case 0:
			return "🔄Long throw, a weapon in the attacking third."
		case 1:
			return "🔄Short throw, quick feet, a chance to cross."
		default:
			return "🔄Battle for the ball, midfield duel ignites."
		}
	case PENALTY:
		switch random {
		case 0:
			return "⚖️Pressure on the taker, keeper eyes the shooter."
		case 1:
			return "⚖️Calm and composed, a chance to convert."
		default:
			return "⚖️Mind games, a stuttering run-up."
		}
	case FOUL:
		switch random {
		case 0:
			return "❌Reckless challenge, a card could be coming."
		case 1:
			return "❌Tactical foul, a professional foul to break play."
		default:
			return "❌Whistle blows, tempers flare."
		}
	default:
		return ""
	}
}

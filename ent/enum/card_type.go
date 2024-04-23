package enum

import (
	"math/rand"
	"slices"
)

type CardType string

const (
	YELLOW CardType = "yellow"
	RED    CardType = "red"
)

func getCardType() []CardType {
	return []CardType{YELLOW, RED}
}

func (CardType) Values() (kinds []string) {
	for _, value := range getCardType() {
		kinds = append(kinds, string(value))
	}
	return
}

func (c CardType) Value() string {
	return string(c)
}

func (c CardType) IsValid() bool {
	return slices.Contains(getCardType(), c)
}

func (c CardType) GetContent() string {
	random := rand.Intn(3)
	switch c {
	case YELLOW:
		switch random {
		case 0:
			return "🟨Caution! One foot in trouble."
		case 1:
			return "🟨Tactical foul? Risky gamble."
		default:
			return "🟨Frustration flares, card shown."
		}
	case RED:
		switch random {
		case 0:
			return "🟥Sent off! Game on edge."
		case 1:
			return "🟥Early shower, dreams dashed."
		default:
			return "🟥Straight red! VAR check incoming? 🟥VAR❓"
		}
	default:
		return ""
	}
}

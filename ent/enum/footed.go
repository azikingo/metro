package enum

import (
	"math/rand"
	"slices"
)

type Footed string

const (
	L    Footed = "left"
	R    Footed = "right"
	BOTH Footed = "both"
)

func getFooted() []Footed {
	return []Footed{L, R, BOTH}
}

func (Footed) Values() (kinds []string) {
	for _, value := range getFooted() {
		kinds = append(kinds, string(value))
	}
	return
}

func (footed Footed) Value() string {
	return string(footed)
}

func (footed Footed) IsValid() bool {
	return slices.Contains(getFooted(), footed)
}

func (footed Footed) IsLeft() bool {
	return footed == L
}

func (footed Footed) IsRight() bool {
	return footed == R
}

func (footed Footed) IsBoth() bool {
	return footed == BOTH
}

func GetRandomFooted() Footed {
	random := rand.Intn(100)
	if random < 85 {
		return R
	} else if random < 95 {
		return L
	} else {
		return BOTH
	}
}

package enum

import (
	"math/rand"
	"slices"
)

type WeatherType string

const (
	SUNNY    WeatherType = "sunny"
	CLOUDY   WeatherType = "cloudy"
	OVERCAST WeatherType = "overcast"
	RAINY    WeatherType = "rainy"
	FOGGY    WeatherType = "foggy"
	SNOWING  WeatherType = "snowing"
)

func getWeatherType() []WeatherType {
	return []WeatherType{SUNNY, CLOUDY, OVERCAST, RAINY, FOGGY, SNOWING}
}

func (WeatherType) Values() (kinds []string) {
	for _, value := range getWeatherType() {
		kinds = append(kinds, string(value))
	}
	return
}

func (w WeatherType) Value() string {
	return string(w)
}

func (w WeatherType) IsValid() bool {
	return slices.Contains(getWeatherType(), w)
}

func GetRandomWeather() WeatherType {
	return getWeatherType()[rand.Intn(len(getWeatherType()))]
}

func (w WeatherType) GetContent() string {
	switch w {
	case SUNNY:
		if rand.Intn(2) == 0 {
			return "☀️Scorching sun, fast-paced action!"
		} else {
			return "🌞Sun's got the players fired up, brace for scorching shots!"
		}
	case CLOUDY:
		if rand.Intn(2) == 0 {
			return "🌥️Cloudy skies, but the action's heating up!"
		} else {
			return "⛅️Don't let the clouds fool you, goals waiting to break through! ⚡"
		}
	case OVERCAST:
		if rand.Intn(2) == 0 {
			return "‍☁️️Grit and passion under the gloomy sky!"
		} else {
			return "️‍☁️Overcast skies, but the players are ready to shine!"
		}
	case RAINY:
		if rand.Intn(2) == 0 {
			return "⛈️️Rain can't dampen the heat on the pitch, get ready for a splash!"
		} else {
			return "️🌧Downpour of goals, slippery skills!"
		}
	case FOGGY:
		if rand.Intn(2) == 0 {
			return "️🌫️Every touch a surprise, tension thick as the fog!"
		} else {
			return "️🌁Foggy weather, but the players are ready to break through!"
		}
	case SNOWING:
		if rand.Intn(2) == 0 {
			return "️❄️Winter wonderland, fierce battle for dominance!"
		} else {
			return "️❄️Snowing, but the goals are still flowing!"
		}
	default:
		return "Unpredictable weather, unstoppable players!"
	}
}

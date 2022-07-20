package model

import (
	"math/rand"
	"time"
)

type WeatherStatus struct {
	Status Status `json:"status"`
}

type CompiledWeatherStatus struct {
	Status         Status `json:"status"`
	StatusCompiled string `json:"status_compiled"`
}

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func RandomValueStatus() (status WeatherStatus) {
	rand.Seed(time.Now().UnixNano())
	minSafe := 0
	maxSafe := 6
	minStandBy := 6
	maxStandBy := 10
	minDanger := 10
	maxDanger := 100
	min := 1
	max := 3
	switch rand.Intn(max-min+1) + min {
	case 1:
		status.Status.Water = rand.Intn(maxSafe-minSafe+1) + minSafe
		status.Status.Wind = rand.Intn(maxSafe-minSafe+1) + minSafe
	case 2:
		status.Status.Water = rand.Intn(maxStandBy-minStandBy+1) + minStandBy
		status.Status.Wind = rand.Intn(maxStandBy-minStandBy+1) + minStandBy
	default:
		status.Status.Water = rand.Intn(maxDanger-minDanger+1) + minDanger
		status.Status.Wind = rand.Intn(maxDanger-minDanger+1) + minDanger
	}
	return status
}

func (s *Status) CheckStatus() string {
	overallStatus := ""
	if s.Water <= 5 && s.Wind <= 6 {
		overallStatus = "SAFE"
	} else if s.Water <= 8 && s.Wind <= 15 {
		overallStatus = "STANDBY"
	} else {
		overallStatus = "DANGER"
	}
	return overallStatus
}

package model

import (
	"math/rand"
	"time"
)

type WeatherStatus struct {
	Status Status `json:"status"`
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

func (s *Status) CheckStatus() (string, string) {
	statusWater := ""
	statusWind := ""
	switch wtr := s.Water; {
	case wtr <= 5:
		statusWater = "aman"
	case wtr <= 8:
		statusWater = "siaga"
	case wtr > 8:
		statusWater = "bahaya"
	default:
		statusWater = "bahaya"
	}

	switch wind := s.Wind; {
	case wind <= 6:
		statusWind = "aman"
	case wind <= 15:
		statusWind = "siaga"
	case wind > 15:
		statusWind = "bahaya"
	default:
		statusWind = "bahaya"
	}
	return statusWater, statusWind
}

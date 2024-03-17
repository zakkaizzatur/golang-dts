package controllers

import (
	"assignment-3/models"
	"math/rand"
)

var currentWater, currentWind int

func GetStatus() models.Status {

	return models.Status{
		Water: currentWater,
		Wind: currentWind,
	}

}

func UpdateStatus() {
	currentWater = rand.Intn(100) + 1
	currentWind = rand.Intn(100) + 1
}



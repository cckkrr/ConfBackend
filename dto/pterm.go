package dto

import (
	"github.com/golang-module/carbon/v2"
)

type PTermCalcedCoordDTO struct {
	PTermId    string               `json:"termid"`
	UpdateTime carbon.DateTimeMilli `json:"updateTime"`
	X          float64              `json:"x"`
	Y          float64              `json:"y"`
	Z          float64              `json:"z"`
}

package constants

import "quasar-backend/model"

var Satellites map[string]model.Point

func InitDefaultSatellites() {
	Satellites = make(map[string]model.Point)
	Satellites["kenobi"] = model.Point{X: -500, Y: -200}
	Satellites["skywalker"] = model.Point{X: 100, Y: -100}
	Satellites["sato"] = model.Point{X: 500, Y: 100}
}
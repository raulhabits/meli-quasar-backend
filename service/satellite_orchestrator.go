package service

import "quasar-backend/model"

var Satellites = make(map[string]model.SatelliteIncomingMessage)

func FillSatelliteInformation(request model.SatelliteIncomingMessage) {
	Satellites[request.Name] = request
}

func RetrieveSavedSatellites() model.TopSecretInputMessages{
	var satellites []model.SatelliteIncomingMessage
	for _, satellite := range Satellites {
		satellites = append(satellites, satellite)
	}
	return model.TopSecretInputMessages{Satellites: satellites}
}
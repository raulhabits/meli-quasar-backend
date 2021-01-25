package model

type TopSecretInputMessages struct {
	Satellites []SatelliteIncomingMessage `json:"satellites"`
}

type SuccessOutputMessage struct {
	Position Point  `json:"position"`
	Message  string `json:"message"`
}

func (entity *TopSecretInputMessages) ExtractMessages() [][]string {
	var messages [][]string
	for _, satellite := range entity.Satellites {
		messages = append(
			messages,
			satellite.Messages,
		)
	}
	return messages
}

func (entity *TopSecretInputMessages) ExtractDistances() []float64 {
	var distances []float64
	for _, satellite := range entity.Satellites {
		distances = append(distances, satellite.Distance)
	}
	return distances
}
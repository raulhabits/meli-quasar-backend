package model

type SatelliteIncomingMessage struct {
	Name string  `json:"name"`
	Message
}

type Message struct {
	Distance float64  `json:"distance"`
	Messages []string  `json:"messages"`
}

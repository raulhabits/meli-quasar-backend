package service

import "quasar-backend/model"

func ValidateAndExtractMessages(request model.TopSecretInputMessages) (model.SuccessOutputMessage, error) {

	var response model.SuccessOutputMessage

	x, y, err := GetLocation(request.ExtractDistances())
	if err != nil {
		response.Message = err.Error()
	} else {
		message, err := GetMessage(request.ExtractMessages())
		if err != nil {
			response.Message = err.Error()
		} else {
			response.Message=  message
			response.Position= model.Point{X:x, Y: y}
		}
	}
	return response, err
}
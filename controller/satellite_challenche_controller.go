package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"quasar-backend/model"
	"quasar-backend/service"
)

func InitSatelliteChallengeRoutes(router *mux.Router) {
	router.HandleFunc("/topsecret", PostTopSecretInputMessages).Methods("POST")
	router.HandleFunc("/topsecret_split/{satellite_name}", PostSatelliteMessageByName).Methods("POST")
	router.HandleFunc("/topsecret_split", GetMessageFromLoadedSatellites).Methods("GET")
	router.HandleFunc("/topsecret_split", ClearLoadedSatellites).Methods("DELETE")
}

 func PostTopSecretInputMessages(w http.ResponseWriter, req *http.Request) {
	 w.Header().Set("Content-Type", "application/json")
	 reqBody, _ := ioutil.ReadAll(req.Body)
	 var messages model.TopSecretInputMessages
	 json.Unmarshal(reqBody, &messages)
	 response, error := service.ValidateAndExtractMessages(messages)
	 if error != nil {
	 	w.WriteHeader(404)
	 }
	 json.NewEncoder(w).Encode(response)
}

func PostSatelliteMessageByName(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(req.Body)
	var message model.SatelliteIncomingMessage
	json.Unmarshal(reqBody, &message)

	vars := mux.Vars(req)
	satelliteName := vars["satellite_name"]
	message.Name = satelliteName
	service.FillSatelliteInformation(message)
	w.WriteHeader(201)
}

func GetMessageFromLoadedSatellites(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	messages := service.RetrieveSavedSatellites()
	response, err := service.ValidateAndExtractMessages(messages)
	if err != nil {
		w.WriteHeader(404)
	}
	json.NewEncoder(w).Encode(response)
}

func ClearLoadedSatellites(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	service.ClearSatellites()
	w.WriteHeader(200)
}

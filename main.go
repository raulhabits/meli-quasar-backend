package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"quasar-backend/constants"
	"quasar-backend/controller"
	"quasar-backend/service"
)

func main() {
	constants.InitDefaultSatellites()

	message, _ := service.GetMessage([][]string{[]string{"este","", "", "mensaje", ""}, []string{"","es", "", "", "secreto"}, []string{"este","", "un", "", ""}})
	println(message)

	x, y, err := service.GetLocation([]float64{1004.987562112089, 400, 2})

	log.Println(x, " ", y, " ", err)

	myRouter := mux.NewRouter()
	controller.InitSatelliteChallengeRoutes(myRouter)

	log.Println("Listing for requests at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
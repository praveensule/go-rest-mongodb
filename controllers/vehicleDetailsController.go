package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"go-rest-mongodb/models"
	. "go-rest-mongodb/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

var vehicleDetailsRepository VehicleDetailsRepository

var GetAllVehicleDetails = func(w http.ResponseWriter, r * http.Request) {
	vehicleDetails, err := vehicleDetailsRepository.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	log.Info(vehicleDetails);
	respondWithJson(w, http.StatusOK, vehicleDetails)
}

var GetVehicleDetailsById = func(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintln(w, "Not implemented!")
}

var CreateVehicleDetail = func(w http.ResponseWriter, r * http.Request) {
	defer r.Body.Close()
	var vehicleDetails models.VehicleDetails
	if err := json.NewDecoder(r.Body).Decode(&vehicleDetails); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	log.Info(vehicleDetails);
	insertResult, err := vehicleDetailsRepository.Insert(vehicleDetails);
	if  err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	vehicleDetails.ID = insertResult.(primitive.ObjectID)
	respondWithJson(w, http.StatusCreated, vehicleDetails)
}

var UpdateVehicleDetail = func(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintln(w, "Not implemented")
}

var DeleteVehicleDetail = func(w http.ResponseWriter, r * http.Request) {
	//Vars returns the route variables for the current request, if any.
	vars := mux.Vars(r)
	//Get id from the current request
	id := vars["id"]
	fmt.Println(id)
	if err := vehicleDetailsRepository.Delete(id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

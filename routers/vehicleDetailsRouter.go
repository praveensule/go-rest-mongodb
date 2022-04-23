package routers

import (
	"github.com/gorilla/mux"
	"go-rest-mongodb/controllers"
)

func AddPlacesRouter(r *mux.Router) *mux.Router {
	s := r.PathPrefix("/vehicleRegistrationNumber").Subrouter()
	s.HandleFunc("", controllers.GetAllVehicleDetails).Methods("GET")
	s.HandleFunc("/{id}", controllers.GetVehicleDetailsById).Methods("GET")
	s.HandleFunc("", controllers.CreateVehicleDetail).Methods("POST")
	s.HandleFunc("", controllers.UpdateVehicleDetail).Methods("PUT")
	s.HandleFunc("/{id}", controllers.DeleteVehicleDetail).Methods("DELETE")
	return s;
}

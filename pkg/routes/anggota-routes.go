package routes

import (
	"learn_gorilla/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterAnggotaRoutes = func(router *mux.Router) {
	router.HandleFunc("/anggota/", controllers.CreateAnggota).Methods("POST")
	router.HandleFunc("/anggota/", controllers.GetAnggota).Methods("GET")
	router.HandleFunc("/anggota/{anggotaId}", controllers.GetAnggotaById).Methods("GET")
	router.HandleFunc("/anggota/{anggotaId}", controllers.UpdateAnggotaById).Methods("PUT")
	router.HandleFunc("/anggota/{anggotaId}", controllers.DeleteAnggota).Methods("DELETE")
}

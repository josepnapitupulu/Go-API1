package routes

import (
	"github.com/gorilla/mux"
	"github.com/josepnapitupulu/API_Baptis/controllers"
)

var RegisterBaptissRoutes = func(router *mux.Router) {
	// router.HandleFunc("/", controllers.Index).Methods("GET")
	// router.HandleFunc("/jemaatBaru", controllers.Create).Methods("POST")
	router.HandleFunc("/baptis/", controllers.CreateBaptis).Methods("POST")
	router.HandleFunc("/baptis/", controllers.GetBaptis).Methods("GET")
	router.HandleFunc("/baptis/{baptisId}", controllers.GetBaptisById).Methods("GET")
	router.HandleFunc("/baptis/{baptisId}", controllers.UpdateBaptis).Methods("PUT")
	router.HandleFunc("/baptis/{baptisId}", controllers.DeleteBaptis).Methods("DELETE")
}

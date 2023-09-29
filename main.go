package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/josepnapitupulu/API_Baptis/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBaptissRoutes(r)
	http.Handle("/", r)

	fmt.Print("Starting Server localhost:7073")
	log.Fatal(http.ListenAndServe("localhost:7073", r))
}
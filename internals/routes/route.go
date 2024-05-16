package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"main.go/internals/controller"
)

func InitRoute() {
	port := 8081
	router := mux.NewRouter()

	router.HandleFunc("/", controller.TestRoute).Methods("GET")

	log.Println("Application running on port http://localhost", port)
	log.Fatal(http.ListenAndServe(":8081", router))

}

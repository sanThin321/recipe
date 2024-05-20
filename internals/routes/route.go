package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"main.go/internals/controller"
)

func InitRoute() {
	port := 8080
	router := mux.NewRouter()

	// Define your API routes first
	router.HandleFunc("/recipe", controller.AddRecipe).Methods("POST")
	// router.HandleFunc("/recipes", controller.GetAllRecipe).Methods("GET")
	// router.HandleFunc("/", controller.TestRoute).Methods("GET")

	// Serve static files from ./view directory
	fhandler := http.FileServer(http.Dir("./view"))
	router.PathPrefix("/").Handler(fhandler)

	log.Println("Application running on port", port)
	log.Fatal(http.ListenAndServe(":8080", router))
}

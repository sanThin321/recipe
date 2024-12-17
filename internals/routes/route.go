package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"main.go/internals/controller"
	// "main.go/internals/controller"
)

func InitRoute() {
	port := 8080
	router := mux.NewRouter()

	// router.HandleFunc("/", controller.TestRoute).Methods("GET")
	
	// user route
	router.HandleFunc("/user_register", controller.Register_user).Methods("POST")
	router.HandleFunc("/user_login", controller.User_login).Methods("POST")
	router.HandleFunc("/logout", controller.Logout)
	router.HandleFunc("/get_user", controller.Get_user).Methods("POST")
	router.HandleFunc("/update_user/{user_id}", controller.Update_user).Methods("PUT")
	router.HandleFunc("/del_user/{user_id}", controller.Delete_user).Methods("DELETE")
	//recipe
	router.HandleFunc("/recipe", controller.AddRecipe).Methods("POST")
	router.HandleFunc("/recipe/{rid}", controller.GetRecipe).Methods("GET")
	router.HandleFunc("/recipe/{rid}", controller.UpdateRecipe).Methods("PUT")
	router.HandleFunc("/recipe/{rid}", controller.DeleteRecipe).Methods("DELETE")
	router.HandleFunc("/recipes", controller.GetAllRecipe)

	fhandler := http.FileServer(http.Dir("./view"))
	router.PathPrefix("/").Handler(fhandler)

	log.Println("Application running on port http://localhost", port)
	log.Fatal(http.ListenAndServe(":8080", router))

}

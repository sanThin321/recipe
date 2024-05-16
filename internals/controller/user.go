package controller

import (
	"fmt"
	"net/http"

	postgres "main.go/data"
	"main.go/internals/model"
)

func TestRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Controller Pass")

	res := model.TestModel()

	postgres.Init()

	if res {
		w.Write([]byte("All Passed"))
	}
}

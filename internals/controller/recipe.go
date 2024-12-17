package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"main.go/internals/model"
	httpresponse "main.go/util/httprespones"
)

func AddRecipe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling AddRecipe request")

	err := r.ParseMultipartForm(10 << 20) // Limit file size to 10 MB
	if err != nil {
		httpresponse.RespondWithError(w, http.StatusBadRequest, "Invalid form data")
		return
	}

	// Get form values
	recipeName := r.FormValue("recipeName")
	ingredient := r.FormValue("ingredients")
	steps := r.FormValue("steps")
	file, _, err := r.FormFile("image")
	if err != nil {
		httpresponse.RespondWithError(w, http.StatusBadRequest, "Error reading image")
		return
	}
	defer file.Close()

	// Read the image data
	imgData, err := io.ReadAll(file)
	if err != nil {
		httpresponse.RespondWithError(w, http.StatusInternalServerError, "Error reading image data")
		return
	}

	// Create the Recipe object
	recipe := model.Recipe{
		RecipeName: recipeName,
		Ingredient: ingredient,
		Steps:      steps,
		Image:      imgData,
	}

	// Save the recipe to the database
	if err := recipe.Create(); err != nil {
		httpresponse.RespondWithError(w, http.StatusInternalServerError, err.Error())
		// fmt.Println("error")
		return
	}

	response := map[string]string{"status": "Recipe added"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func GetAllRecipe(w http.ResponseWriter, r *http.Request) {
	recipes, err := model.GetAllRecipe()
	if err != nil {
		httpresponse.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpresponse.RespondWithJSON(w, http.StatusOK, recipes)
}
func GetRecipe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ridStr := vars["rid"]
	rid, err := strconv.ParseInt(ridStr, 10, 64)
	if err != nil {
		httpresponse.RespondWithError(w, http.StatusBadRequest, "Invalid recipe ID")
		return
	}

	recipe, err := model.GetRecipeByID(rid)
	if err != nil {
		httpresponse.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httpresponse.RespondWithJSON(w, http.StatusOK, recipe)
}
func UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling UpdateRecipe request")
	vars := mux.Vars(r)
	ridStr := vars["rid"]
	rid, err := strconv.ParseInt(ridStr, 10, 64)
	if err != nil {
		httpresponse.RespondWithError(w, http.StatusBadRequest, "Invalid recipe ID")
		return
	}

	err = r.ParseMultipartForm(10 << 20)
	if err != nil {
		httpresponse.RespondWithError(w, http.StatusBadRequest, "Invalid form data")
		return
	}

	recipeName := r.FormValue("recipeName")
	ingredient := r.FormValue("ingredients")
	steps := r.FormValue("steps")
	useExistingImage := r.FormValue("useExistingImage") == "true"

	var imgData []byte
	if !useExistingImage {
		file, _, err := r.FormFile("image")
		if err != nil && err != http.ErrMissingFile {
			httpresponse.RespondWithError(w, http.StatusBadRequest, "Error reading image")
			return
		}
		if file != nil {
			defer file.Close()
			imgData, err = io.ReadAll(file)
			if err != nil {
				httpresponse.RespondWithError(w, http.StatusInternalServerError, "Error reading image data")
				return
			}
		}
	}
	// file, _, err := r.FormFile("image")
	// if err != nil && err != http.ErrMissingFile {
	// 	httpresponse.RespondWithError(w, http.StatusBadRequest, "Error reading image")
	// 	return
	// }
	// defer file.Close()

	// var imgData []byte
	// if err == nil {
	// 	imgData, err = io.ReadAll(file)
	// 	if err != nil {
	// 		httpresponse.RespondWithError(w, http.StatusInternalServerError, "Error reading image data")
	// 		return
	// 	}
	// }

	recipe := model.Recipe{
		Rid:        rid,
		RecipeName: recipeName,
		Ingredient: ingredient,
		Steps:      steps,
	}
	if !useExistingImage {
		recipe.Image = imgData
	} else {
		// Fetch existing image from the database
		existingRecipe, err := model.GetRecipeByID(rid)
		if err != nil {
			httpresponse.RespondWithError(w, http.StatusInternalServerError, "Error fetching existing recipe")
			return
		}
		recipe.Image = existingRecipe.Image
	}

	if err := recipe.Update(); err != nil {
		httpresponse.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httpresponse.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "success"})
}
func DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling Delete Recipe request")
	vars := mux.Vars(r)
	ridStr := vars["rid"]
	rid, err := strconv.ParseInt(ridStr, 10, 64)
	if err != nil {
		httpresponse.RespondWithError(w, http.StatusBadRequest, "Invalid recipe ID")
		return
	}

	recipe := model.Recipe{Rid: rid}
	if err := recipe.Delete(); err != nil {
		httpresponse.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := map[string]string{"status": "success"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
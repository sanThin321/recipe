package controller

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
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

// func GetAllRecipe(w http.ResponseWriter, r *http.Request) {
// 	recipes, err := model.GetAllRecipe()
// 	if err != nil {
// 		httpresponse.RespondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	httpresponse.RespondWithJSON(w, http.StatusOK, recipes)
// }

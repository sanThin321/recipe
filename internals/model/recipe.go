package model

import (
	"errors"
	"fmt"
	// "fmt"
	// "strconv"

	// "fmt"
	// "fmt"

	postgres "main.go/data"
)

type Recipe struct {
	Rid        int64  `json:"rid"`
	RecipeName string `json:"recipename"`
	Ingredient string `json:"ingredient"`
	Steps      string `json:"steps"`
	Image      []byte `json:"image"`
}

const queryGetAllRecipe="SELECT rid, recipename, ingredient, steps, image FROM recipe;"

func (r *Recipe) Create() error {
    if postgres.Db == nil {
        return errors.New("database connection is not initialized")
    }

    query := "INSERT INTO recipe (recipename, ingredient, steps, image) VALUES ($1, $2, $3, $4)"
    _, err := postgres.Db.Exec(query, r.RecipeName, r.Ingredient, r.Steps, r.Image)
    return err
}

func GetAllRecipe() ([]Recipe, error) {
    rows, err := postgres.Db.Query(queryGetAllRecipe)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    // Get column names
    columns, err := rows.Columns()
    if err != nil {
        return nil, err
    }

    // Check if the 'ingredient' column exists
    var ingredientExists bool
    for _, col := range columns {
        if col == "ingredient" {
            ingredientExists = true
            break
        }
    }
    if !ingredientExists {
        return nil, errors.New("ingredient column does not exist")
    }

    var recipes []Recipe
    for rows.Next() {
        var rec Recipe
        if err := rows.Scan(&rec.Rid, &rec.RecipeName, &rec.Ingredient, &rec.Steps, &rec.Image); err != nil {
            return nil, err
        }
        recipes = append(recipes, rec)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }
    return recipes, nil
}

const queryGetRecipeByID = "SELECT rid, recipename, ingredient, steps, image FROM recipe WHERE rid=$1;"

func GetRecipeByID(rid int64) (Recipe, error) {
	var recipe Recipe
	err := postgres.Db.QueryRow(queryGetRecipeByID, rid).Scan(&recipe.Rid, &recipe.RecipeName, &recipe.Ingredient, &recipe.Steps, &recipe.Image)
	if err != nil {
		return Recipe{}, err
	}
	return recipe, nil
}
func (r *Recipe) Update() error {
	if postgres.Db == nil {
		return errors.New("database connection is not initialized")
	}

	query := "UPDATE recipe SET recipename = $1, ingredient = $2, steps = $3, image = $4 WHERE rid = $5"
	_, err := postgres.Db.Exec(query, r.RecipeName, r.Ingredient, r.Steps, r.Image, r.Rid)
	return err
}
func (recipe *Recipe) Delete() error {
	if postgres.Db == nil {
		return errors.New("database connection is not initialized")
	}

	query := "DELETE FROM recipe WHERE rid = $1"
	_, err := postgres.Db.Exec(query, recipe.Rid)
	if err != nil {
		return fmt.Errorf("error deleting recipe: %v", err)
	}

	return nil
}

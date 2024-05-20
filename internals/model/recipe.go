package model

import (
	"fmt"

	postgres "main.go/data"
)

type Recipe struct {
	Rid        int64  `json:"rid"`
	RecipeName string `json:"recipename"`
	Ingredient string `json:"ingredient"`
	Steps      string `json:"steps"`
	Image      []byte `json:"imageurl"`
}

const queryInsertNewRecipe = "INSERT INTO recipe(recipename, ingredient, steps, imageurl) VALUES($1,$2,$3,$4)RETURNING rid;"

// const queryGetAllRecipe="SELECT rid, recipename, ingredient, steps, imageurl FROM recipe;"

func (c *Recipe) Create() error {
	var newRid int64
	err := postgres.Db.QueryRow(queryInsertNewRecipe, c.RecipeName, c.Ingredient, c.Steps, c.Image).Scan(newRid)
	if err != nil {
		return fmt.Errorf("failed to create recipe: %w", err)
	}
	c.Rid = newRid
	return nil
}

// func GetAllRecipe()([]Recipe, error){
//     rows, err:=postgres.Db.Query(queryGetAllRecipe)
//     if err!=nil{
//         return nil, err
//     }
//     defer rows.Close()

//     var recipes []Recipe
//     for rows.Next(){
//         var rec Recipe
//         if err:=rows.Scan(&rec.Rid, &rec.RecipeName, &rec.Ingredient, &rec.Steps, &rec.Image); err!=nil{
//             return nil, err
//         }
//         recipes=append(recipes, rec)
//     }
//     if err := rows.Err();err!=nil{
//         return nil, err
//     }
//     return recipes, nil
// }

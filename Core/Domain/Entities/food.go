package Entities

import (
	"github.com/pluralsight/foodforme/Core/Domain/Enums"
	"github.com/google/uuid"
)

type Food struct {
	id          string
	name        string
	FoodType    string
	Ingredients []Ingredient
	Category    Enums.FoodCategory
}

func NewFood(name, foodType string, ingredients []Ingredient, category Enums.FoodCategory) Food {
	f := Food { uuid.New().String(), name, foodType, ingredients, category}
	return f
}

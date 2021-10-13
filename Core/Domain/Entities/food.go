package Entities

import (
	"github.com/pluralsight/foodforme/Core/Domain/Enums"
)

type Food struct {
	id          int
	name        string
	FoodType    string
	Ingredients []Ingredient
	Category    Enums.FoodCategory
}

func NewFood(name, foodType string, ingredients []Ingredient, category Enums.FoodCategory) *Food {
	return &Food{
		id:          1,
		name:        name,
		FoodType:    foodType,
		Ingredients: ingredients,
		Category:    category,
	}
}

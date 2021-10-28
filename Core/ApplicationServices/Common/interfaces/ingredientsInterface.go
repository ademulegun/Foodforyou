package interfaces

import (
	"github.com/pluralsight/foodforme/Core/ApplicationServices/useCases/ingredient/dto"
	"github.com/pluralsight/foodforme/Core/Domain/Entities"
)

type IngredientsInterface interface {
	SaveIngredient(ingredient Entities.Ingredient) (int, error)
	GetIngredients() ([]dto.GetIngredientDto, error)
	GetIngredient(id string) (dto.GetIngredientDto, error)
	Delete(id string) error
}
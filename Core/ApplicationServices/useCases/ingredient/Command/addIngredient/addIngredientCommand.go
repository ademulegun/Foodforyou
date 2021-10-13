package addingredient

import (
	//"encoding/json"
	"log"

	"github.com/pluralsight/foodforme/Core/ApplicationServices/Common/interfaces"
	"github.com/pluralsight/foodforme/Core/Domain/Entities"
	"github.com/pluralsight/foodforme/Infrastructure/Persistence/repository"
)

type AddIngredientCommand struct {
	Name string
	Category string
}

func (c *AddIngredientCommand) Add(name, category string) (bool, error) {
	var repo interfaces.IngredientsInterface
	ingredient := Entities.NewIngredient(name, category)
	repo = repository.New()
	result, err := repo.SaveIngredient(ingredient)
	if err != nil {
		log.Fatalf(err.Error())
	}
	if result > 0 {
		return true, nil
	}
	return false, err
}
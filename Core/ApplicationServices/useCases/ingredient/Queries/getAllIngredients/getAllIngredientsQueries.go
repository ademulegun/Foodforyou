package getAllIngredients

import (
	"log"
	"github.com/pluralsight/foodforme/Core/ApplicationServices/Common/interfaces"
	"github.com/pluralsight/foodforme/Core/ApplicationServices/useCases/ingredient/dto"
	"github.com/pluralsight/foodforme/Infrastructure/Persistence/repository"
)

type GetAllIngredientsQueries struct{}

func (s *GetAllIngredientsQueries) GetAllIngredients() ([]dto.GetIngredientDto, error) {
	var repo interfaces.IngredientsInterface
	repo = repository.New()
	ingredients, err := repo.GetIngredients()
	if err != nil {
		log.Fatal(err)
	}
	return ingredients, nil
}
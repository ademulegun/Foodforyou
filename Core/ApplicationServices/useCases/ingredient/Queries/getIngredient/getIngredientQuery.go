package getIngredient

import (
	"log"

	"github.com/pluralsight/foodforme/Core/ApplicationServices/Common/interfaces"
	"github.com/pluralsight/foodforme/Core/ApplicationServices/useCases/ingredient/dto"
	"github.com/pluralsight/foodforme/Infrastructure/Persistence/repository"
)

type GetIngredientQuery struct {
	Id int
}

func (s *GetIngredientQuery) GetIngredientById(id string) (dto.GetIngredientDto, error) {
	var repo interfaces.IngredientsInterface
	repo = repository.New()
	ingredient, err := repo.GetIngredient(id)
	if err != nil {
		log.Println(err.Error())
		return ingredient, err
	}
	return ingredient, nil
}
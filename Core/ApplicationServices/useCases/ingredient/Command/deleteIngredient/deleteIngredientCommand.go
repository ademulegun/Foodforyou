package deleteingredient

import (
	"log"

	"github.com/pluralsight/foodforme/Core/ApplicationServices/Common/interfaces"
	"github.com/pluralsight/foodforme/Infrastructure/Persistence/repository"
)

type DeleteIngredient struct {
	Id string
}

func (e *DeleteIngredient) DeleteIngredient(id string) error {
	var repo interfaces.IngredientsInterface
	repo = repository.New()
	err := repo.Delete(id)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
package repository

import (
	"database/sql"

	"github.com/pluralsight/foodforme/Core/ApplicationServices/useCases/ingredient/dto"
	"github.com/pluralsight/foodforme/Core/Domain/Entities"
	"github.com/pluralsight/foodforme/Infrastructure/Persistence"
	"errors"
)

type IngredientRepository struct{}

func New() *IngredientRepository {
	return &IngredientRepository{}
}

func (r *IngredientRepository) SaveIngredient(ingredient Entities.Ingredient) (int, error) {

	result, err := Persistence.DbConn.Exec(`Insert into foodforyou.Ingredients (id, Name, dateAdded, Category) Values (?, ?, ?, ?)`, 
	ingredient.Id(),
	ingredient.IngredientName(),
	ingredient.DateAdded(),
	ingredient.CategoryName())
	if err != nil {
		return 0, err
	}
	inserId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(inserId), nil
}

func (r *IngredientRepository) Delete(id string) error {
	_, err := Persistence.DbConn.Exec(`Delete from foodforyou.Ingredients where id = ? and name = ?`, id)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func (r *IngredientRepository) GetIngredients() ([]dto.GetIngredientDto, error) {
	results, err := Persistence.DbConn.Query("Select id, Name, dateAdded, Category from foodforyou.Ingredients")
	if err != nil {
		return nil, err
	}
	defer results.Close()
	ingredients := make([]dto.GetIngredientDto, 0)
	for results.Next() {
		var ingredient dto.GetIngredientDto
		results.Scan(
			&ingredient.Id, 
			&ingredient.Name, 
			&ingredient.DateAdded,
			&ingredient.Category)
		ingredients = append(ingredients, ingredient)
	}
	return ingredients, nil
}

func (r *IngredientRepository) GetIngredient(id string) (dto.GetIngredientDto, error) {
	row := Persistence.DbConn.QueryRow("Select id, Name, dateAdded, Category from foodforyou.Ingredients where id = ?", id)
	ingredient := dto.GetIngredientDto{}
	err := row.Scan(
			&ingredient.Id, 
			&ingredient.Name, 
			&ingredient.DateAdded,
			&ingredient.Category)
	if err == sql.ErrNoRows {
		return ingredient, nil
	} else if err != nil {
		return ingredient, errors.New(err.Error())
	}
	return ingredient, nil
}
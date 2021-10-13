package Entities

import (
	"errors"
	"time"
	"github.com/google/uuid"
)

type Ingredient struct {
	id        string
	name      string
	dateAdded string
	category string
}

func NewIngredient(name, category string) Ingredient {
	i:= Ingredient { uuid.New().String(), name, time.Now().String(), category}
	return i
}

func (i Ingredient) Id() string{
	return i.id
}

func (i *Ingredient) SetId(id string) error {
	i.id = id
	return nil
}

func (i Ingredient) IngredientName() string{
	return i.name
}

func (i *Ingredient) SetIngredientName(value string) error {
	if len(value) == 0 {
		return errors.New("Ingredient name can not be null")
	}
	i.name = value
	return nil
} 

func (i Ingredient) DateAdded() string {
	return i.dateAdded
}

func (i *Ingredient) SetDate(value time.Time) error {
	i.dateAdded = value.String()
	return nil
}


func (i Ingredient) CategoryName() string{
	return i.category
}

func (i *Ingredient) SetCategory(category string) error{
	if len(category) == 0 {
		return errors.New("Ingredient category can not be")
	}
	i.category = category
	return nil
}

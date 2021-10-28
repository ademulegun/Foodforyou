package Controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	// "strconv"
	"strings"

	addingredient "github.com/pluralsight/foodforme/Core/ApplicationServices/useCases/ingredient/Command/addIngredient"
	deleteingredient "github.com/pluralsight/foodforme/Core/ApplicationServices/useCases/ingredient/Command/deleteIngredient"
	"github.com/pluralsight/foodforme/Core/ApplicationServices/useCases/ingredient/Queries/getAllIngredients"
	"github.com/pluralsight/foodforme/Core/ApplicationServices/useCases/ingredient/Queries/getIngredient"
)

func CreateIngredientHandler(w http.ResponseWriter, r *http.Request) {
	var ingredient addingredient.AddIngredientCommand
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(bodyBytes, &ingredient)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_ , err = ingredient.Add(ingredient.Name, ingredient.Category)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func GetIngredientsHandler(w http.ResponseWriter, r *http.Request) {
	var ingredientInstance getAllIngredients.GetAllIngredientsQueries
	ingredients, err := ingredientInstance.GetAllIngredients()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	j, err := json.Marshal(ingredients)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(j)
}

func GetIngredientHandler(w http.ResponseWriter, r *http.Request) {
	var ingredientInstance getIngredient.GetIngredientQuery
	urlPathSegments := strings.Split(r.URL.Path, "getIngredientById/")
	ingredientId := urlPathSegments[len(urlPathSegments)-1]
	if len(ingredientId) == 0 {
		log.Println("id can not be null")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
    ingredient, err := ingredientInstance.GetIngredientById(ingredientId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	j, err := json.Marshal(ingredient)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(j)
}	

func DeleteIngredientHandler(w http.ResponseWriter, r *http.Request) {
	var ingredientInstance deleteingredient.DeleteIngredient
}
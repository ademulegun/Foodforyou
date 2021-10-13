package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pluralsight/foodforme/Infrastructure/Persistence"
	"github.com/pluralsight/foodforme/Presentation/API/Controllers"
)

func main() {
	Init()
}

func Init() {
	Persistence.SetUpDatabase()
	http.HandleFunc("/addFood", Controllers.CreateFoodHandler)
	http.HandleFunc("/getFood", Controllers.GetFoodHandler)
	http.HandleFunc("/addIngredient", Controllers.CreateIngredientHandler)
	http.HandleFunc("/getIngredients", Controllers.GetIngredientsHandler)
	http.HandleFunc("/getIngredientById/", Controllers.GetIngredientHandler)
	log.Println("Starting server......")
	log.Fatal(http.ListenAndServe(":7000", nil))
	log.Println("Web server started")
}

package Controllers

import (
	// "encoding/json"
	// "io/ioutil"
	"net/http"
)

func CreateFoodHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}

func GetFoodHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}

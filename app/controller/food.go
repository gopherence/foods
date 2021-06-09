package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gopherence/foods/app/model"
)

// FoodHandle responsible for creating and listing foods
func FoodHandle(w http.ResponseWriter, r *http.Request) {
	log.Printf("[foods] %s %s", r.Method, r.URL.Path)

	responseBody := []byte{}
	statusCode := http.StatusMethodNotAllowed

	switch r.Method {
	case "GET":
		responseBody, statusCode = foods(w)
	case "POST":
		responseBody, statusCode = createFood(w, r)
	}

	w.Header().Set("Content-Type", "application/json; charset=ascii")
	w.WriteHeader(statusCode)
	w.Write(responseBody)
}

func foods(w http.ResponseWriter) ([]byte, int) {
	foods, err := model.List()
	if err != nil {
		log.Printf("Error on list foods. %s", err)
	}

	body, err := json.Marshal(foods)
	if err != nil {
		log.Printf("Error on marshall data structure to JSON. %s", err)
		return []byte{}, http.StatusInternalServerError
	}

	return body, http.StatusOK
}

func createFood(w http.ResponseWriter, r *http.Request) ([]byte, int) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error on read request body. %s", err)
		return []byte{}, http.StatusInternalServerError
	}

	var food model.Food
	err = json.Unmarshal(body, &food)
	if err != nil {
		log.Printf("Error on unmarshall JSON to data structure. %s", err)
		return []byte{}, http.StatusBadRequest
	}

	if err := food.Create(); err != nil {
		log.Printf("Error on create food. %s", err)
		return []byte{}, http.StatusInternalServerError
	}

	return body, http.StatusOK
}

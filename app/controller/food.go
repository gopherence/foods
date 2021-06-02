package controller


import (
	"github.com/gopherence/foods/app/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func FoodHandle(w http.ResponseWriter, r *http.Request) {
	log.Printf("[foods] %s %s", r.Method, r.URL.Path)

	switch r.Method {
	case "GET":
		foods(w)
		return
	case "POST":
		createFood(w, r)
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}


func createFood(w http.ResponseWriter, r *http.Request) {
	statusCode := http.StatusOK
	var responseBody interface{}
	var food model.Food
	if err := readJSON(w, r, &food); err != nil {
		log.Printf("Error on convert JSON to food. %s", err)
		statusCode = http.StatusBadRequest
		responseBody = map[string]string{"message": "body is invalid"}
	} else {
		responseBody = food
	}

	if err := food.Create(); err != nil {
		log.Printf("Error on create food. %s", err)
		statusCode = http.StatusInternalServerError
		responseBody = map[string]string{"message": "internal server error"}
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=ascii")
	w.WriteHeader(statusCode)
	writeJSON(w, responseBody)
}

func foods(w http.ResponseWriter) {
	foods, err := model.Find()
	if err != nil {
		log.Printf("Not found foods. %s", err)
	}

	w.Header().Set("Content-Type", "application/json; charset=ascii")
	w.WriteHeader(http.StatusOK)
	writeJSON(w, foods)
}

func writeJSON(w http.ResponseWriter, data interface{}) error {
	body, err := json.Marshal(data)
	if err != nil {
		log.Printf("Couldn't marshall data structure to JSON. %s", err.Error())
		return err
	}

	w.Write(body)
	return nil
}

func readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error on read request body. %s", err.Error())
		return err
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		log.Printf("Error on unmarshall JSON to data structure. %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}
	return err
}
package model

const foodPath = "foods/content"

type Food struct {
	Name string `json:"name"`
}

// Find all foods
func Find() ([]Food, error) {
	return persist.find()
}

// Create register a new food
func (food *Food) Create() error {
	return persist.create(food)
}

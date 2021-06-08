package model

const foodPath = "foods/content"

type Food struct {
	Name string `json:"name"`
}

// List return all foods
func List() ([]Food, error) {
	return persist.list()
}

// Create register a new food
func (food *Food) Create() error {
	return persist.create(food)
}

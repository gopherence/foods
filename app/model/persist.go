package model

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

const dataPath = "var/data"

var persist = &filesystem{}

type filesystem struct {
}

func (fs *filesystem) list() ([]Food, error) {
	foodDataPath := filepath.Join(dataPath, foodPath)

	content, err := ioutil.ReadFile(foodDataPath)
	if err != nil {
		return []Food{}, err
	}

	foods := []Food{}
	if err := json.Unmarshal(content, &foods); err != nil {
		return []Food{}, err
	}

	return foods, nil
}

func (fs *filesystem) create(food *Food) error {
	foodDataPath := filepath.Join(dataPath, foodPath)

	foods, err := fs.list()
	if err != nil {
		return err
	}

	foods = append(foods, *food)

	content, err := json.Marshal(foods)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(foodDataPath, content, os.FileMode(0644))
}

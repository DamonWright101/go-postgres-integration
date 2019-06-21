package logic

import (
	"encoding/json"
)

func BuildPersonFromJson(data string) (Person, error) {
	var people []Person
	err := json.Unmarshal([]byte(data), &people)
	if err != nil {
		return Person{}, err
	}

	return people[0], nil
}

func BuildPersonSliceFromJson(data string) ([]Person, error) {
	var people []Person
	err := json.Unmarshal([]byte(data), &people)
	if err != nil {
		return []Person{}, err
	}

	return people, nil
}

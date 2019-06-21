package logic

import "encoding/json"

type Person struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Fname    string `json:"fname"`
	Lname    string `json:"lname"`
}

func NewPerson(jsonToParse string) (Person, error) {

	var person Person
	jsonInBytes := []byte(jsonToParse)

	err := json.Unmarshal(jsonInBytes, &person)
	if err != nil {
		return Person{}, err
	}

	return person, nil
}

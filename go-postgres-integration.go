package main

import (
	"fmt"
	"log"

	"github.com/DamonWright101/go-postgres-integration/http"
	"github.com/DamonWright101/go-postgres-integration/logic"
	_ "github.com/lib/pq"
)

func main() {

	TestPostgrest()

	http.Server(8080)

}

func TestPostgrest() {

	// get a single person from db and convert it into person obj
	data, err := logic.GetRowFromTableByExactMatch("persons", "username", "username003")
	if err != nil {
		log.Fatalf("%+v", err)
	}

	person, err := logic.BuildPersonFromJson(data)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	fmt.Printf("people: %+v\n\n", person)

	// get a set of people from db and convert it into slice person obj
	data, err = logic.GetRowsFromTableByBetweenTwoValues("persons", "username", "username003", "username006")
	if err != nil {
		log.Fatalf("%+v", err)
	}
	people, err := logic.BuildPersonSliceFromJson(data)

	fmt.Println("people")
	for i := range people {
		fmt.Printf("%d: %+v\n", i, people[i])
	}

}

func unused() {

	somejson := `{ "id": 3, "email": "test3@email.com", "username": "username3", "fname": "test3", "lname": "test" }`
	person, err := logic.NewPerson(somejson)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	data, err := logic.Query("http://localhost:3000/persons?id=eq.4")
	if err != nil {
		log.Fatalf("%+v", err)
	}

	log.Print(person, data, err)
}

package main

import (
	"encoding/json"
	"fmt"
	"go_pl/shared/github"
	"log"
	"os"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {

	dilbert := Employee{ID: 2, Position: "developer"}
	fmt.Println(dilbert)

	position := &dilbert.Position
	*position = "Senior " + *position
	fmt.Println(dilbert.Position)

	dilbert.Salary = 20
	fmt.Println(dilbert)
	RaiseSalary(&dilbert, 20)
	fmt.Println(dilbert)

	fmt.Println("----------")

	var movies = []Movie{
		{Title: "Calablanca", Year: 1942, Color: false, Actors: []string{"Bogart", "Bergman"}},
		{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"McQueen", "Bisset"}},
	}
	fmt.Println(movies)

	//data, err := json.Marshal(movies)
	data, err := json.MarshalIndent(movies, "", " ")

	if err != nil {
		log.Fatalf("Сбой маршалинга JSON: %s", err)
	}
	fmt.Printf("%s\n", data)

	fmt.Println("------------")
	var parsed []struct {
		Title  string
		Actors []string
		Color  bool
	}
	if err := json.Unmarshal(data, &parsed); err != nil {
		log.Fatalf("Demarshal failed: $s", err)
	}

	fmt.Println(parsed)
	terms := []string{"Unhandled", "response", "error"}
	results, err := github.SearchIssues(terms)
	if err != nil {
		log.Fatalf("Issues search failed: %s", err)
	}

	github.PrintIssues(os.Stdout, results)
}

func RaiseSalary(e *Employee, v int) {
	e.Salary = e.Salary + v
}

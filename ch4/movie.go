package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	type Movie struct {
		Title    string
		Year     int  `json:"released"`
		Color    bool `json:"color,omitempty"`
		Actors   []string
		Figurant string
	}

	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}, Figurant: "Edmilson"},
	}

	data, err := json.MarshalIndent(movies, "", "  ")

	if err != nil {
		log.Fatalf("JSON Marshaling failed: %s", err)
	}

	fmt.Printf("%s\n", data)

	var titles []struct{ Title string }

	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON umarshaling failed: %s", err)
	}

	fmt.Println(titles)
}

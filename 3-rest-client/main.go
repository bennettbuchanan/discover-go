package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type IMDB struct {
	Title      string `json:"Title"`
	Year       string `json:"Year"`
	Rated      string `json:"Rated"`
	Released   string `json:"Released"`
	Runtime    string `json:"Runtime"`
	Genre      string `json:"Genre"`
	Director   string `json:"Director"`
	Writer     string `json:"Writer"`
	Actors     string `json:"Actors"`
	Plot       string `json:"Plot"`
	Language   string `json:"Language"`
	Country    string `json:"Country"`
	Awards     string `json:"Awards"`
	Poster     string `json:"Poster"`
	Metascore  string `json:"Metascore"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes  string `json:"imdbVotes"`
	ImdbID     string `json:"imdbID"`
	Type       string `json:"Type"`
	Response   string `json:"Response"`
}

func main() {
	url := "http://www.omdbapi.com/?i=tt0372784&plot=short&r=json"
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("%s", err)
	}

	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var data IMDB
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Printf("%s", err)
	}

	rating, err := strconv.ParseFloat(data.ImdbRating, 64)
	if err != nil {
		fmt.Printf("%s", err)
	}

	i := rating * 10
	y := int(i)
	fmt.Printf("The movie : %s was released in %s - the IMDB rating is %d%% with %s votes.", data.Title, data.Year, y, data.ImdbVotes)
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

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
	fmt.Printf("The movie : %s was released in %s - the IMBD rating is %d%% with %s votes\n", data.Title, data.Year, y, data.ImdbVotes)
}

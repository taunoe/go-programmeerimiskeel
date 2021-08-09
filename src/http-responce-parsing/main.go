package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Read web page, json etc
	s1 := get_content_as_string("https://dog.ceo/api/breeds/image/random")
	fmt.Println(s1)

	s2 := get_content_as_struct("https://dog.ceo/api/breeds/image/random")
	fmt.Println(s2)
	fmt.Println(s2.Status)
	fmt.Println(s2.Message)
}

// For func get_content_as_struct
type doge_responce struct {
	Message string
	Status  string
}

func fetch_url(url string) []byte {
	// web request
	r, err := http.Get(url)

	// if error: exit
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	bytes, err := ioutil.ReadAll(r.Body)
	return bytes
}

func get_content_as_struct(url string) doge_responce {
	bytes := fetch_url(url)

	var data doge_responce // loo struct

	// Parse json to struct
	parsing_err := json.Unmarshal(bytes, &data)
	if parsing_err != nil {
		panic(parsing_err)
	}

	return data
}

func get_content_as_string(url string) string {
	bytes := fetch_url(url)

	return string(bytes)
}

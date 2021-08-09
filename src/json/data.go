package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// DataContainer structure
type DataContainer struct {
	Name string
	Data []*Data
}

// Data structure == json stucture
type Data struct {
	Title       string
	Link        string
	Description string
	Content     string
}

func handle_err(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func get_data(file_path string) DataContainer {
	res, err := ioutil.ReadFile(file_path) // byte array
	handle_err(err)

	var data DataContainer // Data holder
	jErr := json.Unmarshal(res, &data)
	handle_err(jErr)

	return data
}

func get_content(path string) string {
	content, err := ioutil.ReadFile(path)
	handle_err(err)

	return string(content)
}

func write_file(filename string, content []byte) {
	premissions := os.FileMode(07555)
	err := ioutil.WriteFile(filename, content, premissions)
	handle_err(err)
}

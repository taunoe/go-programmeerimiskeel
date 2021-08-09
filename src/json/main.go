package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	json_path := "./data/data.json"
	new_file := "./data/new.json"
	data := get_data(json_path)

	// Print json file content
	/*
		for _, dc := range data.Data {
			fmt.Println(*dc)
		}
	*/

	for _, file := range data.Data {
		file.Content = get_content(file.Link)
		fmt.Println(*&file.Content)
	}

	output, err := json.Marshal(data)
	handle_err(err)
	write_file(new_file, output)
}

package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}

type Author struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Developer bool   `json:"is_developer"`
}

func main() {
	fmt.Println("hi")

	book := Book{Title: "LOTR",
		Author: Author{
			Name:      "Tolkien",
			Age:       0,
			Developer: false,
		}}

	fmt.Printf("%+v \n", book)

	byteArray, err := json.MarshalIndent(book, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))
}

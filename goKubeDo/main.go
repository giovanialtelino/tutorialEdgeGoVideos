package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "my go app")

}

func setupRoutes() {
	http.HandleFunc("/", homePage)
}

func main() {
	setupRoutes()
	http.ListenAndServe(":3000", nil)
}

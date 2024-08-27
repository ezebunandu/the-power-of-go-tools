package main

import (
	"fmt"
	"net/http"

	"github.com/ezebunandu/hello"
)

func main() {
	fmt.Println("Listening on http://localhost:8222")
	http.ListenAndServe(":8222", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request){
			hello.PrintTo(w)
		}),
	)
}

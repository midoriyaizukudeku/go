package main

import (
	"fmt"
	"net/http"
)

func main() {

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {http.ServeFile(w, r, "./public/index.html")})

	fmt.Println("server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
	
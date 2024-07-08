package main

import (
	"fmt"
	"net/http"

	web "web/packages/server"
)

func main() {
	fmt.Println("http://localhost:8080")
	http.HandleFunc("/", web.HomeHandler)
	http.HandleFunc("/ascii-art", web.AsciiArtHandler)

	// Define the file system root
	fs := http.FileServer(http.Dir("./static"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
}

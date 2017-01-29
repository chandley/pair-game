package main

import "fmt"
import (
	"net/http"
	"log"
	"io"
	"os"
)

func WebHandler (w http.ResponseWriter, r *http.Request) {
	img, err := os.Open("./kitten.jpg")
	if err != nil {
		log.Fatal(err) // perhaps handle this nicer
	}
	defer img.Close()
	w.Header().Set("Content-Type", "image/jpeg") // <-- set the content-type header
	io.Copy(w, img)
}

func main()  {
	fmt.Println("started")
	http.HandleFunc("/", WebHandler)
	http.ListenAndServe(":8080", nil)
}

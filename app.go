package main

import "fmt"
import (
	"net/http"
	"log"
	"io"
	"os"
	"html/template"
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

func handler (w http.ResponseWriter, r *http.Request) {
	var myInfo = struct{
		Image string
	}{"kitten"}

	const templ = `<body><img src="http://localhost:8080/images/{{.Image}}.jpg"><p>You clicked {{.Image}}!</p><body>`
	reports, err := template.New("report").
		Parse(templ)
	if err != nil {
		log.Fatal(err)
	}

	err = reports.Execute(w, myInfo)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	const text = `<html>
	<body>
	<a href="http://localhost:8080/clicked"><img src="http://localhost:8080/images/kitten.jpg"/></a>
	</body>
	</html>`
	fmt.Fprintf(w, text)
}

func main()  {
	fmt.Println("started")
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/clicked", handler)
	http.HandleFunc("/images/", WebHandler)
	http.ListenAndServe(":8080", nil)
}

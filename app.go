package main

import "fmt"
import (
	"net/http"
	"log"
	"html/template"
)

func webHandler (w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
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

	type cell struct  {
		Animal string
	}

	puppy := cell{"puppy"}
	kitten := cell{"kitten"}

	var myrow []cell

	myrow = append(myrow, kitten)
	myrow = append(myrow, puppy)
	myrow = append(myrow, kitten)

	var myInfo = struct{
		Image string
		Row []cell
	}{"puppy", myrow,
	}

	const templ = `<body>
	<table>
	{{ $trow := .Row }}

	<tr>
	{{range $tcell := $trow}}
		<a href="http://localhost:8080/clicked"><img src="http://localhost:8080/images/{{$tcell.Animal}}.jpg"></a>
	{{end}}
	</tr>

	</table>
	<p>{{.Image}}</p>
	<body>`

	const cellTemplate = `<a href="http://localhost:8080/clicked"><img src="http://localhost:8080/images/{{.Animal}}.jpg"></a>`

	var loopFunc = func(n int) []struct{} {
		return make([]struct{}, n)
	}

	reports := template.Must(template.New("report").Funcs(template.FuncMap{
		"loop": loopFunc,
	}).Parse(templ))

	reports.Execute(w, myInfo)
}

func main()  {
	fmt.Println("started")
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/clicked", handler)
	http.HandleFunc("/images/", webHandler)
	http.ListenAndServe(":8080", nil)
}

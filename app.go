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

	type row struct {
		Cells []cell
	}

	type board struct {
		Rows []row
	}

	//puppy := cell{"puppy"}
	//kitten := cell{"kitten"}

	first := row{
		[]cell{
			cell{"puppy"},
			cell{"puppy"},
			cell{"kitten"},
		},
	}

	second := row{
		[]cell{
			cell{"kitten"},
			cell{"puppy"},
			cell{"kitten"},
		},
	}

	third := row{
		[]cell{
			cell{"kitten"},
			cell{"kitten"},
			cell{"puppy"},
		},
	}

	pets := board{
		[]row{
			first, second, third,
		},
	}


	const templ = `<body>

	{{ range $rowIndex, $trow := .Rows }}
		<div width="100%">

		{{range $colIndex, $tcell := .Cells}}

			<a href="http://localhost:8080/clicked/{{$rowIndex}}/{{$colIndex}}"><img style="width:200px;height:200px;object-fit: cover" src="http://localhost:8080/images/{{$tcell.Animal}}.jpg"></a>

		{{end}}

		<div>

	{{ end }}


	<body>`

	const cellTemplate = `<a href="http://localhost:8080/clicked"><img src="http://localhost:8080/images/{{.Animal}}.jpg"></a>`

	var loopFunc = func(n int) []struct{} {
		return make([]struct{}, n)
	}

	reports := template.Must(template.New("report").Funcs(template.FuncMap{
		"loop": loopFunc,
	}).Parse(templ))

	reports.Execute(w, pets)
}

func main()  {
	fmt.Println("started")
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/clicked", handler)
	http.HandleFunc("/images/", webHandler)
	http.ListenAndServe(":8080", nil)
}

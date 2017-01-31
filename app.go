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
	var myInfo = struct{
		Image string
	}{"puppy"}

	const templ = `<body>
	<table>
	{{range loop 3}}
		<tr>
		{{range loop 3}}
			<img src="http://localhost:8080/images/puppy.jpg">
		{{end}}
		</tr>
	{{end}}
	</table>
	<p>{{.Image}}</p>
	<body>`

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

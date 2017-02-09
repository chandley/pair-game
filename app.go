package main

import "fmt"
import (
	"net/http"
	"html/template"
	"strconv"
)

type cell struct  {
	Animal string
}

type row struct {
	Cells []cell
}

type board struct {
	Rows []row
}




func webHandler (w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func (s *server) clickHandler(w http.ResponseWriter, r *http.Request) {
	locationString := r.URL.Path[len("/clicked/"):]
	rowIndex, _ := strconv.Atoi(locationString[:1])
	colIndex, _ := strconv.Atoi(locationString[2:])
	fmt.Println("got a click")
	s.currentBoard.Rows[rowIndex].Cells[colIndex] = cell{"martha"}
	http.Redirect(w, r, "/board/", http.StatusFound)
	return
}

type server struct {
	currentBoard *board
}



func (s *server) viewHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Println("showing board")
	const templ = `
	<body>
	{{ range $rowIndex, $trow := .Rows }}
		<div width="100%">
		{{range $colIndex, $tcell := .Cells}}

			<a href="http://localhost:8080/clicked/{{$rowIndex}}/{{$colIndex}}">
			<img style="width:200px; height:200px; object-fit:cover" src="http://localhost:8080/images/{{$tcell.Animal}}.jpg">
			</a>
		{{end}}
		<div>
	{{ end }}
	<body>`

	gameBoard := template.Must(template.New("gameBoard").Parse(templ))

	gameBoard.Execute(w, s.currentBoard)
}

func main()  {
	myBoard := getNewBoard()
	server := &server{myBoard}
	fmt.Println("started")
	http.HandleFunc("/board/", server.viewHandler)
	http.HandleFunc("/clicked/", server.clickHandler)
	http.HandleFunc("/images/", webHandler)
	http.ListenAndServe(":8080", nil)
}

func getNewBoard() *board {
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

	myboard := board{
		[]row{
			first, second, third,
		},
	}

	return &myboard
}

func getAllPuppies() *board {
	first := row{
		[]cell{
			cell{"puppy"},
			cell{"puppy"},
			cell{"puppy"},
		},
	}

	second := row{
		[]cell{
			cell{"puppy"},
			cell{"puppy"},
			cell{"puppy"},
		},
	}

	third := row{
		[]cell{
			cell{"puppy"},
			cell{"puppy"},
			cell{"puppy"},
		},
	}

	myboard := board{
		[]row{
			first, second, third,
		},
	}

	return &myboard
}

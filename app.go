package main

import "fmt"
import (
	"net/http"
	"html/template"
	"strconv"
	"time"
)

type cell struct  {
	Animal string
	Clicked bool
	Paired bool
	Visible bool
}


type board struct {
	Cells [][]cell
}

type server struct {
	currentBoard *board
}

func webHandler (w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func (s *server) clickHandler(w http.ResponseWriter, r *http.Request) {
	if s.clickedCount() >= 2 {
		s.resetClicked()
	}
	locationString := r.URL.Path[len("/clicked/"):]
	rowIndex, _ := strconv.Atoi(locationString[:1])
	colIndex, _ := strconv.Atoi(locationString[2:])
	s.currentBoard.Cells[rowIndex][colIndex].Clicked = true
	s.currentBoard.Cells[rowIndex][colIndex].Visible = true
	http.Redirect(w, r, "/board/", http.StatusFound)
	s.checkForClickedPair()
	return
}

func (s *server) clickedCount() (clickedCount int) {
	for _, thisRow := range s.currentBoard.Cells {
		for _, thisCell := range thisRow {
			if thisCell.Clicked  {
				clickedCount += 1
			}
		}
	}
	return
}

func (s * server) resetClicked() {
	for i, thisRow := range s.currentBoard.Cells {
		for j, thisCell := range thisRow {
			if thisCell.Clicked  {
				s.currentBoard.Cells[i][j].Clicked = false
				if !thisCell.Paired {
					s.currentBoard.Cells[i][j].Visible = false
				}
			}
		}
	}
	return
}

func (s *server) resetIfTwoClicked(w http.ResponseWriter, r *http.Request) {
	if s.clickedCount() >= 2 {
		time.Sleep(500 * time.Millisecond)
		s.resetClicked()
		http.Redirect(w, r, "/board/", http.StatusFound)

	}
}

func (s *server) checkForClickedPair() (bool){
	type indexPair struct{
		x int
		y int
	}
	clickedIndices := []indexPair{}
	for i, thisRow := range s.currentBoard.Cells {
		for j, thisCell := range thisRow {
			if thisCell.Clicked  {
				clickedIndices = append(clickedIndices, indexPair{i,j,})
			}
		}
	}
	if len(clickedIndices) != 2 {
		return false
	}

	if s.currentBoard.Cells[clickedIndices[0].x][clickedIndices[0].y].Animal ==  s.currentBoard.Cells[clickedIndices[1].x][clickedIndices[1].y].Animal {
		fmt.Println("detected a pair", clickedIndices)
		s.currentBoard.Cells[clickedIndices[0].x][clickedIndices[0].y].Paired = true
		s.currentBoard.Cells[clickedIndices[1].x][clickedIndices[1].y].Paired = true
		return true
	}

	return false
}



func (s *server) viewHandler (w http.ResponseWriter, r *http.Request) {

	fmt.Println("showing board")
	const templ = `
	<body>
	{{ range $rowIndex, $trow := .Cells }}
		<div width="100%">
		{{range $colIndex, $tcell := $trow}}
			<a href="http://localhost:8080/clicked/{{$rowIndex}}/{{$colIndex}}">
			<img style="width:200px; height:200px; object-fit:cover" src="http://localhost:8080/images/{{if $tcell.Visible}}{{$tcell.Animal}}.jpg{{else}}mergermarket.jpg {{end}}">
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

	cells := [][]cell{
		{
			{Animal: "puppy"},
			{Animal: "kitten"},
			{Animal: "martha"},
			{Animal: "kitten"},
		},
		{
			{Animal: "kitten"},
			{Animal: "puppy"},
			{Animal: "kitten"},
			{Animal: "puppy"},
		},
		{
			{Animal: "martha"},
			{Animal: "kitten"},
			{Animal: "puppy"},
			{Animal: "kitten"},
		},
	}

	board := board{cells}

	return &board
}


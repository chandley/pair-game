package main

import "fmt"
import (
	"net/http"
	"html/template"
	"strconv"
	"time"
	"math/rand"
)

type server struct {
	board *Board
}

func main()  {
	server := &server{getNewBoard()}
	fmt.Println("started")
	http.HandleFunc("/", server.viewHandler)
	http.HandleFunc("/reset/", server.resetHandler)
	http.HandleFunc("/clicked/", server.clickHandler)
	http.HandleFunc("/images/", webHandler)
	http.ListenAndServe(":8080", nil)
}

func (s *server) viewHandler (w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("board.html"))
	template.Execute(w, s.board)
}

func (s *server) clickHandler(w http.ResponseWriter, r *http.Request) {
	locationString := r.URL.Path[len("/clicked/"):]
	rowIndex, _ := strconv.Atoi(locationString[:1])
	colIndex, _ := strconv.Atoi(locationString[2:])

	if s.board.ClickedCount() >= 2 {
		s.board.ResetClicked()
	}

	s.board.Cells[rowIndex][colIndex].Clicked = true
	s.board.Cells[rowIndex][colIndex].Visible = true
	s.board.checkForClickedPair()
	http.Redirect(w, r, "/", http.StatusFound)
	return
}

func (s *server) resetHandler (w http.ResponseWriter, r *http.Request) {
	s.board = getNewBoard()
	http.Redirect(w, r, "/", http.StatusFound)
}

func webHandler (w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func getNewBoard() *Board {
	cells := [][]Cell{
		{
			{Animal: "puppy"},
			{Animal: "puppy"},
			{Animal: "martha"},
			{Animal: "martha"},
		},
		{
			{Animal: "kitten"},
			{Animal: "kitten"},
			{Animal: "cora"},
			{Animal: "cora"},
		},
		{
			{Animal: "dino"},
			{Animal: "dino"},
			{Animal: "reggie"},
			{Animal: "reggie"},
		},
	}

	board := Board{cells}
	board.Shuffle()
	return &board
}

type Cell struct  {
	Animal string
	Clicked bool
	Paired bool
	Visible bool
}

type Board struct {
	Cells [][]Cell
}

func (b *Board) ClickedCount() (clickedCount int) {
	for _, thisRow := range b.Cells {
		for _, thisCell := range thisRow {
			if thisCell.Clicked  {
				clickedCount += 1
			}
		}
	}
	return
}

func (b *Board) ResetClicked() {
	for i, thisRow := range b.Cells {
		for j, thisCell := range thisRow {
			if thisCell.Clicked  {
				b.Cells[i][j].Clicked = false
				if !thisCell.Paired {
					b.Cells[i][j].Visible = false
				}
			}
		}
	}
	return
}

func (b *Board) Shuffle() {
	rand.Seed(time.Now().UTC().UnixNano())
	for i, row := range b.Cells {
		for j, _ := range row {
			x := rand.Intn(len(b.Cells))
			y := rand.Intn(len(b.Cells[0]))
			fmt.Println(x,y)
			b.Cells[i][j], b.Cells[x][y] = b.Cells[x][y], b.Cells[i][j]
		}
	}
}

func (b *Board) checkForClickedPair() {

	clickedCells := []*Cell{}
	for i, thisRow := range b.Cells {
		for j, thisCell := range thisRow {
			if thisCell.Clicked  {
				clickedCells = append(clickedCells, &b.Cells[i][j])
			}
		}
	}
	if len(clickedCells) != 2 {
		return
	}

	if clickedCells[0].Animal ==  clickedCells[1].Animal {
		clickedCells[0].Paired = true
		clickedCells[1].Paired = true
	}
}





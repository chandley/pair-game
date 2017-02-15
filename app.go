package main

import "fmt"
import (
	"net/http"
	"html/template"
	"strconv"
	"time"
	"math/rand"
)

func main()  {
	server := &server{getNewBoard()}
	fmt.Println("started")
	http.HandleFunc("/", server.viewHandler)
	http.HandleFunc("/reset/", server.resetHandler)
	http.HandleFunc("/clicked/", server.clickHandler)
	http.HandleFunc("/images/", webHandler)
	http.ListenAndServe(":8080", nil)
}

type server struct {
	currentBoard *Board
}

func webHandler (w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func (s *server) clickHandler(w http.ResponseWriter, r *http.Request) {
	if s.currentBoard.ClickedCount() >= 2 {
		s.currentBoard.ResetClicked()
	}
	locationString := r.URL.Path[len("/clicked/"):]
	rowIndex, _ := strconv.Atoi(locationString[:1])
	colIndex, _ := strconv.Atoi(locationString[2:])
	s.currentBoard.Cells[rowIndex][colIndex].Clicked = true
	s.currentBoard.Cells[rowIndex][colIndex].Visible = true
	http.Redirect(w, r, "/", http.StatusFound)
	s.checkForClickedPair()
	return
}

func (s *server) resetIfTwoClicked(w http.ResponseWriter, r *http.Request) {
	if s.currentBoard.ClickedCount() >= 2 {
		time.Sleep(500 * time.Millisecond)
		s.currentBoard.ResetClicked()
		http.Redirect(w, r, "/", http.StatusFound)

	}
}

func (s *server) checkForClickedPair() {

	clickedCells := []*Cell{}
	for i, thisRow := range s.currentBoard.Cells {
		for j, thisCell := range thisRow {
			if thisCell.Clicked  {
				clickedCells = append(clickedCells, &s.currentBoard.Cells[i][j])
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

func (s *server) resetHandler (w http.ResponseWriter, r *http.Request) {
	s.currentBoard = getNewBoard()
	http.Redirect(w, r, "/", http.StatusFound)
}


func (s *server) viewHandler (w http.ResponseWriter, r *http.Request) {
	gameBoard := template.Must(template.ParseFiles("board.html"))
	gameBoard.Execute(w, s.currentBoard)
}

//func (s *server) shuffle() {
//	rand.Seed(time.Now().UTC().UnixNano())
//	for i, row := range s.currentBoard.Cells {
//		for j, _ := range row {
//			a := rand.Intn(len(s.currentBoard.Cells))
//			b := rand.Intn(len(s.currentBoard.Cells[0]))
//			s.currentBoard.Cells[i][j], s.currentBoard.Cells[a][b] = s.currentBoard.Cells[a][b], s.currentBoard.Cells[i][j]
//		}
//	}
//}

func getNewBoard() *Board {
	cells := [][]Cell{
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
			b.Cells[i][j], b.Cells[x][y] = b.Cells[x][y], b.Cells[i][j]
		}
	}
}





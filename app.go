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

type row struct {
	Cells []cell
}

type board struct {
	Rows []row
}

type server struct {
	currentBoard *board
}

func webHandler (w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func (s *server) clickHandler(w http.ResponseWriter, r *http.Request) {
	locationString := r.URL.Path[len("/clicked/"):]
	rowIndex, _ := strconv.Atoi(locationString[:1])
	colIndex, _ := strconv.Atoi(locationString[2:])
	s.currentBoard.Rows[rowIndex].Cells[colIndex].Clicked = true
	s.currentBoard.Rows[rowIndex].Cells[colIndex].Visible = !(s.currentBoard.Rows[rowIndex].Cells[colIndex].Visible)
	http.Redirect(w, r, "/board/", http.StatusFound)
	return
}

func (s *server) clickedCount() (clickedCount int) {
	for _, thisRow := range s.currentBoard.Rows {
		for _, thisCell := range thisRow.Cells {
			if thisCell.Clicked  {
				clickedCount += 1
			}
		}
	}
	return
}

func (s * server) resetClicked() {
	for i, thisRow := range s.currentBoard.Rows {
		for j, thisCell := range thisRow.Cells {
			if thisCell.Clicked  {
				s.currentBoard.Rows[i].Cells[j].Clicked = false
				if !thisCell.Paired {
					s.currentBoard.Rows[i].Cells[j].Visible = false
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
	clickedCells := []*cell{}
	clickedIndices := []indexPair{}
	for i, thisRow := range s.currentBoard.Rows {
		for j, thisCell := range thisRow.Cells {
			if thisCell.Clicked  {
				clickedCells = append(clickedCells, &thisCell)
				clickedIndices = append(clickedIndices, indexPair{i,j,})
			}
		}
	}
	if len(clickedCells) != 2 {
		return false
	}

	if clickedCells[0].Animal == clickedCells[1].Animal {
		fmt.Println("yes pair yay")
		s.currentBoard.Rows[clickedIndices[0].x].Cells[clickedIndices[0].y].Paired = true
		s.currentBoard.Rows[clickedIndices[1].x].Cells[clickedIndices[1].y].Paired = true

		clickedCells[0].Paired = true
		clickedCells[1].Paired = true
		return true
	}

	return false
}



func (s *server) viewHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Println("showing board")
	const templ = `
	<body>
	{{ range $rowIndex, $trow := .Rows }}
		<div width="100%">
		{{range $colIndex, $tcell := .Cells}}
			<a href="http://localhost:8080/clicked/{{$rowIndex}}/{{$colIndex}}">
			<img style="width:200px; height:200px; object-fit:cover" src="http://localhost:8080/images/{{if $tcell.Visible}}{{$tcell.Animal}}.jpg{{else}}mergermarket.jpg {{end}}">
			</a>
		{{end}}
		<div>
	{{ end }}
	<body>`

	gameBoard := template.Must(template.New("gameBoard").Parse(templ))

	gameBoard.Execute(w, s.currentBoard)

	if s.clickedCount() >= 2 {
		time.Sleep(500 * time.Millisecond)
		s.resetClicked()
		http.Redirect(w, r, "/board/", http.StatusFound) // TODO does not redirect!
	}
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
			cell{Animal: "puppy"},
			cell{Animal: "kitten"},
			cell{Animal: "martha"},
		},
	}

	second := row{
		[]cell{
			cell{Animal: "puppy"},
			cell{Animal: "kitten"},
			cell{Animal: "puppy"},
		},
	}

	third := row{
		[]cell{
			cell{Animal: "kitten"},
			cell{Animal: "puppy"},
			cell{Animal: "kitten"},
		},
	}

	myboard := board{
		[]row{
			first, second, third,
		},
	}

	return &myboard
}

//func getAllPuppies() *board {
//	first := row{
//		[]cell{
//			cell{"puppy"},
//			cell{"puppy"},
//			cell{"puppy"},
//		},
//	}
//
//	second := row{
//		[]cell{
//			cell{"puppy"},
//			cell{"puppy"},
//			cell{"puppy"},
//		},
//	}
//
//	third := row{
//		[]cell{
//			cell{"puppy"},
//			cell{"puppy"},
//			cell{"puppy"},
//		},
//	}
//
//	myboard := board{
//		[]row{
//			first, second, third,
//		},
//	}
//
//	return &myboard
//}

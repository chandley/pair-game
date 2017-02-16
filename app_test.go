package main

import "testing"

func getServerWithOneClicked() server {
	cells := [][]Cell{
		{
			Cell{Animal: "puppy"},
			Cell{Animal: "puppy"},
		},
		{
			Cell{Animal: "puppy"},
			Cell{Animal: "puppy"},

		},
	}


	myboard := Board{cells}
	s := server{&myboard}
	s.board.Cells[0][0].Clicked = true
	return s
}

func TestDontMarkClickedNotPair(t *testing.T) {
	s := getServerWithOneClicked()

	s.board.Cells[1][0] = Cell{Animal: "kitten", Clicked: true}
	s.board.checkForClickedPair()

	if s.board.Cells[1][0].Paired || s.board.Cells[0][0].Paired {
		t.Error ("puppy and kitten should not be marked as pair",
			s.board.Cells[1][0].Paired,
			s.board.Cells[0][0].Paired)
	}

}

func TestMarkClickedPair(t *testing.T) {
	s := getServerWithOneClicked()

	s.board.Cells[1][0] = Cell{Animal: "puppy", Clicked: true}
	s.board.checkForClickedPair()

	if !(s.board.Cells[1][0].Paired && s.board.Cells[0][0].Paired) {
		t.Error ("puppy and puppy should both be marked as pair",
			s.board.Cells[1][0].Paired,
			s.board.Cells[0][0].Paired)
	}

}

package main

import "testing"

func TestCountClicked(t *testing.T) {
	s := getServerWithOneClicked()

	if s.clickedCount() != 1 {
		t.Error("should count one clicked cell")
	}
}
func getServerWithOneClicked() server {
	cells := [][]cell{
		{
			cell{Animal: "puppy"},
			cell{Animal: "puppy"},
		},
		{
			cell{Animal: "puppy"},
			cell{Animal: "puppy"},

		},
	}


	myboard := board{cells}
	s := server{&myboard}
	s.currentBoard.Cells[0][0].Clicked = true
	return s
}

func TestResetClicked(t *testing.T) {
	s := getServerWithOneClicked()

	s.resetClicked()

	if s.clickedCount() != 0 {
		t.Error("reset clicked should reset to zero")
	}
}

func TestDontMarkClickedNotPair(t *testing.T) {
	s := getServerWithOneClicked()

	s.currentBoard.Cells[1][0] = cell{Animal: "kitten", Clicked: true}
	s.checkForClickedPair()

	if s.currentBoard.Cells[1][0].Paired || s.currentBoard.Cells[0][0].Paired {
		t.Error ("puppy and kitten should not be marked as pair",
			s.currentBoard.Cells[1][0].Paired,
			s.currentBoard.Cells[0][0].Paired)
	}

}

func TestMarkClickedPair(t *testing.T) {
	s := getServerWithOneClicked()

	s.currentBoard.Cells[1][0] = cell{Animal: "puppy", Clicked: true}
	s.checkForClickedPair()

	if !(s.currentBoard.Cells[1][0].Paired && s.currentBoard.Cells[0][0].Paired) {
		t.Error ("puppy and puppy should both be marked as pair",
			s.currentBoard.Cells[1][0].Paired,
			s.currentBoard.Cells[0][0].Paired)
	}

}

package main

import "testing"

func TestCountClicked(t *testing.T) {
	s := getServerWithOneClicked()

	if s.clickedCount() != 1 {
		t.Error("should count one clicked cell")
	}
}
func getServerWithOneClicked() server {
	first := row{
		[]cell{
			cell{Animal: "puppy"},
			cell{Animal: "puppy"},
			cell{Animal: "puppy"},
		},
	}
	myboard := board{
		[]row{
			first,
		},
	}
	myboard.Rows[0].Cells[0].Clicked = true
	s := server{&myboard}
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

	s.currentBoard.Rows[0].Cells[1] = cell{Animal: "kitten", Clicked: true}
	s.checkForClickedPair()

	if s.currentBoard.Rows[0].Cells[1].Paired || s.currentBoard.Rows[0].Cells[0].Paired {
		t.Error ("puppy and kitten should not be marked as pair",
			s.currentBoard.Rows[0].Cells[1].Paired,
			s.currentBoard.Rows[0].Cells[0].Paired)
	}

}

func TestMarkClickedPair(t *testing.T) {
	s := getServerWithOneClicked()

	s.currentBoard.Rows[0].Cells[1] = cell{Animal: "puppy", Clicked: true}
	s.checkForClickedPair()

	if !(s.currentBoard.Rows[0].Cells[1].Paired && s.currentBoard.Rows[0].Cells[0].Paired) {
		t.Error ("puppy and puppy should both be marked as pair",
			s.currentBoard.Rows[0].Cells[1].Paired,
			s.currentBoard.Rows[0].Cells[0].Paired)
	}

}

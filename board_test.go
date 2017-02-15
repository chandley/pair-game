package main

import "testing"

func TestBoard_ClickedCount(t *testing.T) {
	s := getServerWithOneClicked()

	if s.currentBoard.ClickedCount() != 1 {
		t.Error("should count one clicked cell")
	}
}

func TestResetClicked(t *testing.T) {
	s := getServerWithOneClicked()

	s.currentBoard.ResetClicked()

	if s.currentBoard.ClickedCount() != 0 {
		t.Error("reset clicked should reset to zero")
	}
}


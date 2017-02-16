package main

import "testing"

func TestBoard_ClickedCount(t *testing.T) {
	s := getServerWithOneClicked()

	if s.board.ClickedCount() != 1 {
		t.Error("should count one clicked cell")
	}
}

func TestBoard_ResetClicked(t *testing.T) {
	s := getServerWithOneClicked()

	s.board.ResetClicked()

	if s.board.ClickedCount() != 0 {
		t.Error("reset clicked should reset to zero")
	}
}


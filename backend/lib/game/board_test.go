package game

import (
	"testing"

	myfirestore "github.com/feloy/yatzy.io/backend/lib/firestore"
)

func TestNewBoard2(t *testing.T) {
	board := NewBoard([]string{"user1", "user2"})
	if len(board) != 2 {
		t.Fatalf("Board should contain 2 elements")
	}
	if !board.Contains(BoardElement{-2, 3, "user1", 12, nil}) {
		t.Fatalf("Board should contain -2,3,user1,12")
	}
	if !board.Contains(BoardElement{1, -3, "user2", 12, nil}) {
		t.Fatalf("Board should contain -2,3,user1,12")
	}
}

func TestNewBoard4(t *testing.T) {
	board := NewBoard([]string{"user1", "user2", "user3", "user4"})
	if len(board) != 4 {
		t.Fatalf("Board should contain 2 elements")
	}
	if !board.Contains(BoardElement{-2, 3, "user1", 12, nil}) {
		t.Fatalf("Board should contain -2,3,user1,12")
	}
	if !board.Contains(BoardElement{1, 3, "user2", 12, nil}) {
		t.Fatalf("Board should contain 1,3,user2,12")
	}
	if !board.Contains(BoardElement{1, -3, "user3", 12, nil}) {
		t.Fatalf("Board should contain 1,-3,user3,12")
	}
	if !board.Contains(BoardElement{-2, -3, "user4", 12, nil}) {
		t.Fatalf("Board should contain -2,-3,user4,12")
	}
}

func TestUpdate1(t *testing.T) {
	board := NewBoard([]string{"user1", "user2", "user3", "user4"})
	board.Update([]int{6, 5, 6, 5, 4}, &myfirestore.Click{-2, 3}, "user1")
	if len(board) != 7 {
		t.Fatalf("Board should contain 7 elements but contains %d\n", len(board))
	}
	pts := 26
	if !board.Contains(BoardElement{-2, 3, "user1", 12, &pts}) {
		t.Fatalf("Board should contain -2,3,user1,12,26")
	}
	if !board.Contains(BoardElement{1, 3, "user2", 12, nil}) {
		t.Fatalf("Board should contain 1,3,user2,12")
	}
	if !board.Contains(BoardElement{1, -3, "user3", 12, nil}) {
		t.Fatalf("Board should contain 1,-3,user3,12")
	}
	if !board.Contains(BoardElement{-2, -3, "user4", 12, nil}) {
		t.Fatalf("Board should contain -2,-3,user4,12")
	}
}

func TestRemove(t *testing.T) {
	board := NewBoard([]string{"user1", "user2"})
	if len(board) != 2 {
		t.Fatalf("Board should contain 2 elements")
	}
	if !board.Contains(BoardElement{-2, 3, "user1", 12, nil}) {
		t.Fatalf("Board should contain -2,3,user1,12")
	}
	if !board.Contains(BoardElement{1, -3, "user2", 12, nil}) {
		t.Fatalf("Board should contain -2,3,user1,12")
	}
	board.Remove(-2, 3, "user1")
	if len(board) != 1 {
		t.Fatalf("Board should contain 1 element but contains %d", len(board))
	}
	if !board.Contains(BoardElement{1, -3, "user2", 12, nil}) {
		t.Fatalf("Board should contain -2,3,user1,12")
	}
	board.Remove(1, -3, "xxx")
	if len(board) != 1 {
		t.Fatalf("Board should contain 1 element but contains %d", len(board))
	}
	board.Remove(1, -3, "user2")
	if len(board) != 0 {
		t.Fatalf("Board should contain 0 element but contains %d", len(board))
	}
}

func TestGetFormulasList(t *testing.T) {
	board := NewBoard([]string{"user1", "user2"})
	if len(board) != 2 {
		t.Fatalf("Board should contain 2 elements")
	}
	if len(board.GetFormulasList("user1")) != 1 {
		t.Fatalf("Should have 1 formula for user1, found %d", len(board.GetFormulasList("xxx")))
	}
	if len(board.GetFormulasList("xxx")) != 0 {
		t.Fatalf("Should have 0 formula for xxx, found %d", len(board.GetFormulasList("xxx")))
	}
}

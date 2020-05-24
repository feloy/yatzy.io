package game

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"

	myfirestore "github.com/feloy/yatzy.io/backend/lib/firestore"
	"github.com/feloy/yatzy.io/backend/lib/yatzy"
)

type BoardElement struct {
	X       int    `json:"x"`
	Y       int    `json:"y"`
	UserID  string `json:"userId"`
	Formula int    `json:"formula"`
	Points  *int   `json:"points,omitempty"`
}

var radius = 3

var startPositions = [][][]int{
	{},
	{ // 1 player
		{0, 0},
	},
	{ // 2 players
		{-2, 3},
		{1, -3},
	},
	{ // 3 players
		{-2, 3},
		{3, 0},
		{-2, -3},
	},
	{ // 4 players
		{-2, 3},
		{1, 3},
		{1, -3},
		{-2, -3},
	},
	{},
	{ // 6 players
		{-2, 3},
		{1, 3},
		{3, 0},
		{1, -3},
		{-2, -3},
		{-3, 0},
	},
}

type Move struct {
	X int
	Y int
}

var neighbours [][]Move = [][]Move{
	{
		{-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {-1, -1}, {0, -1},
	},
	{
		{-1, 0}, {1, 0}, {0, 1}, {1, 1}, {0, -1}, {1, -1},
	},
}

type Board []BoardElement

func newBoardElement(positions []int, userID string, formula int) BoardElement {
	element := BoardElement{
		X:      positions[0],
		Y:      positions[1],
		UserID: userID,
	}
	if formula == -1 {
		element.Formula = rand.Intn(13)
	} else {
		element.Formula = formula
	}
	return element
}

func NewBoard(userIDs []string) Board {
	fmt.Printf("Creating board for %d players: %+v\n", len(userIDs), userIDs)
	board := make([]BoardElement, 0, 37)
	positions := startPositions[len(userIDs)]
	for i, userID := range userIDs {
		board = append(board, newBoardElement(positions[i], userID, 12))
	}
	return board
}

func (board *Board) Update(die []int, pos *myfirestore.Click, userID string) {
	done := false
	var points = 0
	for i, cell := range *board {
		if pos.X == cell.X && pos.Y == cell.Y && userID == cell.UserID {
			points = yatzy.ComputePoints(cell.Formula, die)
			if points > 0 {
				(*board)[i].Points = &points
				done = true
				if cell.Formula < 6 && points < 3*(cell.Formula+1) {
					done = false
				}
			}
			break
		}
	}
	if done {
		abs := func(x int) int {
			if x < 0 {
				return -x
			}
			return x
		}

		for _, n := range neighbours[abs(pos.Y)%2] {
			x := pos.X + n.X
			y := pos.Y + n.Y
			if -radius <= y && y <= radius {
				if (-radius+int(math.Floor(float64(abs(y))/2))) <= x && x <= (radius-int(math.Floor((float64(abs(y))+1)/2))) {
					if !board.PositionOccupied(x, y) {
						*board = append(*board, newBoardElement([]int{x, y}, userID, -1))
					}
				}
			}
		}
	} else if points == 0 {
		board.Remove(pos.X, pos.Y, userID)
		// TODO
	}
}

func (board Board) CanReplay(userID string) bool {
	for _, cell := range board {
		if cell.Points == nil && cell.UserID == userID {
			return true
		}
	}
	return false
}

func (board Board) Contains(element BoardElement) bool {
	for _, el := range board {
		if reflect.DeepEqual(el, element) {
			return true
		}
	}
	return false
}

func (board Board) PositionOccupied(x, y int) bool {
	for _, el := range board {
		if el.X == x && el.Y == y {
			return true
		}
	}
	return false
}

func (board *Board) Remove(x, y int, userID string) {
	found := -1
	for i, el := range *board {
		if el.X == x && el.Y == y && el.UserID == userID {
			found = i
			break
		}
	}
	if found >= 0 {
		*board = append((*board)[:found], (*board)[found+1:]...)
	}
}

// GetFormulasList returns the list of possible formulas in the board for a specific user
func (board *Board) GetFormulasList(userID string) []int {
	list := []int{}
	for _, el := range *board {
		if el.UserID == userID && el.Points == nil {
			list = append(list, el.Formula)
		}
	}
	return list
}

func (board *Board) GetFormulaBestPosition(formula int, userID string) (found bool, x, y int) {
	for _, el := range *board {
		if el.Formula == formula && el.UserID == userID && el.Points == nil {
			return true, el.X, el.Y
		}
	}
	return false, 0, 0
}

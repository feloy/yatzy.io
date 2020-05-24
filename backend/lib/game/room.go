package game

import (
	"time"

	myfirestore "github.com/feloy/yatzy.io/backend/lib/firestore"
)

type RoomPlayer struct {
	Bot  bool
	Name string
}

type Room struct {
	RoomSize  int
	StartTime time.Time
	Board     *string
	NPlayers  *int
}

// NewRoomFromFirestoreValue returns a new Room, from values in FirestoreValue
func NewRoomFromFirestoreValue(v myfirestore.FirestoreValue) (*Room, error) {
	var room Room
	roomSize, err := v.GetIntegerValue("roomSize")
	if err != nil {
		return nil, err
	}
	startTime, err := v.GetTimestampValue("startTime")
	if err != nil {
		return nil, err
	}
	var board *string
	b, err := v.GetStringValue("board")
	if err == nil {
		board = &b
	}
	var nplayers *int
	n, err := v.GetIntegerValue("nplayers")
	if err == nil {
		nplayers = &n
	}
	room = Room{
		RoomSize:  roomSize,
		StartTime: startTime,
		Board:     board,
		NPlayers:  nplayers,
	}
	return &room, nil
}

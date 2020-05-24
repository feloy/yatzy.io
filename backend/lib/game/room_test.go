package game

import (
	"testing"

	myfirestore "github.com/feloy/yatzy.io/backend/lib/firestore"
)

func TestNewRoomFromFirestoreValue_Creation(t *testing.T) {
	v := myfirestore.FirestoreValue{
		Fields: map[string]interface{}{
			"roomSize": map[string]interface{}{
				"integerValue": "2",
			},
			"startTime": map[string]interface{}{
				"timestampValue": "2019-02-10T09:31:02.323743Z",
			},
		},
	}

	room, err := NewRoomFromFirestoreValue(v)
	if err != nil {
		t.Fatalf("Error %s\n", err)
	}
	if room.RoomSize != 2 {
		t.Fatalf("RoomSize should be '2' but is '%d'\n", room.RoomSize)
	}
	if room.StartTime.Year() != 2019 {
		t.Fatalf("Year of startTime should be '2019' but is '%d'\n", room.StartTime.Year())
	}
	if room.Board != nil {
		t.Fatalf("Board should be nil but is '%s'\n", *room.Board)
	}
	if room.NPlayers != nil {
		t.Fatalf("NPlayers should be nil but is '%d'\n", *room.NPlayers)
	}
}

func TestNewRoomFromFirestoreValue_UpdateNPlayers(t *testing.T) {
	v := myfirestore.FirestoreValue{
		Fields: map[string]interface{}{
			"roomSize": map[string]interface{}{
				"integerValue": "2",
			},
			"startTime": map[string]interface{}{
				"timestampValue": "2019-02-10T09:31:02.323743Z",
			},
			"nplayers": map[string]interface{}{
				"integerValue": "4",
			},
		},
	}

	room, err := NewRoomFromFirestoreValue(v)
	if err != nil {
		t.Fatalf("Error %s\n", err)
	}
	if room.RoomSize != 2 {
		t.Fatalf("RoomSize should be '2' but is '%d'\n", room.RoomSize)
	}
	if room.StartTime.Year() != 2019 {
		t.Fatalf("Year of startTime should be '2019' but is '%d'\n", room.StartTime.Year())
	}
	if room.Board != nil {
		t.Fatalf("Board should be nil but is '%s'\n", *room.Board)
	}
	if *room.NPlayers != 4 {
		t.Fatalf("NPlayers should be '4' but is '%d'\n", *room.NPlayers)
	}
}

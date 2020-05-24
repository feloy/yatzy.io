package game

import (
	"reflect"
	"testing"

	myfirestore "github.com/feloy/yatzy.io/backend/lib/firestore"
)

func TestNewUserFromFirestoreValue_NewHumanUser(t *testing.T) {
	v := myfirestore.FirestoreValue{
		Fields: map[string]interface{}{
			"name": map[string]interface{}{
				"stringValue": "my name",
			},
			"size": map[string]interface{}{
				"integerValue": "2",
			},
			"tokenId": map[string]interface{}{
				"stringValue": "AZE123",
			},
		},
	}

	user, err := NewUserFromFirestoreValue(v)
	if err != nil {
		t.Fatalf("Error %s\n", err)
	}
	if user.Name != "my name" {
		t.Fatalf("Name should be 'my name' but is '%s'\n", user.Name)
	}
	if *user.Size != 2 {
		t.Fatalf("Size should be '2' but is '%d'\n", *user.Size)
	}
	if *user.TokenID != "AZE123" {
		t.Fatalf("TokenID should be 'AZE123' but is '%s'\n", *user.TokenID)
	}
	if user.Room != nil {
		t.Fatalf("Room should be nil but is '%s'\n", *user.Room)
	}
}

func TestNewUserFromFirestoreValue_NewBotUser(t *testing.T) {
	v := myfirestore.FirestoreValue{
		Fields: map[string]interface{}{
			"name": map[string]interface{}{
				"stringValue": "bot 1",
			},
			"room": map[string]interface{}{
				"stringValue": "azerty",
			},
		},
	}

	user, err := NewUserFromFirestoreValue(v)
	if err != nil {
		t.Fatalf("Error %s\n", err)
	}
	if user.Name != "bot 1" {
		t.Fatalf("Name should be 'bot 1' but is '%s'\n", user.Name)
	}
	if user.Size != nil {
		t.Fatalf("Size should be nil but is '%d'\n", *user.Size)
	}
	if user.TokenID != nil {
		t.Fatalf("TokenID should be nil but is '%s'\n", *user.TokenID)
	}
	if *user.Room != "azerty" {
		t.Fatalf("Room should be 'azerty' but is '%s'\n", *user.Room)
	}
}

func TestNewUserFromFirestoreValue_ReplayDie(t *testing.T) {
	v := myfirestore.FirestoreValue{
		Fields: map[string]interface{}{
			"name": map[string]interface{}{
				"stringValue": "bot 1",
			},
			"shots": map[string]interface{}{
				"integerValue": "2",
			},
			"die": map[string]interface{}{
				"arrayValue": map[string]interface{}{
					"values": []interface{}{
						map[string]interface{}{"integerValue": "3"},
						map[string]interface{}{"integerValue": "2"},
						map[string]interface{}{"integerValue": "1"},
						map[string]interface{}{"integerValue": "4"},
						map[string]interface{}{"integerValue": "5"},
					},
				},
			},
			"replay": map[string]interface{}{
				"arrayValue": map[string]interface{}{
					"values": []interface{}{
						map[string]interface{}{"integerValue": "2"},
						map[string]interface{}{"integerValue": "3"},
						map[string]interface{}{"integerValue": "4"},
					},
				},
			},
		},
	}

	user, err := NewUserFromFirestoreValue(v)
	if err != nil {
		t.Fatalf("Error %s\n", err)
	}
	if user.Name != "bot 1" {
		t.Fatalf("Name should be 'bot 1' but is '%s'\n", user.Name)
	}
	if *user.Shots != 2 {
		t.Fatalf("Shots should be '2' but is '%d'\n", *user.Shots)
	}
	if !reflect.DeepEqual(user.Die, []int{3, 2, 1, 4, 5}) {
		t.Fatalf("Die should be [3,2,1,4,5] but is %v\n", user.Die)
	}
	if !reflect.DeepEqual(user.Replay, []int{2, 3, 4}) {
		t.Fatalf("Replay should be [2,3,4] but is %v\n", user.Replay)
	}
}
func TestNewUserFromFirestoreValue_Click(t *testing.T) {
	v := myfirestore.FirestoreValue{
		Fields: map[string]interface{}{
			"name": map[string]interface{}{
				"stringValue": "bot 1",
			},
			"shots": map[string]interface{}{
				"integerValue": "2",
			},
			"die": map[string]interface{}{
				"arrayValue": map[string]interface{}{
					"values": []interface{}{
						map[string]interface{}{"integerValue": "3"},
						map[string]interface{}{"integerValue": "2"},
						map[string]interface{}{"integerValue": "1"},
						map[string]interface{}{"integerValue": "4"},
						map[string]interface{}{"integerValue": "5"},
					},
				},
			},
			"click": map[string]interface{}{
				"mapValue": map[string]interface{}{
					"fields": map[string]interface{}{
						"x": map[string]interface{}{
							"integerValue": "1",
						},
						"y": map[string]interface{}{
							"integerValue": "-3",
						},
					},
				},
			},
			// click:map[mapValue:map[fields:map[x:map[integerValue:1] y:map[integerValue:-3]]]]
		},
	}

	user, err := NewUserFromFirestoreValue(v)
	if err != nil {
		t.Fatalf("Error %s\n", err)
	}
	if user.Name != "bot 1" {
		t.Fatalf("Name should be 'bot 1' but is '%s'\n", user.Name)
	}
	if *user.Shots != 2 {
		t.Fatalf("Shots should be '2' but is '%d'\n", *user.Shots)
	}
	if !reflect.DeepEqual(user.Die, []int{3, 2, 1, 4, 5}) {
		t.Fatalf("Die should be [3,2,1,4,5] but is %v\n", user.Die)
	}
	if !reflect.DeepEqual(user.Click, &myfirestore.Click{X: 1, Y: -3}) {
		t.Fatalf("Click should be [x:1,y:-3] but is %v\n", user.Click)
	}
}

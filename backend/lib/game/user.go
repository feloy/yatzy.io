package game

import (
	myfirestore "github.com/feloy/yatzy.io/backend/lib/firestore"
)

type User struct {
	Name    string
	Size    *int
	TokenID *string
	Room    *string
	Shots   *int
	Die     []int
	Replay  []int
	Click   *myfirestore.Click
	Finish  *bool
}

// NewUserFromFirestoreValue returns a new Score, from values in FirestoreValue
func NewUserFromFirestoreValue(v myfirestore.FirestoreValue) (*User, error) {
	var user User
	name, err := v.GetStringValue("name")
	if err != nil {
		return nil, err
	}
	var size *int
	s, err := v.GetIntegerValue("size")
	if err == nil {
		size = &s
	}
	var tokenID *string
	t, err := v.GetStringValue("tokenId")
	if err == nil {
		tokenID = &t
	}
	var room *string
	r, err := v.GetStringValue("room")
	if err == nil {
		room = &r
	}
	var shots *int
	sh, err := v.GetIntegerValue("shots")
	if err == nil {
		shots = &sh
	}
	die, err := v.GetIntArrayValue("die")
	if err != nil {
		return nil, err
	}
	replay, err := v.GetIntArrayValue("replay")
	if err != nil {
		return nil, err
	}
	click, _ := v.GetClickValue("click")
	user = User{
		Name:    name,
		Size:    size,
		TokenID: tokenID,
		Shots:   shots,
		Room:    room,
		Die:     die,
		Replay:  replay,
		Click:   click,
	}
	return &user, nil
}

func (o *User) IsHuman() bool {
	return o.Size != nil
}

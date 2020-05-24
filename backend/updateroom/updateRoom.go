package p

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/functions/metadata"
	"google.golang.org/api/iterator"

	myfirestore "github.com/feloy/yatzy.io/backend/lib/firestore"
	"github.com/feloy/yatzy.io/backend/lib/game"
)

var (
	// The project ID, set by the user
	projectID = os.Getenv("PROJECT_ID")
)

// FirestoreEvent is the payload of a Firestore event.
// Please refer to the docs for additional information
// regarding Firestore events.
type FirestoreEvent struct {
	OldValue myfirestore.FirestoreValue `json:"oldValue"`
	Value    myfirestore.FirestoreValue `json:"value"`
}

// OnUpdateRoom is triggered when a room is updated
// - Initiates a board when the room is full
func OnUpdateRoom(ctx context.Context, e FirestoreEvent) error {
	meta, err := metadata.FromContext(ctx)
	if err != nil {
		return fmt.Errorf("metadata.FromContext: %v", err)
	}
	log.Printf("Meta resource: %#v", meta.Resource)

	// Get data from Room
	room, err := game.NewRoomFromFirestoreValue(e.Value)
	if err != nil {
		return err
	}
	log.Printf("room: %+v", room)

	if room.RoomSize == *room.NPlayers && room.Board == nil {
		log.Print("Room is ready")
		client, err := firestore.NewClient(ctx, projectID)
		if err != nil {
			return err
		}
		defer client.Close()

		roomID := myfirestore.GetResourceIDFromMeta(*meta)

		return client.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
			ids, err := GetPlayersIDsInRoom(ctx, tx, client, roomID)
			if err != nil {
				return err
			}
			// Create board
			board := game.NewBoard(ids)
			bytes, err := json.Marshal(board)
			if err != nil {
				return err
			}
			ref := client.Collection("rooms").Doc(roomID)
			err = tx.Set(ref, map[string]interface{}{
				"board": string(bytes),
			}, firestore.MergeAll)
			if err != nil {
				return err
			}
			return sendDie(ctx, tx, client, ids)
		})
	}

	log.Print("Room is not ready. Bye")
	return nil
}

func GetPlayersIDsInRoom(ctx context.Context, tx *firestore.Transaction, client *firestore.Client, roomID string) ([]string, error) {
	query := client.Collection("rooms").Doc(roomID).Collection("players")
	iter := tx.Documents(query)
	ids := []string{}
	for {
		o, err := iter.Next()

		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		ids = append(ids, o.Ref.ID)
	}
	return ids, nil
}

func sendDie(ctx context.Context, tx *firestore.Transaction, client *firestore.Client, users []string) error {
	for _, user := range users {
		ref := client.Collection("users").Doc(user)
		err := tx.Set(ref, map[string]interface{}{
			"die": []int{
				getRandomDice(),
				getRandomDice(),
				getRandomDice(),
				getRandomDice(),
				getRandomDice(),
			},
			"shots": 2,
		}, firestore.MergeAll)
		if err != nil {
			return err
		}
	}
	return nil
}

func getRandomDice() int {
	return 1 + rand.Intn(6)
}

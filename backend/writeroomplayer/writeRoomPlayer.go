package p

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/functions/metadata"
	myfirestore "github.com/feloy/yatzy.io/backend/lib/firestore"
	"google.golang.org/api/iterator"
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

// OnWriteRoomPlayer is triggered when a new player is added to or removed from a room
// Counts the total number of players in the room and updates the `players` value in the room.
func OnWriteRoomPlayer(ctx context.Context, e FirestoreEvent) error {
	meta, err := metadata.FromContext(ctx)
	if err != nil {
		return fmt.Errorf("metadata.FromContext: %v", err)
	}

	roomID := myfirestore.GetParentIDFromMeta(*meta)

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()

	return client.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		count, err := GetPlayersCountInRoom(ctx, tx, client, roomID)
		if err != nil {
			return err
		}
		ref := client.Collection("rooms").Doc(roomID)
		err = tx.Set(ref, map[string]interface{}{
			"nplayers": count,
		}, firestore.MergeAll)

		return err
	})
}

func GetPlayersCountInRoom(ctx context.Context, tx *firestore.Transaction, client *firestore.Client, roomID string) (int, error) {
	query := client.Collection("rooms").Doc(roomID).Collection("players")
	iter := tx.Documents(query)
	count := 0
	for {
		_, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return 0, err
		}
		count++
	}
	return count, nil
}

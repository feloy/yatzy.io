package newuser

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/functions/metadata"

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

// OnNewUser is triggered when a new user is created
func OnNewUser(ctx context.Context, e FirestoreEvent) error {
	log.Printf("projectID: %s", projectID)

	meta, err := metadata.FromContext(ctx)
	if err != nil {
		return fmt.Errorf("metadata.FromContext: %v", err)
	}
	log.Printf("Meta resource: %#v", meta.Resource)

	// Get data from User
	user, err := game.NewUserFromFirestoreValue(e.Value)
	if err != nil {
		return err
	}
	log.Printf("user: %+v", user)

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Printf("error getting firestore client")
		return err
	}
	defer client.Close()

	// Get User ID from URL
	userID := myfirestore.GetResourceIDFromMeta(*meta)
	log.Printf("userID: %s", userID)

	if user.IsHuman() {
		// Verify token
		app, err := firebase.NewApp(context.Background(), nil)
		if err != nil {
			log.Printf("error initializing app: %v\n", err)
			return err
		}
		auth, err := app.Auth(ctx)
		if err != nil {
			log.Printf("error auth: %v\n", err)
			return err
		}

		token, err := auth.VerifyIDToken(ctx, *user.TokenID)
		if err != nil {
			log.Printf("error decoding token\n")
			return err
		}

		return client.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
			var avail bool
			var roomID string

			if *user.BotsInvites > 0 {
				// Force new game when inviting bots
				avail = false
			} else {
				avail, roomID, err = getAvailableRoom(ctx, tx, client, *user.Size)
				if err != nil {
					log.Printf("Error getting if available room\n")
					return err
				}
			}

			if avail {
				err = addUserToExistingRoom(ctx, tx, client, roomID, userID, user.Name, false)
				if err != nil {
					log.Printf("error updating user information")
					return err
				}
			} else {
				// Create a new room and add user to it
				roomID, err = addUserToNewRoom(ctx, tx, client, userID, user.Name, *user.Size, meta.Timestamp)
				if err != nil {
					log.Printf("error adding user to new room")
					return err
				}
				if len(roomID) == 0 {
					log.Printf("Error creating new room\n")
					return err
				}
			}

			// Update user info
			ref := client.Collection("users").Doc(userID)
			err = tx.Set(ref, map[string]interface{}{
				"room":    roomID,
				"userId":  token.UID,
				"tokenId": "",
			}, firestore.MergeAll)
			if err != nil {
				log.Printf("error updating user information")
				return err
			}

			// Create companion Bot users
			for i := 1; i <= *user.BotsInvites; i++ {
				err = CreateBotUser(ctx, tx, client, i, roomID)
				if err != nil {
					log.Printf("error creating bot user")
					return err
				}
			}
			return nil
		})
	}

	// A bot has been created

	return client.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		return addUserToExistingRoom(ctx, tx, client, *user.Room, userID, user.Name, true)
	})
}

// CreateBotUser creates a new Bot user and affects it to the given room
func CreateBotUser(ctx context.Context, tx *firestore.Transaction, client *firestore.Client, i int, roomID string) error {
	ref := client.Collection("users").NewDoc()
	err := tx.Create(ref, map[string]interface{}{
		"name": fmt.Sprintf("bot %d", i),
		"room": roomID,
	})
	if err != nil {
		log.Println("Error creating new bot user")
	}
	return err
}

// addUserToNewRoom creates a new room and adds a user to it
func addUserToNewRoom(ctx context.Context, tx *firestore.Transaction, client *firestore.Client, userID string, name string, roomSize int, startTime time.Time) (string, error) {
	// Create room
	roomRef := client.Collection("rooms").NewDoc()
	err := tx.Create(roomRef, map[string]interface{}{

		"roomSize":  roomSize,
		"startTime": startTime,
	})
	if err != nil {
		log.Println("Error creating new room")
		return "", err
	}

	log.Printf("Added room ID=%s\n", roomRef.ID)

	data := map[string]interface{}{
		"name": name,
		"bot":  false,
	}
	log.Printf("data=%s\n", data)

	// Add player sub-collection
	ref := client.
		Collection("rooms").Doc(roomRef.ID).
		Collection("players").Doc(userID)
	err = tx.Set(ref, data)
	if err != nil {
		log.Println("Error creating new player in room")
		return "", err
	}
	log.Println("Added new player in room")
	return roomRef.ID, nil
}

// addUserToExistingNonFullRoom adds a user to an existing non full room
func addUserToExistingRoom(ctx context.Context, tx *firestore.Transaction, client *firestore.Client,
	room string, userID string, name string, bot bool) error {
	// Add player sub-collection
	data := map[string]interface{}{
		"name": name,
		"bot":  bot,
	}

	ref := client.
		Collection("rooms").Doc(room).
		Collection("players").Doc(userID)
	err := tx.Set(ref, data)
	if err != nil {
		log.Println("Error adding new player in room")
		return err
	}
	log.Println("Added new player in room")
	return nil
}

func getAvailableRoom(ctx context.Context, tx *firestore.Transaction, client *firestore.Client, size int) (bool, string, error) {
	query := client.Collection("rooms").Where("roomSize", "==", size).Where("nplayers", "<", size)
	iter := tx.Documents(query)
	next, err := iter.Next()
	if err == iterator.Done {
		return false, "", nil // No room available
	}
	if err != nil {
		return false, "", err // error occured
	}

	log.Printf("Found available room with size %d: %s", size, next.Ref.ID)
	return true, next.Ref.ID, nil
}

package p

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/functions/metadata"

	myfirestore "github.com/feloy/yatzy.io/backend/lib/firestore"
	"github.com/feloy/yatzy.io/backend/lib/game"
	"github.com/feloy/yatzy.io/backend/lib/yatzy"
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

// OnUpdateUser is triggered when a user is updated
// - Sends new die when replay is queried
func OnUpdateUser(ctx context.Context, e FirestoreEvent) error {
	meta, err := metadata.FromContext(ctx)
	if err != nil {
		return fmt.Errorf("metadata.FromContext: %v", err)
	}
	//	log.Printf("Function triggered by change to: %v", meta.Resource)
	//	log.Printf("%+v", e.Value)

	userID := myfirestore.GetResourceIDFromMeta(*meta)

	// Get data from User
	user, err := game.NewUserFromFirestoreValue(e.Value)
	if err != nil {
		return err
	}
	//	log.Printf("%+v", user)

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()

	if len(user.Replay) > 0 && *user.Shots > 0 {
		*user.Shots--
		for _, r := range user.Replay {
			user.Die[r] = getRandomDice()
		}
		err = client.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
			ref := client.Collection("users").Doc(userID)
			tx.Set(ref, map[string]interface{}{
				"shots":  *user.Shots,
				"die":    user.Die,
				"replay": firestore.Delete, // Stop recursivity
			}, firestore.MergeAll)

			return nil
		})
	} else if user.Click != nil {
		//		fmt.Printf("%+v\n", user.Click)
		err = client.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
			canReplay := updateBoard(ctx, tx, client, userID, user)
			ref := client.Collection("users").Doc(userID)
			if canReplay {
				tx.Set(ref, map[string]interface{}{
					"click": firestore.Delete, // Stop recursivity
					"die": []int{
						getRandomDice(),
						getRandomDice(),
						getRandomDice(),
						getRandomDice(),
						getRandomDice(),
					},
					"shots": 2,
				}, firestore.MergeAll)
			} else {
				tx.Set(ref, map[string]interface{}{
					"click":  firestore.Delete, // Stop recursivity
					"finish": true,
					// endTime
				}, firestore.MergeAll)
			}
			return nil
		})
	} else if user.TokenID == nil && len(user.Replay) == 0 && user.Click == nil {
		// Bot must play
		//		log.Printf("%s will play", user.Name)
		err = client.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
			if *user.Shots == 0 {
				ChooseFormula(ctx, tx, client, userID, user)
			} else {
				ReplayDie(ctx, tx, client, userID)
			}
			return nil
		})
	}
	return nil
}

func GetBoard(ctx context.Context, tx *firestore.Transaction, client *firestore.Client, user *game.User) game.Board {
	ref := client.Collection("rooms").Doc(*user.Room)
	roomSnap, err := tx.Get(ref)
	if err != nil {
		//fmt.Printf("Unable to get room of player\n")
		return nil
	}
	room := roomSnap.Data()
	board := game.Board{}
	json.Unmarshal([]byte(room["board"].(string)), &board)
	return board
}

func updateBoard(ctx context.Context, tx *firestore.Transaction, client *firestore.Client, userID string, user *game.User) bool {
	die := user.Die
	pos := user.Click

	board := GetBoard(ctx, tx, client, user)
	board.Update(die, pos, userID)
	bytes, err := json.Marshal(board)
	if err != nil {
		//fmt.Println("Error marshalling board")
		return false
	}
	ref := client.Collection("rooms").Doc(*user.Room)
	tx.Set(ref, map[string]interface{}{
		"board": string(bytes),
	}, firestore.MergeAll)
	return board.CanReplay(userID)
}

// BOT

func ChooseFormula(ctx context.Context, tx *firestore.Transaction, client *firestore.Client, userID string, user *game.User) {
	board := GetBoard(ctx, tx, client, user)
	formula := GetBestPointsFormula(user.Die, board.GetFormulasList(userID))
	found, x, y := board.GetFormulaBestPosition(formula, userID)
	//fmt.Printf("Best formula: %d - found = %v", formula, found)
	if found {
		//		log.Printf("%s will play in (%d, %d)", user.Name, x, y)
		ref := client.Collection("users").Doc(userID)
		time.Sleep(time.Duration(2+rand.Intn(2)) * time.Second)
		tx.Set(ref, map[string]interface{}{
			"click": map[string]int{
				"x": x,
				"y": y,
			},
		}, firestore.MergeAll)
	} else {
		//		log.Printf("%s cannot play. Bye", user.Name)
	}
}

func ReplayDie(ctx context.Context, tx *firestore.Transaction, client *firestore.Client, userID string) {
	time.Sleep(5 * time.Second)
	ref := client.Collection("users").Doc(userID)
	tx.Set(ref, map[string]interface{}{
		"replay": []int{0, 2, 4},
	}, firestore.MergeAll)
}

// GetBestPointsFormula returns the formula with a max of points, with the provided die
func GetBestPointsFormula(die []int, formulas []int) int {
	bestPoints := -1
	bestFormula := -1
	for _, formula := range formulas {
		points := yatzy.ComputePoints(formula, die)
		if points > bestPoints {
			bestPoints = points
			bestFormula = formula
		}
	}
	return bestFormula
}

func getRandomDice() int {
	return 1 + rand.Intn(6)
}

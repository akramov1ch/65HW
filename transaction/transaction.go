package transaction

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()


func UpdateUserScore(client *redis.Client, userID string, score int) {	
	err := client.Watch(ctx, func(tx *redis.Tx) error{

		currentScore, err := tx.ZScore(ctx, "user_scores", userID).Result()
		if err != nil {
			return err
		}

		newScore := currentScore + float64(score)

		_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			pipe.ZAdd(ctx, "user_scores", redis.Z{
				Score: newScore,
				Member: userID,
			})
			return nil
		})
		return err
	}, "user_scores")


	if err!= nil {
		log.Fatalf("Failed to update user score: %v", err)
	} else {
		fmt.Printf("User %s score successfully updated by %d points.\n", userID, score)
	}
}

package score

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func AddUserScore(client *redis.Client, userID string, score int) {
	fmt.Println("Adding user")
	err := client.ZAdd(ctx, "user_scores", redis.Z{Score: float64(score), Member: userID}).Err()
	if err != nil {
		panic(err)
	}
	items, err := client.ZRangeWithScores(ctx, "user_scores", 0, -1).Result()
	if err != nil {
		log.Fatalf("Failed to get items from sorted set: %v", err)
	}

	for _, item := range items {
		fmt.Printf("Member: %s, Score: %.2f\n", item.Member, item.Score)
	}

}

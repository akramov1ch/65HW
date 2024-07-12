package display

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()


func DisplayLeaderboard(client *redis.Client, n int) {
	fmt.Println("Leaderboard:")
	leaderboard, err := client.ZRevRangeWithScores(ctx, "user_scores", 0, int64(n - 1)).Result()
	if err != nil {
		log.Println("Error retrieving leaderboard: ", err)
	}

	for rk ,m := range leaderboard {
		fmt.Printf("Rank: %d: %s - %.2f points\n", rk + 1, m.Member, m.Score)
	}

}
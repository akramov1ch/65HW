package score

import (
	"math/rand"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func AddMultipleUsers(client *redis.Client, n int) {

	for i := 0; i < n; i++ {
		userID := "user" + strconv.Itoa(i+1)
		score := rand.Intn(10)
		AddUserScore(client, userID, score)
	}
}

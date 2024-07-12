package main

import (
	"context"
	"fmt"
	"65HW/display"
	"65HW/score"
	"65HW/transaction"
	"log"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	defer client.Close()

	pong, err := client.Ping(ctx).Result()
	
	if err != nil {
		log.Printf("Error: %v", err)
	}
	
	fmt.Println(pong)

	score.AddMultipleUsers(client, 1)
	display.DisplayLeaderboard(client, 5)

	transaction.UpdateUserScore(client, "user1", 5)
}

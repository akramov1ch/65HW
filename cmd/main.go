package main

import (
	"context"
	"fmt"
	disp "65HW/internal/display"
	sc "65HW/internal/score"
	tr "65HW/internal/transaction"
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

	sc.AddMultipleUsers(client, 1)
	disp.DisplayLeaderboard(client, 5)

	tr.UpdateUserScore(client, "user1", 5)
}


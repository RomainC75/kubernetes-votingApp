package main

import (
	"context"
	"fmt"
	redis_client "job/redis"
	"os"
	db "shared/postgres/sqlc"

	"github.com/sirupsen/logrus"
)

func main() {
	values := os.Environ()
	for index, v := range values {
		fmt.Printf("%d : %s\n", index, v)
	}

	redisVoteKey := os.Getenv("REDIS_VOTE_KEY")

	db.Connect()
	conn := db.GetConnection()
	ctx := context.Background()
	redis_client.SetRedis()

	rdb := redis_client.GetRedis()

	votes, err := (*conn).ListVotes(ctx)
	if err != nil {
		logrus.Error("=> error trying to get list of votes", err.Error())
	}

	fmt.Println("->  ", votes)

	err = rdb.Set(ctx, redisVoteKey, len(votes), 0).Err()
	if err != nil {
		logrus.Error("error trying to set v in redis : ", err.Error())
	}
}

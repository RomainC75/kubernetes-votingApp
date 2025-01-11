package main

import (
	"context"
	"encoding/json"
	"fmt"
	redis_client "job/redis"
	"job/utils"
	"os"
	"shared/dtos"
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

	votesMap := utils.MapMaker(votes)
	votesResult := dtos.VotesResultsDto{
		Result: votesMap,
	}

	uJson, _ := json.Marshal(votesResult)

	err = rdb.Set(ctx, redisVoteKey, uJson, 0).Err()
	if err != nil {
		logrus.Error("error trying to set v in redis : ", err.Error())
	}
}

package main

import (
	"context"
	"fmt"
	"os"
	db "shared/postgres/sqlc"

	"github.com/sirupsen/logrus"
)

func main() {
	db.Connect()

	values := os.Environ()
	for index, v := range values {
		fmt.Printf("%d : %s\n", index, v)
	}

	conn := db.GetConnection()
	ctx := context.Background()

	votes, err := (*conn).ListVotes(ctx)
	if err != nil {
		logrus.Error("=> error trying to get list of votes", err.Error())
	}

	fmt.Println("->  ", votes)

}

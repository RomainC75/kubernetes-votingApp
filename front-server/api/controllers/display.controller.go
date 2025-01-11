package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	redis_client "front-server/api/redis"
	"front-server/utils"
	"html/template"
	"net/http"
	"os"

	"github.com/go-redis/redis"
)

type PageData struct {
	PageTitle  string
	VoteResult map[string]int
}

func Display(w http.ResponseWriter, r *http.Request) {
	redisKey := os.Getenv("REDIS_VOTE_KEY")
	tmpl := template.Must(template.ParseFiles("templates/layout.html"))
	ctx := context.Background()
	qtyStr, err := redis_client.GetRedis().Get(ctx, redisKey).Result()
	if err == redis.Nil {
		fmt.Println("--> key2 does not exist")
	} else if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"error": err.Error(),
		})
	} else {
		fmt.Println("key2")
	}

	fmt.Println("-> found key", qtyStr)

	voteResult, _ := utils.QuantityAnalysor(qtyStr)
	fmt.Println("-> voteResult : ", voteResult)

	data := PageData{
		PageTitle:  "My TODO listss",
		VoteResult: voteResult.Result,
	}
	tmpl.Execute(w, data)
}

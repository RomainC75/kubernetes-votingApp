package main

import (
	"context"
	"encoding/json"
	"fmt"
	redis_client "front-server/redis"
	"front-server/utils"
	"html/template"
	"net/http"
	"os"

	"github.com/go-redis/redis"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	Quantity  int
	PageTitle string
	Todos     []Todo
}

func main() {
	redis_client.SetRedis()
	// values := os.Environ()
	// for index, v := range values {
	// 	fmt.Printf("%d : %s\n", index, v)
	// }
	serverPort := os.Getenv("FRONT_SERVER_PORT")
	redisKey := os.Getenv("REDIS_VOTE_KEY")

	tmpl := template.Must(template.ParseFiles("templates/layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

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

		qty, err := utils.QuantityAnalysor(qtyStr)

		data := TodoPageData{
			Quantity:  qty,
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})

	fmt.Println("-> ", serverPort)
	http.ListenAndServe(fmt.Sprintf(":%s", serverPort), nil)
}

package api

import (
	"fmt"
	"front-server/api/bootstrap.go"
	"front-server/api/controllers"
	"net/http"
	"os"
)

func RunApi() {
	bootstrap.Bootstrap()

	http.HandleFunc("/", controllers.Display)

	serverPort := os.Getenv("FRONT_SERVER_PORT")
	fmt.Println("-> ", serverPort)
	http.ListenAndServe(fmt.Sprintf(":%s", serverPort), nil)
}

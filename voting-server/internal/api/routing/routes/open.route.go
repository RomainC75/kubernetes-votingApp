package routes

import (
	"net/http"
	"voting-server/internal/api/controllers"
)

func OpenRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /ping", controllers.PingController)
}

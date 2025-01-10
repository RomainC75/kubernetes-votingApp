package routes

import (
	"net/http"
	"voting-server/internal/api/controllers"
)

func VotesRoutes(mux *http.ServeMux) {
	voteCtrl := controllers.NewVoteController()
	mux.HandleFunc("POST /vote", voteCtrl.PostVoteController)
}

package controllers

import (
	"net/http"
	"voting-server/utils"
)

func PingController(w http.ResponseWriter, r *http.Request) {
	res := map[string]any{
		"message": "pong",
	}
	utils.SendJson(w, http.StatusOK, res)
}

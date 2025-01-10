package bootstrap

import (
	db "shared/postgres/sqlc"
	"voting-server/internal/api"
	"voting-server/internal/api/routing"
)

func Bootstrap() {
	db.Connect()

	mux := routing.ConnectRoutes()
	api.RunApi(mux)
}

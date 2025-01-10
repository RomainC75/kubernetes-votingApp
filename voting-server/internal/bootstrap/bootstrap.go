package bootstrap

import (
	validator_helper "shared/dtos/validator"
	db "shared/postgres/sqlc"
	"voting-server/internal/api"
	"voting-server/internal/api/routing"
)

func Bootstrap() {

	db.Connect()
	validator_helper.SetValidate()

	mux := routing.ConnectRoutes()
	api.RunApi(mux)
}

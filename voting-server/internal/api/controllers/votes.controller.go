package controllers

import (
	"encoding/json"
	"net/http"
	"shared/dtos"
	validator_helper "shared/dtos/validator"
	"voting-server/internal/api/services"
	"voting-server/utils"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type VoteController struct {
	VoteSrv *services.VoteSrv
	v       *validator.Validate
}

func NewVoteController() *VoteController {
	return &VoteController{
		VoteSrv: &services.VoteSrv{},
		v:       validator_helper.GetValidate(),
	}
}

func (vCtr *VoteController) PostVoteController(w http.ResponseWriter, r *http.Request) {
	var portVoteBody dtos.VoteDto

	err := json.NewDecoder(r.Body).Decode(&portVoteBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = vCtr.v.Struct(portVoteBody)
	if err != nil {
		logrus.Warnf("validator error : %s \n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := vCtr.VoteSrv.SetVote(portVoteBody.Content)
	utils.SendJson(w, http.StatusOK,
		map[string]any{
			"content": res,
		},
	)
}

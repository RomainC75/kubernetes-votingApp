package controllers

import (
	"encoding/json"
	"net/http"
	"shared/dtos"
	validator_helper "shared/dtos/validator"
	"voting-server/internal/api/services"
	"voting-server/utils"

	"github.com/go-playground/validator/v10"
)

type VoteController struct {
	voteSrv services.VoteSrvI
	v       *validator.Validate
}

func NewVoteController() *VoteController {
	return &VoteController{
		voteSrv: services.NewVoteSrv(),
		v:       validator_helper.GetValidate(),
	}
}

func (vCtr *VoteController) PostVoteController(w http.ResponseWriter, r *http.Request) {
	var newVoteBody dtos.VoteDto

	err := json.NewDecoder(r.Body).Decode(&newVoteBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = vCtr.v.Struct(newVoteBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := vCtr.voteSrv.SetVote(newVoteBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.SendJson(w, http.StatusOK,
		map[string]any{
			"content": res,
		},
	)
}

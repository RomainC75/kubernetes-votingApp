package services

import (
	"shared/dtos"
	db "shared/postgres/sqlc"
)

type VoteSrvI interface {
	SetVote(voteDto dtos.VoteDto) (db.Vote, error)
}

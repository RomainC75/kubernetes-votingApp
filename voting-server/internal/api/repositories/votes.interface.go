package repositories

import (
	"shared/dtos"
	db "shared/postgres/sqlc"
)

type VoteRepoI interface {
	SetNewVote(voteDto dtos.VoteDto) (db.Vote, error)
}

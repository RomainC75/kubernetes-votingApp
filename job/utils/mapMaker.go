package utils

import db "shared/postgres/sqlc"

func MapMaker(votes []db.Vote) map[string]int {
	res := map[string]int{}

	for _, vote := range votes {
		qty, ok := res[vote.Content]
		if ok {
			res[vote.Content] = qty + 1
		} else {
			res[vote.Content] = 1
		}
	}

	return res
}

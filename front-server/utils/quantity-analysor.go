package utils

import (
	"encoding/json"
	"shared/dtos"
)

func QuantityAnalysor(str string) (dtos.VotesResultsDto, error) {

	var data dtos.VotesResultsDto
	err := json.Unmarshal([]byte(str), &data)
	return data, err

}

package main

import (
	"fmt"
	"os"
	"voting-server/internal/bootstrap"
)

func main() {

	values := os.Environ()
	for index, v := range values {
		fmt.Printf("%d : %s\n", index, v)
	}

	bootstrap.Bootstrap()

}

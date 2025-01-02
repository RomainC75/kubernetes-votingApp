package main

import (
	"fmt"
	"os"
)

func main() {
	test := os.Getenv("voting-app-node-external-port")
	fmt.Printf("voting-app-node-external-port : %s \n", test)

	values := os.Environ()
	for index, v := range values {
		fmt.Printf("%d : %s\n", index, v)
	}
}

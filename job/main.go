package main

import (
	"fmt"
	"os"
)

func main() {
	values := os.Environ()
	for index, v := range values {
		fmt.Printf("%d : %s\n", index, v)
	}
}

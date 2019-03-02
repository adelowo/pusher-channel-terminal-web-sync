package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func main() {

	for {
		fmt.Printf("Generating a new UUID -- %s", uuid.New())
		fmt.Println()
		time.Sleep(time.Millisecond * 500)
	}
}

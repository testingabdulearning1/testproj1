package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func main() {
	for {
		id := uuid.New()
		fmt.Println("hello", id)
		time.Sleep(3 * time.Second)
	}
}
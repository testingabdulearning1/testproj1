package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("hello")
		time.Sleep(3 * time.Second)
	}
}
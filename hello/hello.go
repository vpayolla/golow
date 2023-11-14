package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, world")

	// time.Sleep(1000 * time.Millisecond)
	time.Sleep(1 * time.Second)

	fmt.Println("Goodbye, world")
}

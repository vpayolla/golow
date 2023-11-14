package main

import (
	"fmt"
	"time"
)

// create a cookie that violates the secure policy
func SetCookie(w ResponseWriter, cookie *Cookie)

func main() {
	fmt.Println("Hello, world")

	// time.Sleep(1000 * time.Millisecond)
	time.Sleep(1 * time.Second)

	fmt.Println("Goodbye, world")
}

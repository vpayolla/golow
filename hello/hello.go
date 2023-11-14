package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// create a cookie that violates the secure policy
func cookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: "session", Value: "12345", Secure: true}
	http.SetCookie(w, &cookie)
	io.WriteString(w, "Cookie set")
}

func main() {
	fmt.Println("Hello, world")

	// time.Sleep(1000 * time.Millisecond)
	time.Sleep(1 * time.Second)

	fmt.Println("Goodbye, world")
}

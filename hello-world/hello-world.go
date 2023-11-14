package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// create a cookie that violates the secure policy
func cookieHandler(w http.ResponseWriter, r *http.Request) {
	//	cookie := http.Cookie{Name: "session", Value: "12345", Secure: true}
	cookie := http.Cookie{Name: "session", Value: "12345"}
	http.SetCookie(w, &cookie)
	io.WriteString(w, "Cookie set")
}

func Hello() string {
	return "Hello, world"
}

func main() {
	fmt.Println("Hello, World")

	// time.Sleep(1000 * time.Millisecond)
	time.Sleep(1 * time.Second)

	fmt.Println("Ready for business")
	// time.Sleep(1000 * time.Millisecond)
	//	time.Sleep(1 * time.Second)

	fmt.Println("Goodbye, World")
}

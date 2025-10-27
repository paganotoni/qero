package main

import (
	"fmt"
	"net/http"

	"qero/internal"

	// Load environment variables
	_ "go.leapkit.dev/core/tools/envload"
)

func main() {
	s := internal.New()
	fmt.Println("Server started at", s.Addr())
	err := http.ListenAndServe(s.Addr(), s.Handler())
	if err != nil {
		fmt.Println("[error] starting app:", err)
	}
}

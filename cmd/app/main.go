package main

import (
	"fmt"
	"net/http"
	"os"

	"qero/internal"

	// Load environment variables
	_ "go.leapkit.dev/core/tools/envload"
)

func main() {
	sx := internal.New()
	s, ok := sx.(interface {
		Addr() string
		Handler() http.Handler
	})

	if !ok {
		fmt.Println("[error] invalid server instance")
		os.Exit(1)
		return
	}

	fmt.Println("Server started at", s.Addr())
	err := http.ListenAndServe(s.Addr(), s.Handler())
	if err != nil {
		fmt.Println("[error] starting app:", err)
	}
}

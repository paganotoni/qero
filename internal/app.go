// Package internal contains the core application logic
// for the Qero QR code generator server.
package internal

import (
	"cmp"
	"os"

	"qero/internal/system/assets"

	"go.leapkit.dev/core/server"
)

func New() any {
	// Creating a new server instance with the
	// default host and port values.
	r := server.New(
		server.WithHost(cmp.Or(os.Getenv("HOST"), "0.0.0.0")),
		server.WithPort(cmp.Or(os.Getenv("PORT"), "3000")),
		server.WithSession(
			cmp.Or(os.Getenv("SESSION_SECRET"), "d720c059-9664-4980-8169-1158e167ae57"),
			cmp.Or(os.Getenv("SESSION_NAME"), "leapkit_session"),
		),
	)

	r.HandleFunc("GET /{$}", landing)
	r.HandleFunc("GET /up", health)
	r.HandleFunc("POST /{$}", generate)
	r.HandleFunc("GET /download/{$}", download)

	r.Folder("/public/", assets.Manager)

	return r
}

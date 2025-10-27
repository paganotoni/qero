package internal

import (
	"cmp"
	"net/http"
	"os"

	"qero/internal/system/assets"

	"go.leapkit.dev/core/server"
)

// Server interface exposes the methods
// needed to start the server in the cmd/app package
type Server interface {
	Addr() string
	Handler() http.Handler
}

func New() Server {
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

	r.HandleFunc("GET /{$}", qrForm)
	r.HandleFunc("POST /{$}", qrImage)
	r.HandleFunc("GET /download/{$}", downloadQR)

	r.Folder("/public/", assets.Manager)

	return r
}

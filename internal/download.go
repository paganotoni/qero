package internal

import (
	"fmt"
	"net/http"

	qrcode "github.com/skip2/go-qrcode"
	"go.leapkit.dev/core/server"
)

func download(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	text := r.FormValue("Content")
	if text == "" {
		server.Error(w, fmt.Errorf("invalid input for the generator"), http.StatusUnprocessableEntity)
		return
	}

	var png []byte
	png, err := qrcode.Encode(text, qrcode.Medium, 256)
	if err != nil {
		server.Error(w, fmt.Errorf("failed to generate qr code: %w", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename=\"qrcode.png\"")
	w.Header().Set("Content-Type", "image/png")
	w.Write(png)
}

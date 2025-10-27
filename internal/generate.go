package internal

import (
	"encoding/base64"
	"fmt"
	"net/http"

	qrcode "github.com/skip2/go-qrcode"
	"go.leapkit.dev/core/server"
	. "maragu.dev/gomponents/html"
)

func qrImage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	text := r.FormValue("Content")
	if text == "" {
		w.Write([]byte(""))
		return
	}

	var png []byte
	png, err := qrcode.Encode(text, qrcode.Medium, 256)
	if err != nil {
		server.Error(w, fmt.Errorf("Failed to generate QR code: %w", err), http.StatusInternalServerError)
		return
	}

	// Encode the PNG bytes to a Base64 string
	encoded := base64.StdEncoding.EncodeToString(png)
	image := Img(
		Src("data:image/png;base64,"+encoded),
		Alt("QR Code"),
		Class("mx-auto w-full h-auto rounded-lg"),
	)

	image.Render(w)
}

package internal

import (
	"net/http"

	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html"
)

// landing renders the QR generation landing.
func landing(w http.ResponseWriter, r *http.Request) {
	p := page(
		Div(
			Class("text-center max-w-md mx-auto flex flex-col gap-1"),

			H1(
				Class("text-3xl font-bold"),
				Span(
					Span(Text("Welcome to ")),
					Span(Text("Qero!"), Class("text-blue-500")),
				),
			),
			P(Text("Generate shareable QR codes"), Class("text-gray-600 -mt-1 mb-4")),
			Div(Class("w-[300px] h-[300px] bg-gray-200 rounded-lg mb-2"), ID("QR")),
			Form(
				Class("flex flex-col"),
				Action("/download/"),
				Method("GET"),

				Target("_blank"),
				hx.Swap("none"),

				Textarea(
					Class("rounded-lg border-gray-300 border"),
					Placeholder("Enter text or URL here"),
					Name("Content"),
					ID("Content"),

					hx.Post("/"),
					hx.Target("#QR"),
					hx.Swap("innerHTML"),
					hx.Trigger("keyup changed delay:300ms"),
					hx.Include("#Content"),
				),
				Div(
					Class("flex flex-row justify-between mt-2"),
					Div(
						Class("cursor-pointer p-2 px-3 text-gray-600 rounded-lg border border-gray-300 bg-white"),
						Text("Clear"),
						Attr("_", "on click set #Content's value to '' then trigger keyup on #Content"),
					),

					Input(
						Type("Submit"),
						Class("cursor-pointer p-2 px-3 text-gray-600 rounded-lg border border-gray-300 bg-blue-400 text-white"),
						Value("Download"),
					),
				),
			),

			Hr(Class("mt-6 border border-gray-300")),
			P(
				Class("text-sm text-gray-500 mt-4"),
				Text("Made by "),
				A(
					Href("https://antoniopagano.com"),
					Class("text-blue-500 underline"),
					Text("Antonio Pagano"),
				),
			),
		),
	)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	p.Render(w)
}

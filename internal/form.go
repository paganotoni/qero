package internal

import (
	"net/http"

	"qero/internal/system/assets"

	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html"
)

// qrForm renders the QR generation qrForm.
func qrForm(w http.ResponseWriter, r *http.Request) {
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
		),
	)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	p.Render(w)
}

func page(content Node) Node {
	css, err := assets.Manager.PathFor("application.css")
	if err != nil {
		css = ""
	}

	return HTML(
		Lang("en"),
		Class("h-full"),
		Head(
			Meta(
				Charset("UTF-8"),
			),
			Meta(
				Name("viewport"),
				Content("width=device-width, initial-scale=1.0"),
			),
			TitleEl(
				Text("LeapKit"),
			),
			Link(
				Rel("stylesheet"),
				Href(css),
			),
			Script(
				Src("https://unpkg.com/htmx.org@2.0.0"),
			),

			Script(
				Src("https://unpkg.com/hyperscript.org@0.9.14"),
			),
		),
		Body(
			Class("h-full flex flex-row min-h-full relative"),
			Section(
				Class("flex flex-col gap-4 min-h-full w-full"),
				Main(
					Class("overflow-scroll min-h-full p-8 py-10 flex flex-col gap-5 bg-gray-100"),
					content,
				),
			),
		),
	)
}

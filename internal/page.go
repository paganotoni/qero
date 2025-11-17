package internal

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type pathFinder interface {
	Path(string) string
}

// page renders a full HTML page with the given content.
func page(content Node, assets pathFinder) Node {
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
				Href(assets.Path("application.css")),
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

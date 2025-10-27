module qero

go 1.24

require (
	github.com/mattn/go-sqlite3 v1.14.24
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
	go.leapkit.dev/core v0.1.1
	maragu.dev/gomponents v1.2.0
	maragu.dev/gomponents-htmx v0.6.1
)

require (
	github.com/fsnotify/fsnotify v1.8.0 // indirect
	github.com/gorilla/securecookie v1.1.2 // indirect
	github.com/gorilla/sessions v1.3.0 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
	go.antoniopagano.com/tailo v0.0.11 // indirect
	go.leapkit.dev/tools v0.1.2 // indirect
	golang.org/x/sys v0.21.0 // indirect
)

tool (
	go.antoniopagano.com/tailo
	go.leapkit.dev/tools/db
	go.leapkit.dev/tools/dev
)

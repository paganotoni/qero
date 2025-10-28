package internal

import "net/http"

// A simple healthcheck required by Kamal, since this app does not have any
// dependency that needs to be checked it just returns OK.
func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

package handler

import (
	"net/http"
)

func PlayHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	game, present := query["game"]

	if !present {
		w.WriteHeader(400)
		w.Write([]byte("missing `game` query parameter"))
	} else {
		w.WriteHeader(200)
		w.Write([]byte(game[0]))
	}
}

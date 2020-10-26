package controllers

import (
	"net/http"

	"github.com/kabaliserv/tmpfiles/controllers/error"
)

func renderError(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	render := error.JSONMap(code).ToJSON()
	w.Write(render)
}

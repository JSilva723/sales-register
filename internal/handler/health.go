package handler

import (
	"net/http"
	"sales-register/internal/util"
)

func Health(w http.ResponseWriter, r *http.Request) {
	util.WriteJSON(w, map[string]string{
		"status": "ok",
	})
}

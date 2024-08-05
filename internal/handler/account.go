package handler

import (
	"net/http"
	"sales-register/internal/util"
)

func Account(w http.ResponseWriter, r *http.Request) {
	type Req struct {
		ID   int32  `json:"id"`
		Name string `json:"name"`
	}
	req, err := util.ReadJSON[Req](r.Body)
	if err != nil {
		util.WriteError(w, err, http.StatusBadRequest)
		return
	}

	util.WriteJSON(w, req)
}

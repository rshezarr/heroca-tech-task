package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"salt_svc/internal"
)

func (h *Handler) generateSalt(w http.ResponseWriter, r *http.Request) {
	saltResp := internal.Salt{
		GeneratedSalt: h.service.Salt.GenerateSalt(),
	}

	if err := json.NewEncoder(w).Encode(saltResp); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

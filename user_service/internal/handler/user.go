package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"user_svc/internal"
	"user_svc/internal/usecase"
)

type SaltResponse struct {
	GeneratedSalt string `json:"generated_salt"`
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	var user internal.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest(http.MethodPost, "http://localhost:9090/generate-salt", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var saltResp SaltResponse
	if err := json.NewDecoder(resp.Body).Decode(&saltResp); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.service.User.Create(r.Context(), user, saltResp.GeneratedSalt); err != nil {
		log.Println(err)
		if errors.Is(err, usecase.ErrUserExists) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

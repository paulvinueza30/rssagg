package main

import (
	"fmt"
	"net/http"

	"github.com/paulvinueza30/rssagg/internal/database"
	"github.com/paulvinueza30/rssagg/internal/database/auth"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apicfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)

		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}
		user, err := apicfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 404, fmt.Sprintf("couldn't get user: %v", err))
			return
		}
		handler(w, r, user)
	}
}

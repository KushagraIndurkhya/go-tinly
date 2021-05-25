package middleware

import (
	"context"
	"net/http"

	"github.com/KushagraIndurkhya/go-tinly/utills"
)

type contextKey int

const AuthenticatedUserID contextKey = 0

func Auth(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("Token")
		if err == http.ErrNoCookie {
			ctx := context.WithValue(r.Context(), AuthenticatedUserID, "")
			rew := r.WithContext(ctx)

			h.ServeHTTP(w, rew)
		} else {
			isValid, id, err := utills.GetId(token.Value)
			if err != nil || !isValid {
				ctx := context.WithValue(r.Context(), AuthenticatedUserID, "")
				rew := r.WithContext(ctx)

				h.ServeHTTP(w, rew)
			} else {

				ctx := context.WithValue(r.Context(), AuthenticatedUserID, id)
				rew := r.WithContext(ctx)

				h.ServeHTTP(w, rew)
			}
		}
	})
}

package middleware

import (
	"context"
	"net/http"
	"strings"
	"url/configs"
	"url/pkg/jwt"
)

type key string

const (
	EmailContextKey key = "EmailContextKey"
)

func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authedLine := r.Header.Get("Authorization")
		if !strings.HasPrefix(authedLine, "Bearer ") {
			writeUnauthed(w)
			return
		}
		token := strings.TrimPrefix(authedLine, "Bearer ")
		IsValid, data := jwt.NewJWT(config.Auth.Secret).Parse(token)
		if !IsValid {
			writeUnauthed(w)
			return
		}
		r.Context()
		ctx := context.WithValue(r.Context(), EmailContextKey, data.Email)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}

func writeUnauthed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

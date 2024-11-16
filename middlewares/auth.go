package middlewares

import (
	"book-golang/helpers"
	"context"
	"net/http"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("Authorization")

		if accessToken == "" {
			helpers.Response(w, 401, "Unauthorized", nil)
			return
		}

		user, err := helpers.ValidateToken(accessToken)

		if err != nil {
			helpers.Response(w, 401, "Unauthorized", nil)
			return
		} 

		ctx := context.WithValue(r.Context(), "userInfo", user)


		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
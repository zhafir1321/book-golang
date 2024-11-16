package middlewares

import (
	"book-golang/helpers"
	"context"
	"net/http"
)


func Authentication(next http.Handler) http.Handler {
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


func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userInfo := r.Context().Value("userInfo")

		if userInfo == nil {
			helpers.Response(w, 401, "Unauthorized", nil)
			return
		}

		user, ok := userInfo.(*helpers.MyCustomClaims)

		if !ok || user.Role != "admin" {
			helpers.Response(w, 401, "Forbidden", nil)
			return
		}

		

		next.ServeHTTP(w, r)
	})
}

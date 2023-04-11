package middleware

import (
	"context"
	"fmt"
	"net/http"

	"awesomeProject/lib/jwt"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			http.Error(w, "missing Authorization header", http.StatusUnauthorized)
			return
		}
		claims, err := jwt.ParseToken(authorization)
		if err != nil {
			http.Error(w, "expire the token", http.StatusUnauthorized)
		}
		fmt.Println("token正确")
		_, _ = w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))

		ctx := context.WithValue(r.Context(), "claims", claims)

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}

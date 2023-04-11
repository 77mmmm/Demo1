package middleware

import (
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

// 允许10个用户进入
var limiter = rate.NewLimiter(rate.Every(time.Millisecond*31), 10)

func Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if limiter.Allow() == false {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

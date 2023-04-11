package middleware

import (
	"net/http"

	"golang.org/x/time/rate"
)

type Limiter struct {
	Lm *rate.Limiter
}

// !!!!!limiter   多长时间允许多少个进入。还是只能允许10个用户进入   允许10个用户进入

func (limiter *Limiter) Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if limiter.Lm.Allow() == false {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

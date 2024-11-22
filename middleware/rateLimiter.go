package middlewares

import (
	"math"
	"net/http"
	"strconv"
	"time"
)

type RateLimiter struct {
	Period  time.Duration
	MaxRate int64
	DB   any
}

func (r RateLimiter) WriteHeaders (
	w http.ResponseWriter,
	used int64,
	expireTime time.Duration,
){
	limit := r.MaxRate
	remaining := int64(math.Max(float64(limit-used), 0))
	reset := int64(math.Ceil(expireTime.Seconds()))

	w.Header().Add("X-RateLimit-Limit", strconv.FormatInt(limit, 10))
	w.Header().Add("X-RateLimit-Remaining", strconv.FormatInt(remaining, 10))
	w.Header().Add("X-RateLimit-Reset", strconv.FormatInt(reset, 10))

}
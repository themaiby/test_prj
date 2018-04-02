package middleware

import "net/http"
import (
	"github.com/gorilla/context"
	log "github.com/kataras/golog"
	"time"
)

// Query handler for LOG
type LogWrapper struct {
	http.Handler
}

func (wr LogWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debug(r.Method, " | local.domain.com", r.URL, " | ", r.UserAgent(), " | ", r.RemoteAddr, " | ")
	wr.Handler.ServeHTTP(w, r)
}

func SetRequestTime(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		RequestTime := time.Now()
		// Middleware operations
		// Parse body/get token.
		context.Set(r, "RequestTime", RequestTime)

		next.ServeHTTP(w, r)
	})
}

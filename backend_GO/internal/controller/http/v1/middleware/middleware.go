package middleware

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func LoggingMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		start := time.Now()

		lrw := &loggingResponseWriter{ResponseWriter: w}
		next(lrw, r, ps)

		slog.Info("request",
			"method", r.Method,
			"path", r.URL.Path,
			"remote_addr", r.RemoteAddr,
			"user_agent", r.UserAgent(),
			"status", lrw.statusCode,
			"duration", time.Since(start).String(),
		)
	}
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

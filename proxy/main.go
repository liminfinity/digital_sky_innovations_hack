package main

import (
	"log/slog"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"
)

const (
	host        = "localhost"
	pidsPort    = "8000"
	historyPort = "8080"

	proxyPort = "80"
)

func main() {
	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}),
	)

	pidsProxy := MustCreateProxy(log, "http://"+net.JoinHostPort(host, pidsPort))
	historyProxy := MustCreateProxy(log, "http://"+net.JoinHostPort(host, historyPort))

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/pids-service/", func(w http.ResponseWriter, r *http.Request) {
		pidsProxy.ServeHTTP(w, r)
	})

	mux.HandleFunc("/api/v1/history-service/", func(w http.ResponseWriter, r *http.Request) {
		historyProxy.ServeHTTP(w, r)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
	})

	server := &http.Server{
		Addr:    net.JoinHostPort("0.0.0.0", proxyPort),
		Handler: loggingMiddleware(mux),
	}

	log.Info("Gateway server started on", slog.String("address", server.Addr))

	err := server.ListenAndServe()
	if err != nil {
		log.Error("failed to start server", slog.String("err", err.Error()))
		os.Exit(1)
	}
}

func MustCreateProxy(log *slog.Logger, target string) *httputil.ReverseProxy {
	u, err := url.Parse(target)
	if err != nil {
		log.Error("failed to parse target url", slog.String("err", err.Error()))
		panic(err)
	}
	return httputil.NewSingleHostReverseProxy(u)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		lrw := &loggingResponseWriter{ResponseWriter: w}
		next.ServeHTTP(lrw, r)

		slog.Info("request",
			"method", r.Method,
			"path", r.URL.Path,
			"remote_addr", r.RemoteAddr,
			"user_agent", r.UserAgent(),
			"status", lrw.statusCode,
			"duration", time.Since(start).String(),
		)
	})
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

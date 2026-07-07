package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/your-org/kea-ui/internal/config"
	"github.com/your-org/kea-ui/internal/handler"
	"github.com/your-org/kea-ui/internal/web"
)

func main() {
	cfg, err := config.Load("/config/config.json")
	if err != nil {
		logger.Error("cannot load config", "error", err)
		os.Exit(1)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	mux := http.NewServeMux()

	mux.Handle("/static/", web.Static())

	mux.HandleFunc("/", handler.Index())

	logger.Info("server started", "listen", cfg.Server.Listen)

	err = http.ListenAndServe(cfg.Server.Listen, logging(logger, mux))
	if err != nil {
		logger.Error(err.Error())
	}
}

func logging(log *slog.Logger, next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Info(
			"request",
			"method", r.Method,
			"path", r.URL.Path,
			"ip", r.RemoteAddr,
		)

		next.ServeHTTP(w, r)

	})

}

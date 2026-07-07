package main

import (
    "log/slog"
    "net/http"
    "os"
    "github.com/jafsq5/kea-tool-ui/internal/config"
    "github.com/jafsq5/kea-tool-ui/internal/handler"
    "github.com/jafsq5/kea-tool-ui/internal/web"
)

func main() {

    logger := slog.New(
        slog.NewJSONHandler(os.Stdout, nil),
    )

    //Add support env CONFIG_FILE
    configPath := os.Getenv("CONFIG_FILE")
    if configPath == "" {
        configPath = "configs/config.json"
    }

    cfg, err := config.Load(configPath)
    if err != nil {
        logger.Error("cannot load config", "error", err)
        os.Exit(1)
    }

    mux := http.NewServeMux()

    mux.Handle("/static/", web.Static())

    mux.HandleFunc("/", handler.Index())

    logger.Info("starting server",
        "listen", cfg.Server.Listen,
    )

    err = http.ListenAndServe(cfg.Server.Listen, mux)
    if err != nil {
        logger.Error(err.Error())
        os.Exit(1)
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

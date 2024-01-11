package main

import (
	"github.com/terratensor/gmx-server/server/internal/config"
	"log/slog"
	_ "log/slog"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettyLog()
	}
}

func setupPrettyLog() *slog.Logger {

}

func setupJSONLog() *slog.Logger {
	opt := slogpretty.PrettyHandlerOption{}
}

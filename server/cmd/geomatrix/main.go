package main

import (
	"fmt"
	"github.com/terratensor/gmx-server/server/internal/config"
	"github.com/terratensor/gmx-server/server/internal/db/pgstore"
	"github.com/terratensor/gmx-server/server/internal/lib/logger/handlers/slogpretty"
	"github.com/terratensor/gmx-server/server/internal/lib/logger/sl"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)
	// к каждому сообщению будет добавляться поле с информацией о текущем окружении
	log = log.With(slog.String("env", cfg.Env))

	log.Info("initializing server", slog.String("address", cfg.Address))
	log.Debug("logger debug mode enabled")
	//ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	dsn := makeDsn(cfg)
	_, err := pgstore.NewPgStore(dsn)
	if err != nil {
		log.Error("failed to initialize storage", sl.Err(err))
		os.Exit(1)
	}

}

func makeDsn(cfg *config.Config) string {
	dsn := fmt.Sprintf(
		"postgresql://%v:%v@%v:%v/%v?sslmode=disable",
		cfg.Storage.User,
		cfg.Storage.Password,
		cfg.Storage.Host,
		cfg.Storage.Port,
		cfg.Storage.Db,
	)
	return dsn
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}

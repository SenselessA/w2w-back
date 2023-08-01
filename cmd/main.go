package main

import (
	"github.com/SenselessA/w2w_backend/internal/config"
	"github.com/SenselessA/w2w_backend/internal/db/postgres"
	"github.com/SenselessA/w2w_backend/internal/middleware"
	"github.com/SenselessA/w2w_backend/internal/repository"
	"github.com/SenselessA/w2w_backend/internal/services"
	"github.com/SenselessA/w2w_backend/internal/transport"
	"github.com/SenselessA/w2w_backend/pkg/hash"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"
)

var (
	configPath string
)

const (
	CONFIG_DIR  = "configs"
	CONFIG_FILE = "main"
)

func init() {
	// flag.StringVar(&configPath, "config-path", "configs/want2watch.toml", "path to config file")
	logrus.SetFormatter(new(logrus.JSONFormatter))
}

func main() {
	err := os.Unsetenv("PGLOCALEDIR")
	if err != nil {
		logrus.Fatalln(err)
		return
	}

	cfg, err := config.Init(CONFIG_DIR, CONFIG_FILE)
	if err != nil {
		logrus.Fatalln(err)
		return
	}

	middlewares := middleware.InitMiddlewares(cfg.Secret)

	db, err := postgres.Open(&cfg.DB)
	if err != nil {
		logrus.Fatalln(err)
		return
	}
	// defer db.Close(context.Background())
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	hasher := hash.NewSHA1Hasher(cfg.Secret.Jwt)

	repo := repository.New(db)
	service := services.New(repo, hasher)
	handler := transport.NewHandler(service, middlewares)

	err = service.Kodik.StartParseMovies(cfg.Secret.Kodik)
	if err != nil {
		log.Fatal("FATAL KODIK: ", err)
		return
	}

	ticker := time.NewTicker(5 * time.Hour)
	go func(ticker *time.Ticker) {
		for {
			select {
			case <-ticker.C:
				// once 5 hours
				err = service.Kodik.StartParseMovies(cfg.Secret.Kodik)
				if err != nil {
					log.Fatal("FATAL KODIK: ", err)
					return
				}
			}
		}
	}(ticker)

	srv := handler.Init()

	err = srv.Listen(cfg.HTTP.Port)
	if err != nil {
		log.Fatal(err)
		return
	}
}

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"shortener/internal/bootstrap"
	"shortener/internal/handler"
	"shortener/internal/repository"
	"shortener/internal/server"
	"shortener/internal/service"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

const (
	port = "port"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	mode := os.Getenv("MODE")
	var repos repository.Repository
	var db *pgx.Conn
	if mode == "db" {
		err := godotenv.Load("configs/.env")
		if err != nil {
			log.Fatalf("error loading env variables: %s", err.Error())
		}

		db, err = bootstrap.NewPostgresDB()
		if err != nil {
			log.Fatalf("failed to initialize db: %s", err.Error())
		}

		repos = repository.NewRepositoryDB(db)
	} else if mode == "inMemory" {
		repos = repository.NewRepositoryIM()
	} else {
		log.Fatal("mode should be inMemory or db")
	}

	services := service.New(repos)
	handlers := handler.New(services)
	srv := new(server.Server)

	go func() {
		if err := srv.Run(viper.GetString(port), handlers.InitRoutes()); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
	log.Print("started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("error occured on server shutting down: %s", err.Error())
	}
	if mode == "db" {
		if err := db.Close(context.Background()); err != nil {
			log.Printf("error occured on db connection closing: %s", err.Error())
		}
	}
}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

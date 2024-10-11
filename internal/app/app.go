package app

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"api-demo/internal/config"
	"api-demo/internal/controller"
	"api-demo/internal/model"
	"api-demo/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// TODO: add json body parsing/validation --- DONE

// TODO: create database connection and pass it to repository --- DONE
// TODO: implement database with GORM --- ALMOST DONE

// TODO: move database connection to separate function/file
// TODO: create migration script

func Run() {
	// Load config
	config, err := config.NewConfig()
	if err != nil {
		log.Fatal().Msgf("failed to load config: %v", err)
	}

	// Create database connection
	conn, err := sql.Open("pgx", config.DatabaseURL)
	if err != nil {
		log.Fatal().Msgf("failed to connect to database: %v", err)
		return
	}

	// TODO: remove later
	if err := conn.Ping(); err != nil {
		log.Fatal().Msgf("failed to ping database: %v", err)
		return
	} else {
		log.Info().Msg("database connection successful")
	}

	// Create GORM database
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: conn,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal().Msgf("failed to create database connection: %v", err)
		return
	}

	// Migrate database
	// TODO: move to migration script?
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatal().Msgf("failed to migrate database: %v", err)
		return
	}

	// Create user repository
	userRepository := repository.NewUserRepository(db)

	// Create router
	router := mux.NewRouter()

	// Create schema validator
	validate := validator.New(validator.WithRequiredStructEnabled())

	// Create user controller
	UserController := controller.NewUserController(userRepository, validate)

	// Register user handlers
	router.HandleFunc("/user", UserController.CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", UserController.GetUser).Methods("GET")

	// Create http server
	srv := &http.Server{
		Addr:    config.ServerURL,
		Handler: router,
	}

	srvErrChan := make(chan error)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			srvErrChan <- err
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	select {
	case err := <-srvErrChan:
		log.Err(err).Msg("unexpected server error")
	case <-c:
		log.Info().Msg("server received signal")
	}

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer shutdownCancel()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		log.Info().Msg("shutting down server")
		if err := srv.Shutdown(shutdownCtx); err != nil {
			log.Err(err).Msg("server shutdown failed")
		}
	}()

	wg.Wait()
	log.Info().Msg("server shutdown successfully")
}

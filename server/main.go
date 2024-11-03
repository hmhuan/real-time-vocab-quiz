package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/real-time-vocab-quiz/server/configs"
	"github.com/real-time-vocab-quiz/server/dataservices"
	"github.com/real-time-vocab-quiz/server/routes"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()

	err := configs.LoadConfigs()
	if err != nil {
		log.Fatalln("Failed to load configurations", err)
	}

	// 1 Create redis connection
	opt, err := redis.ParseURL("redis://localhost:6379/0") // TODO: hardcode for dev only
	if err != nil {
		log.Fatalln("Failed to create redis connection", err)
	}

	rdb := redis.NewClient(opt)

	// 2 Create connection to database
	pgInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=%s",
		configs.GetDBHost(), configs.GetDBPort(),
		configs.GetDBUser(), configs.GetDBPassword(),
		configs.GetDBName(), configs.GetDBSchema(),
	)
	db, err := sqlx.Connect("postgres", pgInfo)
	if err != nil {
		log.Fatalln("Failed to create db connection", err)
	}
	defer db.Close()
	// Setup conenction pooling for db
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)

	dataServiceHolder := dataservices.NewHolder(db)

	if configs.GetENV() == configs.ReleaseEnv {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	routes.Setup(router, dataServiceHolder, rdb)

	address := fmt.Sprintf("%s:%s", configs.GetHost(), configs.GetPort())

	server := &http.Server{Addr: address, Handler: router}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to listening and serve: %s\n", err)
		}
	}()
	log.Printf("Starting server on %s\n", address)
	<-ctx.Done()

	stop()
	log.Println("Shutting down gracefully...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	os.Exit(0)
}

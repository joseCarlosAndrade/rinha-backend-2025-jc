package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/joseCarlosAndrade/rinha-backend-2025-jc/internal/adapter/api"
	"github.com/joseCarlosAndrade/rinha-backend-2025-jc/internal/adapter/http/server"
	"github.com/joseCarlosAndrade/rinha-backend-2025-jc/internal/adapter/redis"
	"github.com/joseCarlosAndrade/rinha-backend-2025-jc/internal/core/config"
	"github.com/joseCarlosAndrade/rinha-backend-2025-jc/internal/core/service"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

const (
	shutdownTimeout time.Duration = 10*time.Second

)

// this inits a new zap development logger
func initLogger () {
	var err error

	logger, err = zap.NewDevelopment()

	if err != nil {
		log.Fatal("failed to start logger: ", err.Error())
	}
}

// loading env var and filling them into config.App 
func loadEnv() error {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
		return err
	}

	if err := envconfig.Process("app", &config.App); err != nil {
		log.Fatal(err.Error())
	}

	return nil
}

// syncs logger to properly flush all logs
func cleanupLogger(undoFn func()) {
	_ = logger.Sync()

	undoFn()
}	

// TODO: setup redis storage system
func setupStorage(ctx context.Context) (redis.Repository, error) {
	return redis.NewRedisRepository(), nil
}

func main() {
	err := loadEnv()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// initig logger, replacing global logger and defering its cleanup
	initLogger()
	undo := zap.ReplaceGlobals(logger)
	defer cleanupLogger(undo)

	ctx := context.Background()

	// storage handler
	redis, err := setupStorage(ctx)
	if err != nil {
		zap.L().Fatal("could not set redis storage", zap.Error(err))
	}

	// outgoing api calls handler
	api := api.NewAPIRepository()

	// service repository
	svc := service.NewService(redis, api)

	// server handler (using gin engine)
	controller := server.NewController( svc)

	// starting everything
	serv := startServer(controller.Gateway)

	// wainting for signals to shutdown!
	waitForShutdown(ctx, serv)
}

// initializes a server instance and starts it on another go routine. returns a http.Server so its possible to manage it (kill it)
func startServer(router*gin.Engine) *http.Server {
	srv := &http.Server{
		Addr: ":" + config.App.Port,
		Handler: router,
	}

	go func () {
		zap.L().Info("starting code at ", zap.String("port", config.App.Port))

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("startup failed on server", zap.Error(err))
		}
	}()

	return srv
}

// waits for a ctrl c or shutdown signal. this way, the server properly finishes its resources before exiting (if possible)
// uses a timeout context to prevent endless shutdown
func waitForShutdown(ctx context.Context, serv * http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<- quit // waits for a shutdown 

	shutdownCtx, cancel := context.WithTimeout(ctx, shutdownTimeout)
	defer cancel()

	if err := serv.Shutdown(shutdownCtx); err != nil {
		zap.L().Fatal("forcing shutting down")
	}

	zap.L().Info("server exited")
}
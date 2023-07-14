package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	userHandler "github.com/Hideinbruh/test-health/internal/handler/user"
	userRepo "github.com/Hideinbruh/test-health/internal/repository/user"
	userServer "github.com/Hideinbruh/test-health/internal/server/user"
	userService "github.com/Hideinbruh/test-health/internal/service/user"
	"github.com/Hideinbruh/test-health/pkg/logger"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	ctx := context.Background()

	// Чтобы не тратить время на парсинг переменных окружения, я сразу строку передам. Про переменные
	// окружения расскажу на уроке
	pgCfg, err := pgxpool.ParseConfig(
		"host=localhost port=5432 dbname=health user=health-user password=health-password sslmode=disable",
	)
	if err != nil {
		logger.Fatalf("failed to get db config: %s", err.Error())
	}

	//Нужон контекст и конфиг подключения
	dbc, err := pgxpool.ConnectConfig(ctx, pgCfg)
	if err != nil {
		logger.Fatalf("failed to get db connection: %s", err.Error())
	}

	srv := new(userServer.Server)

	// Теперь когда коннект сделан и создан слой репозитория, его надо передать в сервис, чтобы там воспользоваться
	repository := userRepo.NewRepository(dbc)
	service := userService.NewService(repository)
	handler := userHandler.NewHandler(service)

	go func() {
		if err := srv.Run("8080", handler.InitRoutes()); err != nil {
			log.Fatalf("run server: %s\n", err.Error())
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	fmt.Println("app stopped")
}

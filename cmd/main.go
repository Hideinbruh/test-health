package main

import (
	userHandler "awesomeProject3/test-health/internal/handler/user"
	userServer "awesomeProject3/test-health/internal/server/user"
	userService "awesomeProject3/test-health/internal/service/user"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	srv := new(userServer.Server)
	service := userService.NewService()
	handler := userHandler.NewHanlder(ctx, service)

	fmt.Println("my-health app started")
	go func() {
		if err := srv.Run("8080", handler.InitRoutes()); err != nil {
			log.Fatalf("run server: %s\n", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	fmt.Println("my-health app stopped")

}

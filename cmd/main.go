package main

import (
	userHandler "awesomeProject3/internal/handler/user"
	userServer "awesomeProject3/internal/server/user"
	userService "awesomeProject3/internal/service/user"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	srv := new(userServer.Server)
	service := userService.NewService()
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

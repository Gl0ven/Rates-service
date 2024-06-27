package main

import (
	"Gl0ven/kata_projects/rates/internal/app"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	err := godotenv.Load("/rates/.env")
	if err != nil {
		panic(err)
	}

	app := app.Bootstrap()

	signChan := make(chan os.Signal, 1)
	signal.Notify(signChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		app.Logger.Info("starting grpc server on port: " + app.Port)
		if err := app.Server.Serve(app.Listener); err != nil {
			app.Logger.Fatal("server initialization error", zap.Error(err))
		}
	}()

	<-signChan

	app.Server.GracefulStop()

	app.Logger.Info("Server stopped gracefully")
}

package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"runExec/internal/config"
	"runExec/internal/runner"
)

func Execute() (cfg *config.Config, err error) {

	fmt.Print("starting...")
	cfg, err = config.InitConfig()
	if err != nil {
		log.Fatalf("can't load config: %v", err)
		return
	}

	logger, err := config.InitLogger(cfg.Debug, cfg.AppName)
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
		return
	}
	cfg.Logger = logger
	logger.Info("initializing the service...")

	// handle service stop
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		sig := <-quit
		logger.Info(fmt.Sprintf("caught sig: %+v", sig))
	}()
	return
}

func Work(cfg *config.Config) {

	if err := runner.CheckPath(cfg.PathCheck); err != nil {
		cfg.Logger.Fatal("Нет такой папки")
		os.Exit(1)
	}
	if err := runner.RunApp(cfg.PathApp); err != nil {
		cfg.Logger.Fatal("Не смогли запустить приложение")
		os.Exit(1)
	}

}

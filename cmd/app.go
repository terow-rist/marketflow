package cmd

import (
	"fmt"
	"log"
	"log/slog"
	"marketflow/internal/adapter/config"
	flag "marketflow/internal/adapter/config"
	httpserver "marketflow/internal/adapter/handler/http"
	"marketflow/internal/adapter/logger"
	"net/http"
)

func Run() {

	// Parse flags
	err := flag.Parse()
	if err != nil {
		log.Fatal(err)
	}

	// Load env variables
	config := config.New()

	//Set logger
	logger.Set()
	slog.Info("Staring application", "app", config.App.Name, "env", config.App.Env)

	//Futere init db

	//Server setup
	temp := httpserver.NewTemp()

	mux := httpserver.NewRouter(temp)
	slog.Info(fmt.Sprintf("Listening on port: %d", flag.Port))
	err = http.ListenAndServe(fmt.Sprintf(":%d", flag.Port), mux)
	log.Fatal(err)
}

package cmd

import (
	"log"
	"log/slog"
	cfg "marketflow/internal/adapter/config"
	flag "marketflow/internal/adapter/config"
	"marketflow/internal/adapter/exchange"
	"marketflow/internal/adapter/logger"
	"marketflow/internal/core/domain"
	"marketflow/internal/core/service"
)

func Run() {
	// Parse flags
	err := flag.Parse()
	if err != nil {
		log.Fatalf("Failed to parse flags. %s", err.Error())
	}

	// Load env variables
	config := cfg.Init()

	//Set logger
	logger.Set()
	slog.Info("Staring application", "app", config.Info.Name, "env", config.Info.Env)

	//stream
	updates1 := make(chan domain.PriceUpdate)
	updates2 := make(chan domain.PriceUpdate)
	updates3 := make(chan domain.PriceUpdate)

	go exchange.NewListener("exchange1:40101", "exchange1", updates1).Start()
	go exchange.NewListener("exchange2:40102", "exchange2", updates2).Start()
	go exchange.NewListener("exchange3:40103", "exchange3", updates3).Start()

	//service init
	patternService := service.NewPatternService()

	result1 := patternService.FanOut(updates1)
	result2 := patternService.FanOut(updates2)
	result3 := patternService.FanOut(updates3)

	merged := patternService.FanIn(result1, result2, result3)
	for update := range merged {
		slog.Info("Merged update",
			"exchange", update.Exchange,
			"symbol", update.Symbol,
			"price", update.Price)
	}

	//Futere init db
	//db, err := PostgreSQL.Connect(ctx, config)

	//Server setup
	// temp := httpserver.NewTemp()

	// mux := httpserver.NewRouter(temp)
	// slog.Info(fmt.Sprintf("Listening on port: %d", flag.Port))
	// err = http.ListenAndServe(fmt.Sprintf(":%d", flag.Port), mux)
	// log.Fatal(err)
}

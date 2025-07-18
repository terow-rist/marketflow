package cmd

import (
	"log"
	"log/slog"
	"marketflow/internal/adapter/config"
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
		log.Fatal(err)
	}

	// Load env variables
	config := config.New()

	//Set logger
	logger.Set()
	slog.Info("Staring application", "app", config.App.Name, "env", config.App.Env)

	//stream
	updates1 := make(chan domain.PriceUpdate)
	updates2 := make(chan domain.PriceUpdate)
	updates3 := make(chan domain.PriceUpdate)

	listener1 := exchange.NewListener("exchange1:40101", "exchange1", updates1)
	listener2 := exchange.NewListener("exchange2:40102", "exchange2", updates2)
	listener3 := exchange.NewListener("exchange3:40103", "exchange3", updates3)

	go listener1.Start()
	go listener2.Start()
	go listener3.Start()

	pd1 := service.FanOut(updates1)
	pd2 := service.FanOut(updates2)
	pd3 := service.FanOut(updates3)

	merged := service.FanIn(pd1, pd2, pd3)
	for update := range merged {
		slog.Info("Merged update",
			"exchange", update.Exchange,
			"symbol", update.Symbol,
			"price", update.Price)
	}

	//Futere init db

	//Server setup
	// temp := httpserver.NewTemp()

	// mux := httpserver.NewRouter(temp)
	// slog.Info(fmt.Sprintf("Listening on port: %d", flag.Port))
	// err = http.ListenAndServe(fmt.Sprintf(":%d", flag.Port), mux)
	// log.Fatal(err)
}

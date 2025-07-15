package cmd

import (
	"bufio"
	"log"
	"log/slog"
	"marketflow/internal/adapter/config"
	flag "marketflow/internal/adapter/config"
	"marketflow/internal/adapter/logger"
	"net"
	"os"
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
	tempCatch()

	// //Futere init db

	// //Server setup
	// temp := httpserver.NewTemp()

	// mux := httpserver.NewRouter(temp)
	// slog.Info(fmt.Sprintf("Listening on port: %d", flag.Port))
	// err = http.ListenAndServe(fmt.Sprintf(":%d", flag.Port), mux)
	// log.Fatal(err)
}

func tempCatch() {
	conn, err := net.Dial("tcp", "exchange1:40101") // Connect to exchange1
	if err != nil {
		slog.Error("Failed to connect to exchange1", "error", err)
		os.Exit(1)
	}
	defer conn.Close()

	slog.Info("Connected to exchange1")

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		msg := scanner.Text()
		slog.Info("Price update", "data", msg)
	}

	if err := scanner.Err(); err != nil {
		slog.Error("Error reading from exchange1", "error", err)
	}
}

package exchange

import (
	"bufio"
	"encoding/json"
	"log/slog"
	"marketflow/internal/core/domain"
	"net"
	"os"
	"time"
)

type Listener struct {
	Address  string
	Exchange string
	Out      chan<- domain.PriceUpdate
}

func NewListener(address, exchange string, out chan<- domain.PriceUpdate) *Listener {
	return &Listener{
		Address:  address,
		Exchange: exchange,
		Out:      out,
	}
}

// Not ready yet.
func (l *Listener) Start() {
	conn, err := net.Dial("tcp", l.Address) // Connect to exchange1
	if err != nil {
		slog.Error("Failed to connect", "exchange", l.Exchange, "error", err)
		os.Exit(1)
	}
	defer conn.Close()

	slog.Info("Connected to exchange", "exchange", l.Exchange)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {

		rawMsg := scanner.Text()
		var update struct {
			Symbol    string  `json:"symbol"`
			Price     float64 `json:"price"`
			Timestamp int64   `json:"timestamp"`
		}

		if err := json.Unmarshal([]byte(rawMsg), &update); err != nil {
			slog.Error("Failed to parse JSON", "exchange", l.Exchange)
			continue
		}

		// Convert epoch millis to time.Time
		ts := time.UnixMilli(update.Timestamp)

		l.Out <- domain.PriceUpdate{
			Symbol:    update.Symbol,
			Price:     update.Price,
			Timestamp: ts,
			Exchange:  l.Exchange,
		}

	}

	if err := scanner.Err(); err != nil {
		slog.Error("Error reading from exchange1", "error", err)
	}
}

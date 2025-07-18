package service

import (
	"log/slog"
	"marketflow/internal/core/domain"
)

// TODO refactor fan out to separate
func FanOut(in <-chan domain.PriceUpdate, workerCount int, worker func(domain.PriceUpdate)) {
	for i := 0; i < workerCount; i++ {
		go func(id int) {
			for update := range in {
				slog.Info("Worker processing update", "worker", id, "symbol", update.Symbol, "price", update.Price)
				worker(update)
			}
		}(i + 1)
	}
}

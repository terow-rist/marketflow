package service

import (
	"log/slog"
	"marketflow/internal/core/domain"
	"sync"
)

func FanOut(input <-chan domain.PriceUpdate) <-chan domain.PriceUpdate {
	out := make(chan domain.PriceUpdate)
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			for update := range input {
				worker(update)
				out <- update
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func worker(update domain.PriceUpdate) {
	slog.Info("Processed update", "exchange", update.Exchange, "symbol", update.Symbol, "price", update.Price)
}

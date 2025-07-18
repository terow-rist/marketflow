package service

import (
	"marketflow/internal/core/domain"
	"sync"
)

func FanIn(inputs ...<-chan domain.PriceUpdate) <-chan domain.PriceUpdate {
	out := make(chan domain.PriceUpdate)
	var wg sync.WaitGroup

	wg.Add(len(inputs))
	for _, ch := range inputs {
		go func(ch <-chan domain.PriceUpdate) {
			defer wg.Done()
			for update := range ch {
				out <- update
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

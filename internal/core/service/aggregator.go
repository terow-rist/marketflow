package service

import "marketflow/internal/core/domain"

func FanIn(inputs ...<-chan domain.PriceUpdate) <-chan domain.PriceUpdate {
	out := make(chan domain.PriceUpdate)
	for _, ch := range inputs {
		go func() {
			for update := range ch {
				out <- update
			}
		}()
	}
	return out
}

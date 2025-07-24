package port

import "marketflow/internal/core/domain"

type PatternService interface {
	FanIn(inputs ...<-chan domain.PriceUpdate) <-chan domain.PriceUpdate
	FanOut(input <-chan domain.PriceUpdate) <-chan domain.PriceUpdate
}

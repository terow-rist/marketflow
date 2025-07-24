package service

// import (
// 	"fmt"
// 	"log/slog"
// 	"marketflow/internal/core/domain"
// 	"sync"
// )

// func worker(jobs <-chan domain.PriceUpdate, results chan<- string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for job := range jobs {
// 		results <- "Processed update exchange " + job.Exchange + "symbol" + job.Symbol + "price" + fmt.Sprintf("%f", job.Price)
// 	}
// }

// func collectResults(results <-chan string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for res := range results {
// 		slog.Info(res)
// 	}
// }

// func dispatcher(jobs <-chan domain.PriceUpdate, workerCount int) chan string {
// 	results := make(chan string)

// 	// Start workers
// 	var wg sync.WaitGroup
// 	wg.Add(workerCount)
// 	for w := 0; w < workerCount; w++ {
// 		go worker(jobs, results, &wg)
// 	}

// 	// Close results channel once all workers are done
// 	go func() {
// 		wg.Wait()
// 		close(results)
// 	}()

// 	// Start collecting results
// 	var resultsWg sync.WaitGroup
// 	resultsWg.Add(1)
// 	go collectResults(results, &resultsWg)

// 	// Wait for collector to finish
// 	resultsWg.Wait()

// 	return results

// }

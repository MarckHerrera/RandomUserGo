package aggregator

import (
	"context"
	"log"
	"randomNasa/internal/dto"
	"randomNasa/internal/interfaces"
	"randomNasa/internal/models"
	"sync"
	"time"
)

type RandomNasaResult struct {
	Info Info
	Results[]models.User
	Err error
	Duration time.Duration
}

type Info struct {
	Type string
	TotalUsers int
}

func fetchRandomNasa(ctx context.Context, service interfaces.RandomNasaService, results chan RandomNasaResult)  {
	start := time.Now()
	users, err := service.Search(ctx)

	duration := time.Since(start)

	conteo := 0

	if users != nil {
		conteo = len(users.Results)
	}


	results <- RandomNasaResult{
		Info: Info{
			Type: "success",
			TotalUsers: conteo,
		},
		Results: users.Results,
		Err: err,
		Duration: duration,
	}
}

func Aggregate(ctx context.Context, service interfaces.RandomNasaService) (*dto.GetResponse, error) {
	const peticiones = 20
	results := make(chan RandomNasaResult, peticiones)

	var wg sync.WaitGroup
	for i := 0; i < peticiones; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			fetchRandomNasa(ctx, service, results)
		}()

	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var allRandomNasa []models.User

	for result := range results  {
		/* logsito de errores */

		if result.Err != nil {
			log.Printf("Error: %v", result.Err)
			continue
		}
		allRandomNasa = append(allRandomNasa, result.Results...)
	}

	return &dto.GetResponse{
		Info: dto.ResponseInfo{
			Type: "success",
			TotalUsers: len(allRandomNasa),
		},
		Results: allRandomNasa,
	}, nil
}
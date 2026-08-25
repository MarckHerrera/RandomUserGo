package aggregator

import (
	"context"
	"sync"
	"log"

	"randomusers/internal/interfaces"
	"randomusers/internal/models"
	"randomusers/internal/dto"
	"randomusers/internal/pipeline"
)

const (
	totalPeticiones = 20
	numWorkers = 5
)

/* Como el aggregator y el pipeline lo usan lo pase a models */

/* type RandomUsersResult struct {
	Info Info
	Results []models.User
	Err error
	Duration time.Duration
}

type Info struct {
	Type string
	TotalUsers int
} */


/* Este fetch lo tengo que hacer en un worker, pero no lo tengo que hacer en el aggregator */

/* func fetchRandomUser(ctx context.Context, service interfaces.RandomUserService, results chan RandomUsersResult) {
	start := time.Now()
	users, err := service.Search(ctx)

	duration := time.Since(start)

	conteo := 0

	if users != nil {
		conteo = len(users.Results)
	}
	log.Printf("Peticion terminada: %d usuarios, err=%v, duracion=%s", conteo, err, duration)

	results <- RandomUsersResult{
		Info: Info{
			Type: "success",
			TotalUsers: len(users.Results),
		},
		Results: users.Results,
		Err: err,
		Duration: duration,
	}
}
 */
 
func Aggregate(ctx context.Context, service interfaces.RandomUserService) (*dto.GetResponse, error) {
	jobs := make(chan int, totalPeticiones)
	results := make(chan models.RandomUserJobResult, totalPeticiones)
	
	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			pipeline.Worker(ctx, workerID, service, jobs, results)
		}(w)
	}

	for i := 1; i <= totalPeticiones; i++ {
		jobs <- i
	}

	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	var allRandomUsers []models.User

	for result := range results {
		if result.Err != nil {
			log.Printf("Error en la peticion: %v (pero seguimos :) )", result.Err)
			continue
		}
		allRandomUsers = append(allRandomUsers, result.Results...)
	}

	fullResult := &dto.GetResponse{
		Info: dto.ResponseInfo{
			Type: "success",
			TotalUsers: len(allRandomUsers),
		},
		Results: allRandomUsers,
	}

	return fullResult, nil
}
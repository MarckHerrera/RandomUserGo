package pipeline

import (
	"context"
	"time"
	"log"
	"fmt"

	"randomusers/internal/interfaces"
	"randomusers/internal/models"
)

const maxRetry = 3

func Worker(ctx context.Context, id int, service interfaces.RandomUserService, jobs <-chan int, results chan<- models.RandomUserJobResult) {

	for jobID := range jobs {
		start := time.Now()
		users, err := Retry(ctx, service, jobID, maxRetry)

		duration := time.Since(start)

		conteo := 0

		if users != nil {
			conteo = len(users.Results)
		}

		log.Printf("Worker: %d | JobID: %d | %d usuarios | err=%v | duracion=%s", id, jobID, conteo, err, duration)

		var reultUsers []models.User
		if users != nil {
			reultUsers = users.Results
		}

		results <- models.RandomUserJobResult{
			Info: models.ResultInfo{
				Type: "success",
				TotalUsers: conteo,
			},
			Results: reultUsers,
			Err: err,
			Duration: duration,
		}
	}
}

func Retry(ctx context.Context, service interfaces.RandomUserService, jobID int, maxRetry int) (*models.RandomUser, error) {
	var lastErr error

	for intento := 1; intento <= maxRetry; intento++ {
		users, err := service.Search(ctx)

		if err == nil && users != nil && len(users.Results) > 0 {
			if intento > 1 {
				log.Printf("Job %d: recuperado en el intento %d", jobID, intento)
			}
			return users, nil
		}

		if err != nil {
			lastErr = err
		}else{
			lastErr = fmt.Errorf("respuesta vacia en intento %d", intento)
		}

		log.Printf("Job %d: intento %d | err=%v", jobID, intento, lastErr)

		if ctx.Err() != nil {
			return nil, ctx.Err()
		}

		time.Sleep(time.Duration(intento) * 200 * time.Millisecond)
	}

	return nil, lastErr
}
package workers

import (
	"context"
	"log"
	"seven-solutions-challenge/internal/adapters/outbound/db/mongo/repositories"
	"seven-solutions-challenge/internal/app/ports"
	"sync"
	"time"
)

type ListUsersWorker struct {
	name     string
	userRepo repositories.IUserRepo
}

func NewListUsersWorker(userRepo repositories.IUserRepo) ports.IWorkers {
	return &ListUsersWorker{
		name:     "LIST_USER_WORKER",
		userRepo: userRepo,
	}
}

// Start implements IBackgroudWorkers.
func (b *ListUsersWorker) Start(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			log.Println("Listing user(s) worker shutting down...")
			return
		default:
			resp, err := b.userRepo.List(ctx)
			if err != nil {
				log.Fatal(err)
			}

			log.Println("Total user(s):", len(resp))
			time.Sleep(10 * time.Second)
		}
	}
}

func (b *ListUsersWorker) GetWorkerName() string {
	return b.name
}

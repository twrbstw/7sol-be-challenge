package ports

import (
	"context"
	"sync"
)

type IWorkers interface {
	Start(ctx context.Context, wg *sync.WaitGroup)
	GetWorkerName() string
}

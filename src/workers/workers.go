package workers

import "context"

type IWorkers interface {
	Start(ctx context.Context)
	GetWorkerName() string
}

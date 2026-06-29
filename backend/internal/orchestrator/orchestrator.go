package orchestrator

import (
	"context"
)
type Scheduler interface {
    Run(ctx context.Context) error
    StartRound(ctx context.Context) error
    CloseRound(ctx context.Context) error
    ResolveRound(ctx context.Context) error
}



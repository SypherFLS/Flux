package scheduler

import (
	"context"
)
type Engine interface {
    OpenRound(ctx context.Context) error
    CloseRound(ctx context.Context) error
    ResolveRound(ctx context.Context) error
}
type V1 struct {

}
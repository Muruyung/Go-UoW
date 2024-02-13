package gouow

import (
	"context"
)

type UnitOfWorkInterface interface {
	NewSession(ctx *context.Context) error
	BeginTx(ctx context.Context, operation func(ctxTx context.Context) error) error
	Commit() error
	Rollback(err error) error
}

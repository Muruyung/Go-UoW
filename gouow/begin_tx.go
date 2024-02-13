package gouow

import (
	"context"

	"github.com/Muruyung/Go-UoW/logger"
)

// BeginTx begin unit of work transaction
func (uow *unitOfWorkInteractor) BeginTx(ctx context.Context, operation func(context.Context) error) error {
	logger.Info("Begin transaction process...")

	var (
		ctxVal    = ctx.Value(TX_KEY)
		err       error
		isSession = false
	)

	if ctxVal != nil {
		uow.dbTx = ctxVal.(*TX)
		isSession = uow.dbTx.UseTx
	}

	if !isSession {
		err = uow.NewSession(&ctx)
		if err != nil {
			return err
		}
	}

	err = operation(ctx)
	if err != nil {
		errRollback := uow.Rollback(err)
		if errRollback != nil {
			return errRollback
		}
		return err
	}

	if isSession {
		return nil
	}

	return uow.Commit()
}

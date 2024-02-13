package gouow

import (
	"context"

	"github.com/Muruyung/Go-UoW/logger"
)

type unitOfWorkInteractor struct {
	db   interface{}
	dbTx *TX
	ctx  context.Context
}

// Init initialize new unit of work package
func Init(db interface{}) UnitOfWorkInterface {
	logger.Info("Initialize unit of work")
	return &unitOfWorkInteractor{
		db:  db,
		ctx: context.TODO(),
	}
}

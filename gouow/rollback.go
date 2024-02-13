package gouow

import (
	"database/sql"

	"gorm.io/gorm"

	"github.com/Muruyung/Go-UoW/logger"
)

// Rollback function for rollback all transaction, it used when an error occurs in your business process
func (uow *unitOfWorkInteractor) Rollback(err error) error {
	logger.Info("Transaction rollback process...")

	returnRollback := func() error {
		logger.Info("Transaction rollback success")
		return nil
	}

	switch dbTx := uow.dbTx.Tx.(type) {
	case *gorm.DB:
		err = dbTx.Rollback().Error
		uow.dbTx = nil
		if err == nil || err != gorm.ErrInvalidTransaction {
			return returnRollback()
		}

	case *sql.Tx:
		err = dbTx.Rollback()
		uow.dbTx = nil
		if err == nil || err == sql.ErrTxDone || err == sql.ErrConnDone {
			return returnRollback()
		}
	}

	logger.Error("Transaction rollback failed")
	return err
}

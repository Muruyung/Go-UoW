package gouow

import (
	"database/sql"

	"gorm.io/gorm"

	"github.com/Muruyung/Go-UoW/logger"
)

// Commit function for commit transaction process
func (uow *unitOfWorkInteractor) Commit() error {
	logger.Info("Transaction commit process...")

	errCommit := func(err error) error {
		errRollback := uow.Rollback(err)
		if errRollback != nil {
			return errRollback
		}
		return err
	}

	switch dbTx := uow.dbTx.Tx.(type) {
	case *gorm.DB:
		uow.dbTx.Tx = dbTx.Commit()
		if err := uow.dbTx.Tx.(*gorm.DB).Error; err != nil {
			return errCommit(err)
		}

	case *sql.Tx:
		err := dbTx.Commit()
		if err != nil {
			return errCommit(err)
		}
	}
	uow.dbTx.UseTx = false

	logger.Info("Transaction commit success")

	return nil
}

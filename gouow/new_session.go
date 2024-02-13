package gouow

import (
	"context"
	"database/sql"
	"fmt"

	"gorm.io/gorm"

	"github.com/Muruyung/Go-UoW/logger"
)

// NewSession create new session for start unit of work transaction
func (uow *unitOfWorkInteractor) NewSession(ctx *context.Context) error {
	logger.Info("Create new session process...")
	uow.dbTx = new(TX)
	errSession := func(err error) error {
		msg := fmt.Sprintf("Create new session failed: %v", err)
		logger.Error(msg)
		return err
	}

	switch db := uow.db.(type) {
	case *gorm.DB:
		uow.dbTx.Tx = db.Session(&gorm.Session{SkipDefaultTransaction: true}).Begin()
		if err := uow.dbTx.Tx.(*gorm.DB).Error; err != nil {
			return errSession(err)
		}

	case *sql.DB:
		tx, err := db.Begin()
		if err != nil {
			return errSession(err)
		}
		uow.dbTx.Tx = tx
	}

	uow.dbTx.UseTx = true
	*ctx = context.WithValue(*ctx, TX_KEY, uow.dbTx)

	return nil
}

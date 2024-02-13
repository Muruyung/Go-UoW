package gouow

import (
	"database/sql"

	"gorm.io/gorm"
)

// GormDB get gorm db engine
func (tx *TX) GormDB() *gorm.DB {
	return tx.Tx.(*gorm.DB)
}

// SqlDB get sql db engine
func (tx *TX) SqlDB() *sql.Tx {
	return tx.Tx.(*sql.Tx)
}

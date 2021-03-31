package repository

import (
	"gorm.io/gorm"
)

//开启事务
func BeginTransaction(db *gorm.DB, process func(tx *gorm.DB) error) error {
	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	err := process(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

package ormx

import (
	"database/sql"
	"fmt"
	"github.com/changebooks/database"
	"github.com/changebooks/log"
	"github.com/changebooks/orm"
	"time"
)

// id is primary key's value
func (x *Orm) Delete(idRegister *log.IdRegister,
	driver *database.Driver, table string, id interface{}) (affectedRows int64, err error) {
	tag := "Delete"

	var db *sql.DB
	if driver != nil {
		db = driver.GetDb()
	}

	start := time.Now()

	affectedRows, err, affectedRowsErr, query := orm.Delete(db, table, id)

	done := time.Now()
	remark := NewDeleteRemark(driver, start, done, table, id, query)

	if affectedRowsErr != nil {
		x.logger.E(tag, AffectedRowsFailure, remark, affectedRowsErr, "", idRegister)
	}

	if err == nil {
		x.logger.I(tag, fmt.Sprintf("affected's rows: %d", affectedRows), remark, idRegister)
	} else {
		x.logger.E(tag, Failure, remark, err, "", idRegister)
	}

	return
}

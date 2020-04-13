package ormx

import (
	"database/sql"
	"fmt"
	"github.com/changebooks/database"
	"github.com/changebooks/log"
	"github.com/changebooks/orm"
	"time"
)

// attributes is a struct or a pointer of struct (eg. struct or &struct)
func (x *Orm) Insert(idRegister *log.IdRegister,
	driver *database.Driver, table string, attributes interface{}) (result sql.Result, err error) {
	tag := "Insert"

	var db *sql.DB
	if driver != nil {
		db = driver.GetDb()
	}

	start := time.Now()

	result, err, query, args := orm.Insert(db, table, attributes)

	done := time.Now()
	remark := NewInsertRemark(driver, start, done, table, attributes, query, args...)

	if err == nil {
		affectedRows, affectedRowsErr := result.RowsAffected()
		if affectedRowsErr != nil {
			x.logger.E(tag, AffectedRowsFailure, remark, affectedRowsErr, "", idRegister)
		}

		x.logger.I(tag, fmt.Sprintf("affected's rows: %d", affectedRows), remark, idRegister)
	} else {
		x.logger.E(tag, Failure, remark, err, "", idRegister)
	}

	return
}

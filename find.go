package ormx

import (
	"database/sql"
	"fmt"
	"github.com/changebooks/database"
	"github.com/changebooks/log"
	"github.com/changebooks/orm"
	"time"
)

// result is a pointer (eg. &[]struct or &[]*struct)
func (x *Orm) Find(idRegister *log.IdRegister,
	driver *database.Driver, result interface{}, query string, args ...interface{}) (affectedRows int64, err error) {
	tag := "Find"

	var db *sql.DB
	if driver != nil {
		db = driver.GetDb()
	}

	start := time.Now()

	affectedRows, err, closeErr := orm.Find(db, result, query, args...)

	done := time.Now()
	remark := NewFindRemark(driver, start, done, query, args...)

	if closeErr != nil {
		x.logger.E(tag, CloseFailure, remark, closeErr, "", idRegister)
	}

	if err == nil {
		x.logger.I(tag, fmt.Sprintf("affected's rows: %d", affectedRows), remark, idRegister)
	} else {
		x.logger.E(tag, Failure, remark, err, "", idRegister)
	}

	return
}

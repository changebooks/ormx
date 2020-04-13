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
// id is primary key's value
func (x *Orm) Update(idRegister *log.IdRegister,
	driver *database.Driver, table string, attributes map[string]interface{}, id interface{}) (affectedRows int64, err error) {
	tag := "Update"

	var db *sql.DB
	if driver != nil {
		db = driver.GetDb()
	}

	start := time.Now()

	affectedRows, err, affectedRowsErr, query, args := orm.Update(db, table, attributes, id)

	done := time.Now()
	remark := NewUpdateRemark(driver, start, done, table, attributes, id, query, args...)

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

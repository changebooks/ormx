package ormx

import (
	"github.com/changebooks/database"
	"github.com/changebooks/databasex"
	"time"
)

type FindRemark struct {
	Dsn     string        `json:"dsn"`
	Query   string        `json:"query"`
	Args    interface{}   `json:"args"`
	Command string        `json:"command"`
	Total   time.Duration `json:"total"`
	Start   time.Time     `json:"start"`
	Done    time.Time     `json:"done"`
}

func NewFindRemark(driver *database.Driver,
	start time.Time, done time.Time, query string, args ...interface{}) *FindRemark {
	dsn := ""
	if driver != nil {
		dsn = driver.GetDsn()
	}

	command := databasex.ReplacePlaceholder(query, args...)
	total := done.Sub(start)

	return &FindRemark{
		Dsn:     dsn,
		Query:   query,
		Args:    args,
		Command: command,
		Total:   total,
		Start:   start,
		Done:    done,
	}
}

type InsertRemark struct {
	Dsn        string        `json:"dsn"`
	Table      string        `json:"table"`
	Attributes interface{}   `json:"attributes"`
	Query      string        `json:"query"`
	Args       interface{}   `json:"args"`
	Command    string        `json:"command"`
	Total      time.Duration `json:"total"`
	Start      time.Time     `json:"start"`
	Done       time.Time     `json:"done"`
}

func NewInsertRemark(driver *database.Driver, start time.Time, done time.Time,
	table string, attributes interface{}, query string, args ...interface{}) *InsertRemark {
	dsn := ""
	if driver != nil {
		dsn = driver.GetDsn()
	}

	command := databasex.ReplacePlaceholder(query, args...)
	total := done.Sub(start)

	return &InsertRemark{
		Dsn:        dsn,
		Table:      table,
		Attributes: attributes,
		Query:      query,
		Args:       args,
		Command:    command,
		Total:      total,
		Start:      start,
		Done:       done,
	}
}

type UpdateRemark struct {
	Dsn        string                 `json:"dsn"`
	Table      string                 `json:"table"`
	Attributes map[string]interface{} `json:"attributes"`
	Id         interface{}            `json:"id"`
	Query      string                 `json:"query"`
	Args       interface{}            `json:"args"`
	Command    string                 `json:"command"`
	Total      time.Duration          `json:"total"`
	Start      time.Time              `json:"start"`
	Done       time.Time              `json:"done"`
}

func NewUpdateRemark(driver *database.Driver, start time.Time, done time.Time,
	table string, attributes map[string]interface{}, id interface{}, query string, args ...interface{}) *UpdateRemark {
	dsn := ""
	if driver != nil {
		dsn = driver.GetDsn()
	}

	command := databasex.ReplacePlaceholder(query, args...)
	total := done.Sub(start)

	return &UpdateRemark{
		Dsn:        dsn,
		Table:      table,
		Attributes: attributes,
		Id:         id,
		Query:      query,
		Args:       args,
		Command:    command,
		Total:      total,
		Start:      start,
		Done:       done,
	}
}

type DeleteRemark struct {
	Dsn     string        `json:"dsn"`
	Table   string        `json:"table"`
	Id      interface{}   `json:"id"`
	Query   string        `json:"query"`
	Command string        `json:"command"`
	Total   time.Duration `json:"total"`
	Start   time.Time     `json:"start"`
	Done    time.Time     `json:"done"`
}

func NewDeleteRemark(driver *database.Driver, start time.Time, done time.Time,
	table string, id interface{}, query string) *DeleteRemark {
	dsn := ""
	if driver != nil {
		dsn = driver.GetDsn()
	}

	command := databasex.ReplacePlaceholder(query, id)
	total := done.Sub(start)

	return &DeleteRemark{
		Dsn:     dsn,
		Table:   table,
		Id:      id,
		Query:   query,
		Command: command,
		Total:   total,
		Start:   start,
		Done:    done,
	}
}

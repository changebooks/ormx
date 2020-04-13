package ormx

import (
	"errors"
	"github.com/changebooks/log"
)

type Orm struct {
	logger *log.Logger
}

func New(logger *log.Logger) (*Orm, error) {
	if logger == nil {
		return nil, errors.New("logger can't be nil")
	}

	return &Orm{logger: logger}, nil
}

func (x *Orm) GetLogger() *log.Logger {
	return x.logger
}

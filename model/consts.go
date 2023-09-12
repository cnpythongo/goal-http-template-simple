package model

import "database/sql/driver"

type UserStatusType string

const (
	INACTIVE UserStatusType = "INACTIVE"
	ACTIVE   UserStatusType = "ACTIVE"
	FREEZE   UserStatusType = "FREEZE"
	REMOVE   UserStatusType = "REMOVE"
)

func (st *UserStatusType) Scan(value interface{}) error {
	*st = UserStatusType(value.([]byte))
	return nil
}

func (st UserStatusType) Value() (driver.Value, error) {
	return string(st), nil
}

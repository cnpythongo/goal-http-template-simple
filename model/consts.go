package model

import "database/sql/driver"

type UserStatusType string

const (
	UserStatusInactive UserStatusType = "INACTIVE"
	UserStatusActive   UserStatusType = "ACTIVE"
	UserStatusFreeze   UserStatusType = "FREEZE"
	UserStatusDelete   UserStatusType = "DELETE"
)

func (st *UserStatusType) Scan(value interface{}) error {
	*st = UserStatusType(value.([]byte))
	return nil
}

func (st UserStatusType) Value() (driver.Value, error) {
	return string(st), nil
}

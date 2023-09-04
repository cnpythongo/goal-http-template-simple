package account

import "database/sql/driver"

type userStatusType string

const (
	INACTIVE userStatusType = "INACTIVE"
	ACTIVE   userStatusType = "ACTIVE"
	FREEZE   userStatusType = "FREEZE"
	REMOVE   userStatusType = "REMOVE"
)

func (st *userStatusType) Scan(value interface{}) error {
	*st = userStatusType(value.([]byte))
	return nil
}

func (st userStatusType) Value() (driver.Value, error) {
	return string(st), nil
}

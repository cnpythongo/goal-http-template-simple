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

type SystemConfigScope string

const (
	SystemConfigScopeGlobal SystemConfigScope = "global"
	SystemConfigScopeAdmin  SystemConfigScope = "admin"
	SystemConfigScopeApp    SystemConfigScope = "app"
)

func (st *SystemConfigScope) Scan(value interface{}) error {
	*st = SystemConfigScope(value.([]byte))
	return nil
}

func (st SystemConfigScope) Value() (driver.Value, error) {
	return string(st), nil
}

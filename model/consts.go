package model

import "database/sql/driver"

type UserStatusType string
type UserGender int64

const (
	UserStatusInactive UserStatusType = "INACTIVE"
	UserStatusActive   UserStatusType = "ACTIVE"
	UserStatusFreeze   UserStatusType = "FREEZE"
	UserStatusDelete   UserStatusType = "DELETE"

	UserGenderMale    UserGender = 1
	UserGenderFemale  UserGender = 2
	UserGenderUnknown UserGender = 3
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

type AccountHistoryDevice string

const (
	AccountHistoryDeviceWeb     AccountHistoryDevice = "web"
	AccountHistoryDeviceAndroid AccountHistoryDevice = "android"
	AccountHistoryDeviceIOS     AccountHistoryDevice = "ios"
)

func (st *AccountHistoryDevice) Scan(value interface{}) error {
	*st = AccountHistoryDevice(value.([]byte))
	return nil
}

func (st AccountHistoryDevice) Value() (driver.Value, error) {
	return string(st), nil
}

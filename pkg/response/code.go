package response

import (
	"github.com/cnpythongo/goal/pkg/config"
)

const (
	SuccessCode = 0
	FailCode    = 1

	UnknownError           = 1000
	PayloadError           = 1001
	ParamsError            = 1002
	AuthError              = 1003
	AuthTokenError         = 1004
	AuthTokenGenerateError = 1005
	AuthForbiddenError     = 1006
	AuthLoginRequireError  = 1007

	QueryError            = 1100
	DBAttributesCopyError = 1101
	DeleteError           = 1102
	UpdateError           = 1103
	CreateError           = 1104
	DataExistError        = 1105
	DataNotExistError     = 1106

	AccountEmailExistsError  = 1200
	AccountUserOrPwdError    = 1201
	AccountUserNotExistError = 1202
	AccountUserInactiveError = 1203
	AccountUserFreezeError   = 1204
)

var MsgMapping = map[string]map[int]string{
	"en":    MessageEn,
	"zh_cn": MessageZHCN,
}

func GetCodeMsg(code int) string {
	lang := config.GetConfig().App.Language
	mapping, ok := MsgMapping[lang]
	if !ok {
		return "unsupport language"
	}
	msg, ok := mapping[code]
	if !ok {
		return "unknown error"
	}
	return msg
}

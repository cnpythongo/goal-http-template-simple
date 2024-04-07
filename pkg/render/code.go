package render

import (
	"goal-app/pkg/config"
)

const (
	OK    = 0
	Error = 1

	UnknownError           = 1000
	PayloadError           = 1001
	ParamsError            = 1002
	AuthError              = 1003
	AuthTokenError         = 1004
	AuthTokenGenerateError = 1005
	AuthForbiddenError     = 1006
	AuthLoginRequireError  = 1007
	AuthCaptchaError       = 1008

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
	AccountOldPasswordError  = 1205

	FlixCreditBalanceError       = 2206
	FlixCreditBalanceReduceError = 2207

	UploadFileError = 3000
)

var MsgMapping = map[string]map[int]string{
	"en":    MessageEn,
	"zh_cn": MessageZHCN,
}

func GetCodeMsg(code int, other interface{}) string {
	lang := config.GetConfig().App.Language
	mapping, ok := MsgMapping[lang]
	if !ok {
		return "language key not exists"
	}
	msg, ok := mapping[code]
	if !ok {
		if other == nil {
			return "unknown error"
		} else {
			return other.(string)
		}
	}
	return msg
}

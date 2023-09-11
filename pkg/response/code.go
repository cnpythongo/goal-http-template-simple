package response

import (
	"github.com/cnpythongo/goal/pkg/config"
)

const (
	SuccessCode = 0
	FailCode    = 1

	UnknownError           = 1000
	PayloadError           = 1100
	ParamsError            = 1200
	AuthError              = 1300
	AuthTokenError         = 1301
	AuthTokenGenerateError = 1302
	AuthForbiddenError     = 1303
	AuthRequireError       = 1304
	DBQueryError           = 1400
	DBAttributesCopyError  = 1401

	AccountUserExistError      = 2000
	AccountEmailExistsError    = 2001
	AccountCreateError         = 2002
	AccountUserOrPwdError      = 2003
	AccountUserNotExistError   = 2004
	AccountQueryUserError      = 2005
	AccountQueryUserParamError = 2006
	AccountQueryUserListError  = 2007
	AccountUserInactiveError   = 2008
	AccountUserFreezeError     = 2009
)

var MsgMapping = map[string]map[int]string{
	"en":    MessageEn,
	"zh_cn": MessageZHCN,
}

func GetCodeMsg(code int) string {
	lang := config.GetConfig().App.Language
	mapping, ok := MsgMapping[lang]
	if !ok {
		return ""
	}
	msg, ok := mapping[code]
	if !ok {
		return ""
	}
	return msg
}

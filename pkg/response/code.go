package response

import (
	"github.com/cnpythongo/goal/pkg/common/config"
)

const (
	SuccessCode = 0
	FailCode    = 1

	UnknownError = 1000
	PayloadError = 1100

	AccountUserExistError      = 2000
	AccountEmailExistsError    = 2001
	AccountCreateError         = 2002
	AccountUserIdError         = 2003
	AccountUserNotExistError   = 2004
	AccountQueryUserError      = 2005
	AccountQueryUserParamError = 2006
	AccountQueryUserListError  = 2007
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

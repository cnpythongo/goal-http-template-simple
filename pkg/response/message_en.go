package response

var MessageEn = map[int]string{
	SuccessCode: "ok",
	FailCode:    "fail",

	UnknownError:          "Unknown error",
	PayloadError:          "Post data error",
	ParamsError:           "Params error",
	AuthError:             "认证失败",
	AuthTokenError:        "Auth token error",
	AuthTokenTimeoutError: "Auth token timeout",

	AccountUserExistError:      "用户名已存在，请换一个",
	AccountEmailExistsError:    "邮箱地址已存在，请换一个",
	AccountCreateError:         "创建用户失败",
	AccountUserIdError:         "用户id不正确",
	AccountUserNotExistError:   "用户不存在",
	AccountQueryUserError:      "查询用户失败",
	AccountQueryUserParamError: "查询用户参数不正确",
	AccountQueryUserListError:  "查询用户列表数据失败",
	AccountUserInactiveError:   "用户账号未激活",
	AccountUserFreezeError:     "用户账号被冻结",
}

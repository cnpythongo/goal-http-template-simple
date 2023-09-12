package response

var MessageEn = map[int]string{
	SuccessCode: "ok",
	FailCode:    "fail",

	UnknownError:       "Unknown error",
	PayloadError:       "Post data error",
	ParamsError:        "Params error",
	AuthError:          "认证失败",
	AuthTokenError:     "Auth token error",
	AuthForbiddenError: "无权访问",
	AuthRequireError:   "请登录后再执行操作",

	AccountUserExistError:      "用户名已存在，请换一个",
	AccountEmailExistsError:    "邮箱地址已存在，请换一个",
	AccountCreateError:         "创建用户失败",
	AccountUserOrPwdError:      "用户名或密码不正确",
	AccountUserNotExistError:   "用户不存在",
	AccountQueryUserError:      "查询用户失败",
	AccountQueryUserParamError: "查询用户参数不正确",
	AccountQueryUserListError:  "查询用户列表数据失败",
	AccountUserInactiveError:   "用户账号未激活",
	AccountUserFreezeError:     "用户账号被冻结",
	AccountUserDeleteError:     "删除用户账号失败",
	AccountUserUpdateError:     "更新用户数据失败",

	DBQueryError: "查询失败",
}

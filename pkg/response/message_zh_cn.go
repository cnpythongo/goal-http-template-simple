package response

var MessageZHCN = map[int]string{
	SuccessCode: "ok",
	FailCode:    "失败",

	UnknownError:          "未知错误",
	PayloadError:          "提交表单数据不正确",
	ParamsError:           "请求参数错误",
	AuthTokenError:        "会话令牌不正确，请重新登录",
	AuthTokenTimeoutError: "会话令牌过期，请重新登录",

	AccountUserExistError:      "用户名已存在，请换一个",
	AccountEmailExistsError:    "邮箱地址已存在，请换一个",
	AccountCreateError:         "创建用户失败",
	AccountUserIdError:         "用户id不正确",
	AccountUserNotExistError:   "用户不存在",
	AccountQueryUserError:      "查询用户失败",
	AccountQueryUserParamError: "查询用户参数不正确",
	AccountQueryUserListError:  "查询用户列表数据失败",
}

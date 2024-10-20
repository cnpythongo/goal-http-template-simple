package render

var MessageEn = map[int]string{
	OK:    "ok",
	Error: "fail",

	UnknownError:           "未知错误",
	PayloadError:           "提交表单数据不正确",
	ParamsError:            "请求参数错误",
	AuthError:              "认证失败",
	AuthTokenError:         "会话令牌不正确或已过期，请重新登录",
	AuthTokenGenerateError: "生成会话令牌失败",
	AuthForbiddenError:     "无权访问",
	AuthLoginRequireError:  "请登录后再执行操作",
	AuthCaptchaError:       "验证码不正确",

	QueryError:                "查询失败",
	DBAttributesCopyError:     "属性操作失败",
	DeleteError:               "删除数据失败",
	UpdateError:               "更新数据失败",
	CreateError:               "创建数据失败",
	DataExistError:            "数据已存在",
	DataNotExistError:         "数据不存在",
	AssociatedDataExistsError: "存在关联数据",
	DataNameExistError:        "名称已存在",

	AccountEmailExistsError:  "邮箱地址已存在",
	AccountUserOrPwdError:    "用户名或密码不正确",
	AccountUserInactiveError: "用户账号未激活",
	AccountUserFreezeError:   "用户账号被冻结",
	AccountOldPasswordError:  "旧密码不正确",
	AccountUserNotExistError: "用户账号不存在",

	FlixCreditBalanceError:       "余额不足",
	FlixCreditBalanceReduceError: "余额扣减失败",
}

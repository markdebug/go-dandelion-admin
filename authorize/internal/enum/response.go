package enum

import error_support "github.com/gly-hub/go-dandelion/error-support"

// 基础错误码
var (
	// SystemUnknownError 系统未知错误
	SystemUnknownError = &error_support.Error{Code: 51001, Msg: "系统未知错误"}
	// DataBaseError 数据库错误
	DataBaseError = &error_support.Error{Code: 51002, Msg: "数据库错误"}
	// RedisError Redis错误
	RedisError = &error_support.Error{Code: 51003, Msg: "Redis错误"}
)

// 业务错误码
var (
	// UserNameOrPasswordError 用户名或密码错误
	UserNameOrPasswordError = &error_support.Error{Code: 52001, Msg: "用户名或密码错误"}
	// CaptchaError 验证码错误
	CaptchaError = &error_support.Error{Code: 52002, Msg: "验证码错误"}
)

package errs

// 系统/中间件错误 5xxxix
var (
	DBError              = NewError(500001, "DB错误")
	LoginExpirationError = NewError(501003, "登录已过期")
	CopierError          = NewError(500204, "模型转换错误")
)

// BlogNotExistError 业务错误 2xxxix
var (
	BlogNotExistError = NewError(203001, "博客不存在")
)

package errs

// 系统/中间件错误 5xxxix
var (
	DBError              = NewError(500001, "DB错误")
	LoginExpirationError = NewError(500002, "登录已过期")
	CopierError          = NewError(500003, "模型转换错误")
)

// BlogNotExistError 业务错误 2xxxix
var (
	BlogNotExistError        = NewError(200001, "博客不存在")
	WorkerIdExcessOfQuantity = NewError(200002, "ID超过限制")
	ClockMovedBackwards      = NewError(200003, "系统时钟出现错误")
)

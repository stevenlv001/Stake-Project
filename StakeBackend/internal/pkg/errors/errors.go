package errors

import "fmt"

// ErrorCode 错误码类型
type ErrorCode string

const (
	// 通用错误
	ErrInvalidRequest ErrorCode = "INVALID_REQUEST"
	ErrInternalError  ErrorCode = "INTERNAL_ERROR"
	ErrUnauthorized   ErrorCode = "UNAUTHORIZED"
	ErrForbidden      ErrorCode = "FORBIDDEN"
	ErrNotFound       ErrorCode = "NOT_FOUND"
	ErrTooManyRequests ErrorCode = "TOO_MANY_REQUESTS"

	// 业务错误 - 质押相关
	ErrInvalidAmount      ErrorCode = "INVALID_AMOUNT"
	ErrInsufficientBalance ErrorCode = "INSUFFICIENT_BALANCE"
	ErrBlacklisted        ErrorCode = "BLACKLISTED"
	ErrContractPaused     ErrorCode = "CONTRACT_PAUSED"
	ErrStakeLimitExceeded ErrorCode = "STAKE_LIMIT_EXCEEDED"

	// 业务错误 - 交易相关
	ErrTxNotFound    ErrorCode = "TX_NOT_FOUND"
	ErrTxPending     ErrorCode = "TX_PENDING"
	ErrTxFailed      ErrorCode = "TX_FAILED"
	ErrTxTimeout     ErrorCode = "TX_TIMEOUT"

	// 业务错误 - 管理员相关
	ErrInvalidAdminRole ErrorCode = "INVALID_ADMIN_ROLE"
	ErrAdminOnly        ErrorCode = "ADMIN_ONLY"
)

// AppError 应用错误结构
type AppError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	Details string    `json:"details,omitempty"`
}

func (e *AppError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("[%s] %s: %s", e.Code, e.Message, e.Details)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// New 创建新的应用错误
func New(code ErrorCode, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

// WithDetails 添加详细信息
func (e *AppError) WithDetails(details string) *AppError {
	e.Details = details
	return e
}

// 预定义的错误实例
var (
	// 通用错误
	ErrBadRequest     = New(ErrInvalidRequest, "请求参数错误")
	ErrInternalServer = New(ErrInternalError, "服务器内部错误")
	ErrUnauthenticated = New(ErrUnauthorized, "未授权访问")
	ErrPermissionDenied = New(ErrForbidden, "权限不足")
	ErrResourceNotFound = New(ErrNotFound, "资源不存在")
	ErrRateLimited    = New(ErrTooManyRequests, "请求过于频繁")

	// 质押错误
	ErrInvalidStakeAmount = New(ErrInvalidAmount, "质押金额格式错误或超出范围")
	ErrInsufficientStakeBalance = New(ErrInsufficientBalance, "质押余额不足")
	ErrAddressBlacklisted = New(ErrBlacklisted, "地址已被列入黑名单")
	ErrContractIsPaused = New(ErrContractPaused, "合约已暂停")
	ErrExceedsStakeLimit = New(ErrStakeLimitExceeded, "超过质押限额")

	// 交易错误
	ErrTransactionNotFound = New(ErrTxNotFound, "交易未被追踪")
	ErrTransactionPending = New(ErrTxPending, "交易等待确认中")
	ErrTransactionFailed = New(ErrTxFailed, "交易执行失败")
	ErrTransactionTimeout = New(ErrTxTimeout, "交易确认超时")

	// 管理员错误
	ErrInvalidRole = New(ErrInvalidAdminRole, "无效的管理员角色")
	ErrAdminAccessRequired = New(ErrAdminOnly, "需要管理员权限")
)

// Wrap 包装底层错误
func Wrap(err error, code ErrorCode, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Details: err.Error(),
	}
}

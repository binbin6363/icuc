package err

import "fmt"

// icuc定义的错误
type ICIUError struct {
	Code int    // 错误码
	Msg  string // 错误消息
}

// 通用错误码定义 100 - 10000
const (
	CodeSystem   = 100  // CodeSystem 系统错误
	CodeUnknown  = 101  // CodeUnknown 未知错误
	CodeNoAuth   = 1000 // CodeNoAuth 缺少认证鉴权信息
	CodeAuthFail = 1001 // CodeAuthFail 认证鉴权信息失败
)

// 错误定义
var (
	ErrSystem  = New(CodeSystem, "系统错误")
	ErrUnknown = New(CodeUnknown, "未知错误")

	ErrNoAuth   = New(CodeNoAuth, "没有认证信息")
	ErrAuthFail = New(CodeAuthFail, "认证信息鉴权失败")
)

func (ie *ICIUError) Error() string {
	return fmt.Sprintf("ICIU Code: %d, Msg: %s", ie.Code, ie.Msg)
}

// New 新建ICIUError
func New(code int, msg string) *ICIUError {
	return &ICIUError{Code: code, Msg: msg}
}

// Code 从error获取错误码
func Code(e error) int {
	if e == nil {
		return 0
	}
	if ie, ok := e.(*ICIUError); ok {
		return ie.Code
	}

	return CodeUnknown
}

// Msg 从error获取错误消息
func Msg(e error) string {
	if e == nil {
		return ""
	}
	if ie, ok := e.(*ICIUError); ok {
		return ie.Msg
	}
	return Msg(ErrUnknown)
}

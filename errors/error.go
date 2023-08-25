package errors

import (
	"errors"
	"fmt"
	"io"

	stderrors "github.com/pkg/errors"
)

const (
	defaultErrorCode = "999999"
)

type Error interface {
	Wraped() []error
	Code() string

	error

	private()
}

var _ Error = &CodeError{}

type CodeError struct {
	cause   error  // 原始错误信息
	code    string // 错误码
	message string // 错误消息
}

func (e CodeError) Error() string {
	if len(e.code) == 0 || e.code == defaultErrorCode {
		return e.message
	}

	if len(e.message) == 0 {
		return fmt.Sprint("err:", e.code)
	}

	return fmt.Sprintf("code=%s message=%s", e.code, e.message)
}

func (e *CodeError) Cause() error { return e.cause }

func (e *CodeError) Code() string { return e.code }

func (e *CodeError) Message() string { return e.message }

func (e *CodeError) Unwrap() error { return stderrors.Unwrap(e.cause) }

func (e *CodeError) Is(err error) bool {
	if se, ok := err.(*CodeError); ok {
		return se.code == e.code
	}

	return false
}

func (e *CodeError) As(any) bool {
	return false
}

func (e *CodeError) WithMessage(args ...any) *CodeError {
	return new(e.cause, e.code, fmt.Sprint(args...))
}

func (e *CodeError) WithMessagef(format string, args ...any) *CodeError {
	return new(e.cause, e.code, fmt.Sprintf(format, args...))
}

func (e *CodeError) WithCode(code string) *CodeError {
	return new(e.cause, code, e.message)
}

func (e *CodeError) WithStack(err error) *CodeError {
	return new(stderrors.WithStack(err), e.code, e.message)
}

func (e *CodeError) WithCause(err error) *CodeError {
	return new(err, e.code, e.message)
}

func (e *CodeError) Wraped() []error {
	return []error{}
}

func (e *CodeError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			if e.cause != nil {
				fmt.Fprintf(s, "%+v\n", e.Cause())
			}
			io.WriteString(s, e.message)
			return
		}
		fallthrough
	case 's', 'q':
		io.WriteString(s, e.Error())
	}
}

func (*CodeError) private() {}

func new(err error, code string, args ...any) *CodeError {
	return &CodeError{
		code:    code,
		message: fmt.Sprint(args...),
		cause:   err,
	}
}

func newf(err error, code string, format string, args ...any) *CodeError {
	return &CodeError{
		code:    code,
		message: fmt.Sprintf(format, args...),
		cause:   err,
	}
}

func ErrorWithCode(code string, args ...any) *CodeError {
	return new(nil, code, args...)
}

func Errorf(code string, format string, args ...any) *CodeError {
	return newf(nil, code, format, args...)
}

func WithStack(err error) error {
	return new(stderrors.WithStack(err), defaultErrorCode)
}

// 附加消息
func Wrap(err error, code string, args ...any) error {
	return new(stderrors.WithStack(err), code, args...)
}

// 自定义消息并附加堆栈信息
func Wrapf(err error, code string, format string, args ...any) error {
	return newf(stderrors.WithStack(err), code, format, args...)
}

// 原始错误信息
func Cause(err error) error {
	return stderrors.Unwrap(err)
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func As(err error, target any) bool {
	if target == nil {
		panic("errors: target cannot be nil")
	}

	return errors.As(err, target)
}

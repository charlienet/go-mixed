package errors_test

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/charlienet/go-mixed/errors"
	pkgerr "github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

const (
	defaultErrorCode = "999999"
	testCode         = "0093832"
)

var (
	globalError = errors.Error(defaultErrorCode, "全局错误对象")
	errorbyCode = errors.Error(testCode, "全局错误对象")
)

func TestPersetError(t *testing.T) {
	t.Log(globalError)
	t.Log(errorbyCode)
}

func TestWithCode(t *testing.T) {
	code := "098765"
	ne := globalError.WithCode(code)

	assert.Equal(t, code, ne.Code())
	assert.Equal(t, globalError.Code(), defaultErrorCode)
}

func TestWithMessage(t *testing.T) {
	se := errorbyCode.WithMessage("新的消息")
	t.Log(se)
	assert.Equal(t, testCode, se.Code())
}

func TestNewError(t *testing.T) {
	var e error = errors.Error("123456", "测试")

	err := e.(*errors.CodeError)
	t.Log(e, err.Code())
}

func TestNewWrapError(t *testing.T) {
	err := errors.Wrapf(newError(), "33333", "测试")
	t.Logf("%+v", err)
}

func TestWithStack(t *testing.T) {
	err := newError()
	e := errors.Wrapf(err, "888888", "这是附加%s消息", "测试段")
	t.Logf("%+v", e)
}

func TestLogMessage(t *testing.T) {
	t.Logf("%+v", errors.Error("88888", "错误消息"))
	t.Log(errors.Error("77777"))
	t.Log(errors.Error("77777", "测试"))
}

func TestIs(t *testing.T) {
	code1 := "000090"
	code2 := "000091"
	e1 := errors.Error(code1)
	e2 := errors.Error(code1)
	e3 := errors.Error(code2)

	t.Log(errors.Is(e1, e2))
	t.Log(errors.Is(e1, e3))
}

func TestAs(t *testing.T) {
	sss := &errors.CodeError{}
	ret := errors.As(globalError, sss)
	t.Log(ret)
}

func TestCause(t *testing.T) {
	err := newError()
	we := errors.Wrap(err, "22222")

	ue := errors.Cause(we)
	t.Logf("%+v", ue)
}

func TestCaller(t *testing.T) {
	cc := c()
	t.Logf("Caller:%+v", cc)
}

type ErrorString struct {
	s string
}

func (e *ErrorString) Error() string {
	return e.s
}

func TestAs2(t *testing.T) {
	var targetErr *ErrorString
	err := fmt.Errorf("new error:[%w]", &ErrorString{s: "target err"})
	t.Log(errors.As(err, &targetErr))
	t.Log(targetErr)
}

func c() *stack {
	return caller()
}

func caller() *stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(2, pcs[:])
	var st stack = pcs[0:n]
	return &st
}

type stack []uintptr

func (s *stack) Format(st fmt.State, verb rune) {
	switch verb {
	case 'v':
		switch {
		case st.Flag('+'):
			for _, pc := range *s {
				f := pkgerr.Frame(pc)
				fmt.Fprintf(st, "\n%+v", f)
			}
		}
	}
}

func newError() error {
	return fmt.Errorf("原生错误信息")
}

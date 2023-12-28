package runtime

import (
	"bytes"
	"fmt"
	"runtime"
)

func WrapStackError(v interface{}) *StackError {
	return newStackError(v, 3)
}

func WrapStackErrorSkip(v interface{}, skip int) *StackError {
	return newStackError(v, skip)
}

type StackError struct {
	value interface{}
	stack bytes.Buffer
}

var _ error = (*StackError)(nil)

func (p *StackError) Error() string {
	return fmt.Sprintf("%v", p.value)
}

func (p *StackError) Stack() string {
	return p.stack.String()
}

func newStackError(v interface{}, skip int) *StackError {
	stack := make([]uintptr, 50)
	length := runtime.Callers(skip, stack)
	frames := runtime.CallersFrames(stack[:length])

	buf := bytes.Buffer{}
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		buf.WriteString(fmt.Sprintf("%s:%d (0x%x)\n\t%s\n", frame.File, frame.Line, frame.PC, frame.Function))
	}

	return &StackError{value: v, stack: buf}
}

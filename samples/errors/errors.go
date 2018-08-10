package main

import (
	"fmt"
	"path/filepath"
	"runtime"
)

// START_ST_OMIT
type StackTrace []StackFrame

func FromRuntimeFrames(fs *runtime.Frames) StackTrace {
	var frames []StackFrame
	f, more := fs.Next()
	for more {
		frames = append(frames, StackFrame{File: f.File, Line: f.Line, Function: f.Function})
		f, more = fs.Next()
	}
	return frames
}

type StackFrame struct {
	File     string
	Line     int
	Function string
}

func (st StackTrace) String() string {
	out := ""
	for _, f := range st {
		out = out + fmt.Sprintf("%s - %s:%d\n", f.Function, f.File, f.Line)
	}
	return out
}

// END_ST_OMIT

// START_BE_OMIT
type BetterError struct {
	File       string
	Line       int
	StackTrace StackTrace
	Message    string
}

func (be BetterError) Error() string {
	return fmt.Sprintf("%s caused by \n%s ", be.Message, be.StackTrace)
}

func NewBetterError(msg string) BetterError {
	_, path, line, _ := runtime.Caller(1)
	trace := make([]uintptr, 20)
	count := runtime.Callers(1, trace)
	frames := FromRuntimeFrames(runtime.CallersFrames(trace[:count]))
	_, file := filepath.Split(path)
	return BetterError{
		File: file, Line: line, StackTrace: frames, Message: msg,
	}
}

// END_BE_OMIT

// START_MAIN_OMIT
func SayHello() error {
	return NewBetterError("oops")
}

func main() {
	err := SayHello()
	if err != nil {
		panic(err)
	}
}

// END_MAIN_OMIT

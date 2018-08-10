package main

import (
	"fmt"
	"path/filepath"
	"runtime"
)

// START_OMIT
func decorate(message string) string {
	_, path, line, ok := runtime.Caller(1)
	_, file := filepath.Split(path)
	if ok {
		return fmt.Sprintf("[%s:%d] - %s", file, line, message)
	}
	return message
}

func SayName() {
	fmt.Println(decorate("Attempting to say a name"))
}

func main() {
	SayName()
}

// END_OMIT

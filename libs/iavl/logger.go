package iavl

import (
	"fmt"
	"sync"
)

const (
	FlagOutputModules = "iavl-output-modules"
)

const (
	IavlErr   = 0
	IavlInfo  = 1
	IavlDebug = 2
)

var (
	once    sync.Once
	logFunc LogFuncType = nil

	OutputModules map[string]int
)

type LogFuncType func(level int, format string, args ...interface{})

func SetLogFunc(l LogFuncType) {
	once.Do(func() {
		logFunc = l
	})
}

func iavlLog(module string, level int, format string, args ...interface{}) {
	if v, ok := OutputModules[module]; ok && v != 0 && logFunc != nil {
		format = fmt.Sprintf("%s", format)
		logFunc(level, format, args...)
	}
}

// Code generated by depstubber. DO NOT EDIT.
// This is a simple stub for github.com/astaxie/beego/logs, strictly for use in testing.

// See the LICENSE file for information about the licensing of the original library.
// Source: github.com/astaxie/beego/logs (exports: ; functions: NewLogger,Alert,Critical,Debug,Emergency,Error,Info,Informational,Notice,Trace,Warn,Warning)

// Package logs is a stub of github.com/astaxie/beego/logs, generated by depstubber.
package logs

func Alert(_ interface{}, _ ...interface{}) {}

type BeeLogger struct{}

func (_ *BeeLogger) Alert(_ string, _ ...interface{}) {}

func (_ *BeeLogger) Async(_ ...int64) *BeeLogger {
	return nil
}

func (_ *BeeLogger) Close() {}

func (_ *BeeLogger) Critical(_ string, _ ...interface{}) {}

func (_ *BeeLogger) Debug(_ string, _ ...interface{}) {}

func (_ *BeeLogger) DelLogger(_ string) error {
	return nil
}

func (_ *BeeLogger) Emergency(_ string, _ ...interface{}) {}

func (_ *BeeLogger) EnableFuncCallDepth(_ bool) {}

func (_ *BeeLogger) Error(_ string, _ ...interface{}) {}

func (_ *BeeLogger) Flush() {}

func (_ *BeeLogger) GetLevel() int {
	return 0
}

func (_ *BeeLogger) GetLogFuncCallDepth() int {
	return 0
}

func (_ *BeeLogger) Info(_ string, _ ...interface{}) {}

func (_ *BeeLogger) Informational(_ string, _ ...interface{}) {}

func (_ *BeeLogger) Notice(_ string, _ ...interface{}) {}

func (_ *BeeLogger) Reset() {}

func (_ *BeeLogger) SetLevel(_ int) {}

func (_ *BeeLogger) SetLogFuncCallDepth(_ int) {}

func (_ *BeeLogger) SetLogger(_ string, _ ...string) error {
	return nil
}

func (_ *BeeLogger) SetPrefix(_ string) {}

func (_ *BeeLogger) Trace(_ string, _ ...interface{}) {}

func (_ *BeeLogger) Warn(_ string, _ ...interface{}) {}

func (_ *BeeLogger) Warning(_ string, _ ...interface{}) {}

func (_ *BeeLogger) Write(_ []byte) (int, error) {
	return 0, nil
}

func Critical(_ interface{}, _ ...interface{}) {}

func Debug(_ interface{}, _ ...interface{}) {}

func Emergency(_ interface{}, _ ...interface{}) {}

func Error(_ interface{}, _ ...interface{}) {}

func Info(_ interface{}, _ ...interface{}) {}

func Informational(_ interface{}, _ ...interface{}) {}

func NewLogger(_ ...int64) *BeeLogger {
	return nil
}

func Notice(_ interface{}, _ ...interface{}) {}

func Trace(_ interface{}, _ ...interface{}) {}

func Warn(_ interface{}, _ ...interface{}) {}

func Warning(_ interface{}, _ ...interface{}) {}
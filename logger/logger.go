package logger

import "log"

// WeakLogger only has methods Info, Warn and Error
// as it is only used inner the module iter
type WeakLogger interface {
	Info(msg ...any)
	Warn(msg ...any)
	Error(msg ...any)
}

var _defaultLogger WeakLogger = &weakLogger{}

type weakLogger struct{}

func (wl *weakLogger) Info(msg ...any) {
	log.Println(msg...)
}

func (wl *weakLogger) Warn(msg ...any) {
	log.Println(msg...)
}

func (wl *weakLogger) Error(msg ...any) {
	log.Println(msg...)
}

func SetLogger(l WeakLogger) {
	_defaultLogger = l
}

func Info(msg ...any) {
	_defaultLogger.Info(msg)
}

func Warn(msg ...any) {
	_defaultLogger.Warn(msg)
}

func Error(msg ...any) {
	_defaultLogger.Error(msg)
}

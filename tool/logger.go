package tool

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

type logger struct {
	debugLogger *log.Logger
	traceLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	fatalLogger *log.Logger
	panicLogger *log.Logger
}

var (
	DebugHandle io.Writer = os.Stdout
	TraceHandle io.Writer = os.Stdout
	InfoHandle  io.Writer = os.Stdout
	WarnHandle  io.Writer = os.Stdout
	ErrorHandle io.Writer = os.Stderr
	FatalHandle io.Writer = os.Stderr
	PanicHandle io.Writer = os.Stderr
)

var debugOnce sync.Once
var traceOnce sync.Once
var infoOnce sync.Once
var warnOnce sync.Once
var errorOnce sync.Once
var fatalOnce sync.Once
var panicOnce sync.Once

var Logger = &logger{}

func (logger *logger) Debugf(format string, v ...interface{}) {
	debugOnce.Do(func() {
		fmt.Fprintln(DebugHandle, "debugOnce.Do init Debugf")
		logger.debugLogger = log.New(DebugHandle, "[DEBUG]\t", log.LstdFlags|log.Lshortfile)
	})
	logger.debugLogger.Output(2, fmt.Sprintf(format, v...))
}

func (logger *logger) Debug(v ...interface{}) {
	debugOnce.Do(func() {
		fmt.Fprintln(DebugHandle, "debugOnce.Do init Debug")
		logger.debugLogger = log.New(DebugHandle, "[DEBUG]\t", log.LstdFlags|log.Lshortfile)
	})
	logger.debugLogger.Output(2, fmt.Sprintln(v...))
}

func (logger *logger) Tracef(format string, v ...interface{}) {
	traceOnce.Do(func() {
		fmt.Fprintln(TraceHandle, "traceOnce.Do init Tracef")
		logger.traceLogger = log.New(TraceHandle, "[TRACE]\t", log.LstdFlags|log.Lshortfile)
	})
	logger.traceLogger.Output(2, fmt.Sprintf(format, v...))
}

func (logger *logger) Trace(v ...interface{}) {
	traceOnce.Do(func() {
		fmt.Fprintln(TraceHandle, "traceOnce.Do init Trace")
		logger.traceLogger = log.New(TraceHandle, "[TRACE]\t", log.LstdFlags|log.Lshortfile)
	})
	logger.traceLogger.Output(2, fmt.Sprintln(v...))
}

func (logger *logger) Infof(format string, v ...interface{}) {
	infoOnce.Do(func() {
		fmt.Fprintln(InfoHandle, "infoOnce.Do init Infof")
		logger.infoLogger = log.New(InfoHandle, "[INFO]\t", log.LstdFlags|log.Lshortfile)
	})
	logger.infoLogger.Output(2, fmt.Sprintf(format, v...))
}

func (logger *logger) Info(v ...interface{}) {
	infoOnce.Do(func() {
		fmt.Fprintln(InfoHandle, "infoOnce.Do init Info")
		logger.infoLogger = log.New(InfoHandle, "[INFO]\t", log.LstdFlags|log.Lshortfile)
	})
	logger.infoLogger.Output(2, fmt.Sprintln(v...))
}

func (logger *logger) Warningf(format string, v ...interface{}) {
	warnOnce.Do(func() {
		fmt.Fprintln(WarnHandle, "warnOnce.Do init WarningF")
		logger.warnLogger = log.New(WarnHandle, "[WARNING]\t", log.LstdFlags|log.Lshortfile)
	})
	logger.warnLogger.Output(2, fmt.Sprintf(format, v...))
}

func (logger *logger) Warning(v ...interface{}) {
	warnOnce.Do(func() {
		fmt.Fprintln(WarnHandle, "warnOnce.Do init Warning")
		logger.warnLogger = log.New(WarnHandle, "[WARNING]\t", log.LstdFlags|log.Lshortfile)
	})
	logger.warnLogger.Output(2, fmt.Sprintln(v...))
}

func (logger *logger) Errorf(format string, v ...interface{}) {
	errorOnce.Do(func() {
		fmt.Fprintln(ErrorHandle, "errorOnce.Do init Errorf")
		logger.errorLogger = log.New(ErrorHandle, "[ERROR]\t", log.LstdFlags|log.Lshortfile)
	})
	logger.errorLogger.Output(2, fmt.Sprintf(format, v...))
}

func (logger *logger) Error(v ...interface{}) {
	errorOnce.Do(func() {
		fmt.Fprintln(ErrorHandle, "errorOnce.Do init Error")
		logger.errorLogger = log.New(ErrorHandle, "[ERROR]\t", log.LstdFlags|log.Lshortfile)
	})
	logger.errorLogger.Output(2, fmt.Sprintln(v...))
}

func (logger *logger) Fatalf(format string, v ...interface{}) {
	fatalOnce.Do(func() {
		fmt.Fprintln(FatalHandle, "fatalOnce.Do init Fatalf")
		logger.fatalLogger = log.New(FatalHandle, "[FATAL]\t", log.LstdFlags|log.Lshortfile)
	})
	logger.fatalLogger.Output(2, fmt.Sprintf(format, v...))
}

func (logger *logger) Fatal(v ...interface{}) {
	fatalOnce.Do(func() {
		fmt.Fprintln(FatalHandle, "fatalOnce.Do init Fatal")
		logger.fatalLogger = log.New(FatalHandle, "[FATAL]\t", log.LstdFlags|log.Lshortfile)
	})
	logger.fatalLogger.Output(2, fmt.Sprintln(v...))
}

func (logger *logger) Panicf(format string, v ...interface{}) {
	panicOnce.Do(func() {
		fmt.Fprintln(PanicHandle, "panicOnce.Do init Panicf")
		logger.panicLogger = log.New(PanicHandle, "[PANIC]\t", log.LstdFlags|log.Lshortfile)
	})
	logger.panicLogger.Output(2, fmt.Sprintf(format, v...))
}

func (logger *logger) Panic(v ...interface{}) {
	panicOnce.Do(func() {
		fmt.Fprintln(PanicHandle, "panicOnce.Do init Panic")
		logger.panicLogger = log.New(PanicHandle, "[PANIC]\t", log.LstdFlags|log.Lshortfile)
	})
	logger.panicLogger.Output(2, fmt.Sprintln(v...))
}

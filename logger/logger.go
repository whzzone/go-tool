package logger

import (
	"fmt"
	"github.com/fatih/color"
	"time"
)

/*
	自定义颜色
	customLogger := logger.NewLogger()
	customLogger.infoColor = color.New(color.FgMagenta)
	customLogger.warnColor = color.New(color.FgBlue)
	customLogger.errColor = color.New(color.FgHiRed)
	customLogger.msgColor = color.New(color.FgHiGreen)

	customLogger.Info("custom info message")
	customLogger.Warn("custom warn message")
	customLogger.Error("custom error message")
	customLogger.Msg("custom msg message")
*/

// Logger 是一个日志记录器
type Logger struct {
	infoColor  *color.Color // INFO 级别的颜色
	warnColor  *color.Color // WARN 级别的颜色
	errColor   *color.Color // ERROR 级别的颜色
	debugColor *color.Color // DEBUG 级别的颜色
}

// DefLogger 默认
var DefLogger = NewLogger()

// NewLogger 返回一个 logger 实例
func NewLogger() *Logger {
	l := Logger{}

	// 创建四种颜色
	l.infoColor = color.New(color.FgGreen)
	l.warnColor = color.New(color.FgYellow)
	l.errColor = color.New(color.FgRed)
	l.debugColor = color.New(color.FgWhite)

	return &l
}

// Info 输出 INFO 级别的日志
func (l *Logger) Info(msg string) {
	l.infoColor.Printf("[%s] %s\n", time.Now().Format("2006-01-02 15:04:05"), msg)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.infoColor.Printf("[%s] %s", time.Now().Format("2006-01-02 15:04:05"), fmt.Sprintf(format, v...))
}

// Warn 输出 WARN 级别的日志
func (l *Logger) Warn(msg string) {
	l.warnColor.Printf("[%s] %s\n", time.Now().Format("2006-01-02 15:04:05"), msg)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.warnColor.Printf("[%s] %s", time.Now().Format("2006-01-02 15:04:05"), fmt.Sprintf(format, v...))
}

// Error 输出 ERROR 级别的日志
func (l *Logger) Error(msg string) {
	l.errColor.Printf("[%s] %s\n", time.Now().Format("2006-01-02 15:04:05"), msg)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.errColor.Printf("[%s] %s", time.Now().Format("2006-01-02 15:04:05"), fmt.Sprintf(format, v...))
}

// Debug 输出 DEBUG 级别的日志
func (l *Logger) Debug(msg string) {
	l.debugColor.Printf(fmt.Sprintf("[%s] %s\n", time.Now().Format("2006-01-02 15:04:05"), msg))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debugColor.Printf(fmt.Sprintf("[%s] %s", time.Now().Format("2006-01-02 15:04:05"), fmt.Sprintf(format, v...)))
}

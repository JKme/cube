package gologger

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelError:
		return "ERROR"
	}
	return ""
}

var (
	formatLogger *Logger
	levelMap     = map[string]Level{
		"DEBUG": LevelDebug,
		"INFO":  LevelInfo,
		"ERROR": LevelError,
	}
)

func init() {
	InitLog("INFO")

}

func InitLog(l string) {
	switch {
	case l == "INFO":
		formatLogger = newLog(os.Stdout).setLevel(levelMap[l])
	default:
		formatLogger = newLog(os.Stdout).setFlags(log.Ldate | log.Ltime | log.Lshortfile)
	}
}

type Logger struct {
	level Level
	l     *log.Logger
}

func newLog(w io.Writer) *Logger {
	return &Logger{
		l: log.New(w, "", 0),
	}
}

func (l *Logger) setFlags(flag int) *Logger {
	l.l.SetFlags(flag)
	return l
}

func (l *Logger) setLevel(level Level) *Logger {
	l.level = level
	return l
}

func (l *Logger) doLog(level Level, v ...interface{}) bool {
	if level < l.level {
		return false
	}
	if level == 1 {
		l.l.Output(3, fmt.Sprintln(fmt.Sprintln(v...)))
	} else {
		l.l.Output(3, "\t"+level.String()+"\t"+fmt.Sprintln(fmt.Sprintln(v...)))

	}
	return true
}

func (l *Logger) doLogf(level Level, format string, v ...interface{}) bool {
	if level < l.level {
		return false
	}
	if level == 1 {
		l.l.Output(3, fmt.Sprintln(fmt.Sprintf(format, v...)))
	} else {
		l.l.Output(3, "\t"+level.String()+"\t"+fmt.Sprintln(fmt.Sprintf(format, v...)))

	}
	return true
}

func Debug(v ...interface{}) {
	formatLogger.doLog(LevelDebug, v...)
}

func Info(v ...interface{}) {
	formatLogger.doLog(LevelInfo, v...)
}

func Warn(v ...interface{}) {
	formatLogger.doLog(LevelWarn, v...)
}

func Error(v ...interface{}) {
	formatLogger.doLog(LevelError, v...)
	os.Exit(1)
}

func Errorf(format string, v ...interface{}) {
	formatLogger.doLogf(LevelError, format, v...)
	os.Exit(1)
}

func Warnf(format string, v ...interface{}) {
	formatLogger.doLogf(LevelWarn, format, v...)
}

func Debugf(format string, v ...interface{}) {
	formatLogger.doLogf(LevelDebug, format, v...)
}

func Infof(format string, v ...interface{}) {
	formatLogger.doLogf(LevelInfo, format, v...)
}

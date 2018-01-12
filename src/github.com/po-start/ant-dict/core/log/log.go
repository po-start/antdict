package log

import (
	"os"
	//"strings"

	golog "github.com/siddontang/go/log"

	_ "github.com/po-start/ant-dict/core/errors"
)

var consoleLogger *golog.Logger = StdLogger()
var fileLogger *golog.Logger = StdLogger()

// 全局变量
var GlobalSysLogger *golog.Logger = consoleLogger

// 创建日志文件
func New(fileName string, maxBytes int, backupCount int) error {
	sysFile, err := golog.NewRotatingFileHandler(fileName, maxBytes, backupCount)
	if err != nil {
		return err
	}
	fileLogger = golog.New(sysFile, golog.Lfile|golog.Ltime|golog.Llevel)
	return nil
}

// 开启或关闭写入文件
func IsWriteFile(b bool) {
	if b {
		GlobalSysLogger = fileLogger
	} else {
		GlobalSysLogger = consoleLogger
	}
	return
}

// 开闭文件
func Close() {
	consoleLogger.Close()
	fileLogger.Close()
}

func StdLogger() *golog.Logger {
	h, _ := golog.NewStreamHandler(os.Stdout)
	return golog.NewDefault(h)
}

func SetLogLevel(level string) {
	/*  switch strings.ToLower(level) {
	    case "debug":
	        golog.SetLevel(golog.LevelDebug)
	    case "info":
	        golog.SetLevel(golog.LevelInfo)
	    case "warn":
	        golog.SetLevel(golog.LevelWarn)
	    case "error":
	        golog.SetLevel(golog.LevelError)
	    default:
	        golog.SetLevel(golog.LevelError)
	    }*/
	golog.SetLevelByName(level)
}

func Trace(i ...interface{}) {
	GlobalSysLogger.Trace(i)
}

func Debug(i ...interface{}) {
	GlobalSysLogger.Debug(i)
}

func Info(i ...interface{}) {
	GlobalSysLogger.Info(i)
}

func Warn(i ...interface{}) {
	GlobalSysLogger.Warn(i)
}

func Error(i ...interface{}) {
	GlobalSysLogger.Error(i)
}

func Fatal(i ...interface{}) {
	GlobalSysLogger.Fatal(i)
}

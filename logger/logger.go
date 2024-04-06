package logger

import (
	"BlogServ/config"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	BASEDIR = "/etc/blogserv"

	DEBUG = "DEBUG"
	INFO  = "INFO"
	ERROR = "ERROR"
)

type logger struct {
	fp    string
	funcn string
	l     *logrus.Logger
}

var (
	log = &logger{
		l: logrus.New(),
	}
)

func init() {
	// 日志输出路径
	logFile, err := os.OpenFile(BASEDIR+"/server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.l.Errorf("无法打开日志文件：%v", err)
	}
	//defer func(logFile *os.File) {
	//	err := logFile.Close()
	//	if err != nil {
	//		log.l.Errorf("关闭日志文件失败：%v", err)
	//	}
	//}(logFile)

	writers := io.MultiWriter(logFile, os.Stdout)
	log.l.SetOutput(writers)

	// 日志时间格式
	log.l.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02T15:04:05",
	})

	// 日志等级
	var level logrus.Level
	c := config.AllConfig()
	logLevel := strings.ToUpper(c.Logging.Level)
	switch logLevel {
	case DEBUG:
		level = logrus.DebugLevel
	case INFO:
		level = logrus.InfoLevel
	case ERROR:
		level = logrus.ErrorLevel
	}
	log.l.SetLevel(level)
}

// getCallerInfo 获取调用者的函数名，调用行
func (l *logger) getCallerInfo() {
	// 调用链往上翻三层，找到调用函数的信息
	pc, file, line, ok := runtime.Caller(3)
	if !ok {
		return
	}

	funcName := runtime.FuncForPC(pc).Name()
	l.funcn = path.Base(funcName)
	_, fileName := path.Split(file)
	l.fp = fmt.Sprintf("%s:%d", fileName, line)
}

// printer 日志处理函数
func (l *logger) printer(level logrus.Level, msg ...interface{}) {
	log.getCallerInfo()
	l.l.WithFields(logrus.Fields{"filePath": l.fp, "func": l.funcn}).Log(level, msg...)
}

// printerf 格式化日志处理
func (l *logger) printerf(level logrus.Level, format string, msg ...interface{}) {
	log.getCallerInfo()
	l.l.WithFields(logrus.Fields{"filePath": l.fp, "func": l.funcn}).Logf(level, format, msg...)
}

func Info(msg ...interface{}) {
	log.printer(logrus.InfoLevel, msg...)
}

func Infof(format string, msg ...interface{}) {
	log.printerf(logrus.InfoLevel, format, msg...)
}

func Debug(msg ...interface{}) {
	log.printer(logrus.DebugLevel, msg...)
}

func Debugf(format string, msg ...interface{}) {
	log.printerf(logrus.DebugLevel, format, msg...)
}

func Error(msg ...interface{}) {
	log.printer(logrus.ErrorLevel, msg...)
}

func Errorf(format string, msg ...interface{}) {
	log.printerf(logrus.ErrorLevel, format, msg...)
}

func Fatal(msg ...interface{}) {
	log.printer(logrus.FatalLevel, msg...)
}

func Fatalf(format string, msg ...interface{}) {
	log.printerf(logrus.FatalLevel, format, msg...)
}

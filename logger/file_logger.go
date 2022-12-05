package logger

import (
	"fmt"
	"github.com/CarrotVegeta/showstart/conf"
	"github.com/CarrotVegeta/showstart/pkg/utils"
	"io"
	"log"
	"os"
	"path"
	"time"
)

var FileLog *fileLogger

type fileLogger struct {
	logLevel LogLevel
	logInfo  *os.File
	logError *os.File
	infoFp   string
	errorFp  string
	fn       string
	date     time.Time
}

func NewFileLogger() {
	if conf.Conf.Logger == nil {
		conf.Conf.Logger = conf.NewLogger()
	}
	logLevel, err := parseLogLevel(conf.Conf.Logger.LogLevel)
	if err != nil {
		log.Println(err)
	}
	f := &fileLogger{
		infoFp:   conf.Conf.Logger.InfoFp,
		errorFp:  conf.Conf.Logger.ErrorFn,
		logLevel: logLevel,
	}
	if f.infoFp == "" {
		f.infoFp = "./log/info"
	}
	if f.errorFp == "" {
		f.errorFp = "./log/error"
	}
	if f.fn == "" {
		f.fn = "run.log"
	}
	err = f.createFile()
	if err != nil {
		panic(err.Error())
	}
	FileLog = f
	return
}

func (fl *fileLogger) log(level LogLevel, format string, a ...interface{}) {
	if fl.enable(level) {
		if !fl.isOneDay() {
			err := fl.createFile()
			if err != nil {
				format = err.Error()
			}
		}
		msg := fmt.Sprintf(format, a...)
		funcName, fileName, lineNo := getInfo(3)

		fmt.Fprintf(fl.getFileIO(level), "[%s] [%s] [%s : %s : %d] %s \n", time.Now().Format("2006-01-02 15:04:05 "), getLogString(level), fileName, funcName, lineNo, msg)
	}
}
func (fl *fileLogger) getFileIO(level LogLevel) io.Writer {
	switch level {
	case ERROR:
		return fl.logError
	default:
		return fl.logInfo
	}
}
func (fl *fileLogger) enable(logLevel LogLevel) bool {
	return fl.logLevel <= logLevel
}
func (fl *fileLogger) Info(format string, a ...interface{}) {
	if fl.enable(INFO) {
		fl.log(INFO, format, a...)
	}
}
func (fl *fileLogger) isOneDay() bool {
	return time.Now().Format("2006-01-02") == fl.date.Format("2006-01-02")
}
func (fl *fileLogger) Debug(format string, a ...interface{}) {
	if fl.enable(DEBUG) {
		fl.log(DEBUG, format, a...)
	}
}

func (fl *fileLogger) Trace(format string, a ...interface{}) {
	if fl.enable(TRACE) {
		fl.log(TRACE, format, a...)
	}
}

func (fl *fileLogger) Warning(format string, a ...interface{}) {
	if fl.enable(WARING) {
		fl.log(WARING, format, a...)
	}
}

func (fl *fileLogger) Error(format string, a ...interface{}) {
	if fl.enable(ERROR) {
		fl.log(ERROR, format, a...)
	}
}

func (fl *fileLogger) Fatal(format string, a ...interface{}) {
	if fl.enable(FATAL) {
		fl.log(FATAL, format, a...)
	}

}
func (fl *fileLogger) Close() {
	fl.logError.Close()
	fl.logInfo.Close()
}

func (fl *fileLogger) pathExistsAndCreate(path string) error {
	ok, err := utils.PathExists(path)
	if err != nil {
		return fmt.Errorf("logger path exists :%v", err.Error())
	}
	if !ok {
		err := os.MkdirAll(path, 0644)
		if err != nil {
			return fmt.Errorf("logger create path failed:%v", err.Error())
		}
	}
	return nil
}
func (fl *fileLogger) createFile() error {
	infoFileName := path.Join(fl.infoFp, fl.fn)
	errorFileName := path.Join(fl.errorFp, fl.fn)
	if !fl.isOneDay() {
		if ok, _ := utils.PathExists(infoFileName); ok {
			os.Rename(infoFileName, fmt.Sprintf("%s.log", time.Now().Format("2006-01-02")))
		}
		if ok, _ := utils.PathExists(errorFileName); ok {
			os.Rename(errorFileName+".err", fmt.Sprintf("%s.log.err", time.Now().Format("2006-01-02")))
		}
	}
	if err := fl.pathExistsAndCreate(fl.infoFp); err != nil {
		return fmt.Errorf("infoFp path exists and create:%v", err)
	}
	if err := fl.pathExistsAndCreate(fl.errorFp); err != nil {
		return fmt.Errorf("errorFp path exists and create:%v", err)
	}
	logInfo, err := os.OpenFile(infoFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("new filelogger failed:%v", err.Error())
	}
	logError, err := os.OpenFile(errorFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("new filelogger err file  failed:%v", err.Error())
	}
	fl.logInfo = logInfo
	fl.logError = logError
	fl.date = time.Now()
	return nil
}

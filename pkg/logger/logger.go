package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	loggers = make(map[string]Logger)
	mux     sync.RWMutex
)

type (
	Logger interface {
		Info(bool, ...interface{})
		Infof(bool, string, ...interface{})
		Warn(bool, ...interface{})
		Warnf(bool, string, ...interface{})
		Error(bool, ...interface{})
		Errorf(bool, string, ...interface{})
		Fatal(bool, ...interface{})
		Fatalf(bool, string, ...interface{})
	}
)

// defaultLogger impl the default logger by logrus
type defaultLogger struct {
	cmdLogger     *logrus.Logger
	fileLogger    *logrus.Logger
	name          string
	outputDirPath string
}

func (d *defaultLogger) Info(logToFile bool, msg ...interface{}) {
	d.info(logToFile, strings.Join([]string{d.name, fmt.Sprint(msg...)}, ": "))
}
func (d *defaultLogger) Infof(logToFile bool, format string, msg ...interface{}) {
	d.info(logToFile, strings.Join([]string{d.name, fmt.Sprintf(format, msg...)}, ": "))
}
func (d *defaultLogger) info(logToFile bool, msg string) {
	if logToFile && d.outputDirPath == "" {
		d.cmdLogger.Warn("logger output dir path is \"\", so redirect the logging into cmd")
	}
	if logToFile && d.outputDirPath != "" {
		d.fileLogger.Info(msg)
	} else {
		d.cmdLogger.Info(msg)
	}
}
func (d *defaultLogger) Warn(logToFile bool, msg ...interface{}) {
	d.warn(logToFile, strings.Join([]string{d.name, fmt.Sprint(msg...)}, ": "))
}
func (d *defaultLogger) Warnf(logToFile bool, format string, msg ...interface{}) {
	d.warn(logToFile, strings.Join([]string{d.name, fmt.Sprintf(format, msg...)}, ": "))
}
func (d *defaultLogger) warn(logToFile bool, msg string) {
	if logToFile && d.outputDirPath == "" {
		d.cmdLogger.Warn("logger output dir path is \"\", so redirect the logging into cmd")
	}
	if logToFile && d.outputDirPath != "" {
		d.fileLogger.Warn(msg)
	} else {
		d.cmdLogger.Warn(msg)
	}
}
func (d *defaultLogger) Error(logToFile bool, msg ...interface{}) {
	d.error(logToFile, strings.Join([]string{d.name, fmt.Sprint(msg...)}, ": "))
}
func (d *defaultLogger) Errorf(logToFile bool, format string, msg ...interface{}) {
	d.error(logToFile, strings.Join([]string{d.name, fmt.Sprintf(format, msg...)}, ": "))
}
func (d *defaultLogger) error(logToFile bool, msg string) {
	if logToFile && d.outputDirPath == "" {
		d.cmdLogger.Warn("logger output dir path is \"\", so redirect the logging into cmd")
	}
	if logToFile && d.outputDirPath != "" {
		d.fileLogger.Error(msg)
	} else {
		d.cmdLogger.Error(msg)
	}
}
func (d *defaultLogger) Fatal(logToFile bool, msg ...interface{}) {
	d.fatal(logToFile, strings.Join([]string{d.name, fmt.Sprint(msg...)}, ": "))
}
func (d *defaultLogger) Fatalf(logToFile bool, format string, msg ...interface{}) {
	d.fatal(logToFile, strings.Join([]string{d.name, fmt.Sprintf(format, msg...)}, ": "))
}
func (d *defaultLogger) fatal(logToFile bool, msg string) {
	if logToFile && d.outputDirPath == "" {
		d.cmdLogger.Warn("logger output dir path is \"\", so redirect the logging into cmd")
	}
	if logToFile && d.outputDirPath != "" {
		d.fileLogger.Fatal(msg)
	} else {
		d.cmdLogger.Fatal(msg)
	}
}

func GetLogger(name string, outputDirPath string) Logger {
	mux.Lock()
	if _, ok := loggers[name]; !ok {
		initLogger(name, outputDirPath)
	}
	mux.Unlock()

	mux.RLock()
	defer mux.RUnlock()
	return loggers[name]
}

func initLogger(name string, outputDirPath string) {
	_logger := &defaultLogger{}

	cmdLogr := logrus.New()
	cmdLogr.SetFormatter(&logrus.TextFormatter{})
	cmdLogr.SetOutput(os.Stdout)
	_logger.cmdLogger = cmdLogr

	if outputDirPath != "" {
		fileLogr := logrus.New()
		fileLogr.SetFormatter(&logrus.TextFormatter{})
		timeUnix := time.Now().Unix()
		p, _ := filepath.Abs(outputDirPath)
		_, err := os.Stat(p)
		if os.IsNotExist(err) {
			os.Mkdir(p, os.ModePerm)
		}

		filename := fmt.Sprintf("%s_%s", name, strconv.FormatInt(timeUnix, 10))
		file, err := os.OpenFile(filepath.Join(p, filename+".log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fileLogr.Fatal(err)
		}
		writer := io.Writer(file)
		fileLogr.SetOutput(writer)
		_logger.fileLogger = fileLogr
	}

	_logger.name = name
	loggers[name] = _logger
}

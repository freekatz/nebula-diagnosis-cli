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
		Info(...interface{})
		Infof(string, ...interface{})
		Warn(...interface{})
		Warnf(string, ...interface{})
		Error(...interface{})
		Errorf(string, ...interface{})
		Fatal(...interface{})
		Fatalf(string, ...interface{})
	}
)

// defaultLogger impl the default logger by logrus
type defaultLogger struct {
	cmdLogger     *logrus.Logger
	fileLogger    *logrus.Logger
	name          string
	outputDirPath string
	logToFile     bool
}

func (d *defaultLogger) Info(msg ...interface{}) {
	d.info(strings.Join([]string{d.name, fmt.Sprint(msg...)}, ": "))
}
func (d *defaultLogger) Infof(format string, msg ...interface{}) {
	d.info(strings.Join([]string{d.name, fmt.Sprintf(format, msg...)}, ": "))
}
func (d *defaultLogger) info(msg string) {
	if d.logToFile && d.outputDirPath == "" {
		d.cmdLogger.Warn("logger output dir path is \"\", so redirect the logging into cmd")
	}
	d.cmdLogger.Info(msg)
	if d.logToFile && d.outputDirPath != "" {
		d.fileLogger.Info(msg)
	}
}
func (d *defaultLogger) Warn(msg ...interface{}) {
	d.warn(strings.Join([]string{d.name, fmt.Sprint(msg...)}, ": "))
}
func (d *defaultLogger) Warnf(format string, msg ...interface{}) {
	d.warn(strings.Join([]string{d.name, fmt.Sprintf(format, msg...)}, ": "))
}
func (d *defaultLogger) warn(msg string) {
	if d.logToFile && d.outputDirPath == "" {
		d.cmdLogger.Warn("logger output dir path is \"\", so redirect the logging into cmd")
	}
	d.cmdLogger.Warn(msg)
	if d.logToFile && d.outputDirPath != "" {
		d.fileLogger.Warn(msg)
	}
}
func (d *defaultLogger) Error(msg ...interface{}) {
	d.error(strings.Join([]string{d.name, fmt.Sprint(msg...)}, ": "))
}
func (d *defaultLogger) Errorf(format string, msg ...interface{}) {
	d.error(strings.Join([]string{d.name, fmt.Sprintf(format, msg...)}, ": "))
}
func (d *defaultLogger) error(msg string) {
	if d.logToFile && d.outputDirPath == "" {
		d.cmdLogger.Warn("logger output dir path is \"\", so redirect the logging into cmd")
	}
	d.cmdLogger.Error(msg)
	if d.logToFile && d.outputDirPath != "" {
		d.fileLogger.Error(msg)
	}
}
func (d *defaultLogger) Fatal(msg ...interface{}) {
	d.fatal(strings.Join([]string{d.name, fmt.Sprint(msg...)}, ": "))
}
func (d *defaultLogger) Fatalf(format string, msg ...interface{}) {
	d.fatal(strings.Join([]string{d.name, fmt.Sprintf(format, msg...)}, ": "))
}
func (d *defaultLogger) fatal(msg string) {
	if d.logToFile && d.outputDirPath == "" {
		d.cmdLogger.Warn("logger output dir path is \"\", so redirect the logging into cmd")
	}
	d.cmdLogger.Fatal(msg)
	if d.logToFile && d.outputDirPath != "" {
		d.fileLogger.Fatal(msg)
	}
}

func GetLogger(name string, outputDirPath string, logToFile bool) Logger {
	mux.Lock()
	if _, ok := loggers[name]; !ok {
		initLogger(name, outputDirPath, logToFile)
	}
	mux.Unlock()

	mux.RLock()
	defer mux.RUnlock()
	return loggers[name]
}

func initLogger(name string, outputDirPath string, logToFile bool) {
	_logger := &defaultLogger{}

	cmdLogr := logrus.New()
	cmdLogr.SetFormatter(&logrus.TextFormatter{})
	cmdLogr.SetOutput(os.Stdout)
	_logger.cmdLogger = cmdLogr

	if logToFile {
		fileLogr := logrus.New()
		fileLogr.SetFormatter(&logrus.TextFormatter{})
		timeUnix := time.Now().Unix()
		p, _ := filepath.Abs(outputDirPath)
		_, err := os.Stat(p)
		if os.IsNotExist(err) {
			os.MkdirAll(p, os.ModePerm)
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
	_logger.outputDirPath = outputDirPath
	_logger.logToFile = logToFile
	loggers[name] = _logger
}

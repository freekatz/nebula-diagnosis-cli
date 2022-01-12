package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	cmdLoggers  = make(map[string]*CMDLogger)
	cmdMux      sync.RWMutex
	fileLoggers = make(map[string]*FileLogger)
	fileMux     sync.RWMutex
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
	logr *logrus.Logger
}

func (d *defaultLogger) Info(msg ...interface{}) {
	d.info(fmt.Sprint(msg...))
}
func (d *defaultLogger) Infof(format string, msg ...interface{}) {
	d.info(fmt.Sprintf(format, msg...))
}
func (d *defaultLogger) info(msg string) {
	d.logr.Info(msg)
}
func (d *defaultLogger) Warn(msg ...interface{}) {
	d.warn(fmt.Sprint(msg...))
}
func (d *defaultLogger) Warnf(format string, msg ...interface{}) {
	d.warn(fmt.Sprintf(format, msg...))
}

func (d *defaultLogger) warn(msg string) {
	d.logr.Warn(msg)
}
func (d *defaultLogger) Error(msg ...interface{}) {
	d.error(fmt.Sprint(msg...))
}
func (d *defaultLogger) Errorf(format string, msg ...interface{}) {
	d.error(fmt.Sprintf(format, msg...))
}
func (d *defaultLogger) error(msg string) {
	d.logr.Error(msg)
}
func (d *defaultLogger) Fatal(msg ...interface{}) {
	d.fatal(fmt.Sprint(msg...))
}
func (d *defaultLogger) Fatalf(format string, msg ...interface{}) {
	d.fatal(fmt.Sprintf(format, msg...))
}
func (d *defaultLogger) fatal(msg string) {
	d.logr.Fatal(msg)
}

type CMDLogger struct {
	*defaultLogger
}

func GetCmdLogger(name string) Logger {
	cmdMux.Lock()
	if _, ok := cmdLoggers[name]; !ok {
		initCMDLogger(name)
	}
	cmdMux.Unlock()

	cmdMux.RLock()
	defer cmdMux.RUnlock()
	return cmdLoggers[name]
}

func initCMDLogger(name string) {
	logr := logrus.New()
	logr.SetFormatter(&logrus.TextFormatter{})
	logr.SetOutput(os.Stdout)

	cmdLogger := &CMDLogger{&defaultLogger{}}
	cmdLogger.logr = logr
	cmdLoggers[name] = cmdLogger
}

type FileLogger struct {
	*defaultLogger
}

func GetFileLogger(name string, dirPath string) Logger {
	fileMux.Lock()
	if _, ok := fileLoggers[name]; !ok {
		initFileLogger(name, dirPath)
	}
	fileMux.Unlock()

	fileMux.RLock()
	defer fileMux.RUnlock()
	return fileLoggers[name]
}

func initFileLogger(name string, dirPath string) {
	logr := logrus.New()
	logr.SetFormatter(&logrus.TextFormatter{})

	timeUnix := time.Now().Unix()
	p, _ := filepath.Abs(dirPath)
	_, err := os.Stat(p)
	if os.IsNotExist(err) {
		os.Mkdir(p, os.ModePerm)
	}

	filename := fmt.Sprintf("%s_%s", name, strconv.FormatInt(timeUnix, 10))
	file, err := os.OpenFile(filepath.Join(p, filename+".log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logr.Fatal(err)
	}
	writer := io.Writer(file)

	logr.SetOutput(writer)
	fileLogger := &FileLogger{&defaultLogger{}}
	fileLogger.logr = logr
	fileLoggers[name] = fileLogger
}

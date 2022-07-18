package lib

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/fajardiyanto/flt-go-logger/interfaces"
	"path"
	"runtime"
	"strings"
)

type LogLevel = logrus.Level

type Modules struct {
	ServiceName  string
	Logger       *logrus.Logger
	Entry        *logrus.Entry
	Level        LogLevel
	outputFormat string
}

func NewLib() interfaces.Logger {
	return &Modules{}
}

func (*Modules) New() interfaces.Logger {
	return NewLib()
}

func (c *Modules) Init() interfaces.Logger {
	c.Logger = logrus.New()

	c.Entry = c.Logger.WithFields(logrus.Fields{})

	return c
}

func (c *Modules) SetFormat(format string) interfaces.Logger {
	switch strings.ToLower(format) {
	case "text":
		c.Logger.SetFormatter(&logrus.TextFormatter{
			CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
				_, filename := path.Split(frame.File)
				return "", filename
			},
			FullTimestamp: true,
		})
	case "json":
		c.Logger.SetFormatter(&logrus.JSONFormatter{
			CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
				_, filename := path.Split(frame.File)
				return "", filename
			},
		})
	}
	return c
}

func (c *Modules) SetLevel(level string) interfaces.Logger {
	switch level {
	case "debug":
		c.Logger.SetLevel(logrus.DebugLevel)
		c.Level = logrus.DebugLevel

	case "error":
		c.Logger.SetLevel(logrus.ErrorLevel)
		c.Level = logrus.ErrorLevel

	case "info":
		c.Logger.SetLevel(logrus.InfoLevel)
		c.Level = logrus.InfoLevel

	case "warn":
		c.Logger.SetLevel(logrus.WarnLevel)
		c.Level = logrus.WarnLevel

	default:
		c.Logger.SetLevel(logrus.InfoLevel)
		c.Level = logrus.InfoLevel

	}

	return c
}

func (c *Modules) SetReportCaller(reportCaller bool) {
	c.Logger.SetReportCaller(reportCaller)
}

func (c *Modules) Info(m interface{}, opts ...interface{}) interfaces.Logger {
	fields, ms := interfaces.GetFields(m, opts)
	c.Entry.WithFields(fields).Info(ms...)

	return c
}

func (c *Modules) Infof(f string, opts ...interface{}) interfaces.Logger {
	fields, ms := interfaces.GetFields(nil, opts)
	c.Entry.WithFields(fields).Infof(f, ms...)

	return c
}

func (c *Modules) Warn(m interface{}, opts ...interface{}) interfaces.Logger {
	fields, ms := interfaces.GetFields(m, opts)
	c.Entry.WithFields(fields).Warn(ms...)

	return c
}

func (c *Modules) Warnf(f string, opts ...interface{}) interfaces.Logger {
	fields, ms := interfaces.GetFields(nil, opts)
	c.Entry.WithFields(fields).Warnf(f, ms...)

	return c
}

func (c *Modules) Debug(m interface{}, opts ...interface{}) interfaces.Logger {
	fields, ms := interfaces.GetFields(m, opts)
	c.Entry.WithFields(fields).Debug(ms...)

	return c
}

func (c *Modules) Debugf(f string, opts ...interface{}) interfaces.Logger {
	fields, ms := interfaces.GetFields(nil, opts)
	c.Entry.WithFields(fields).Debugf(f, ms...)

	return c
}

func (c *Modules) Error(m interface{}, opts ...interface{}) interfaces.Logger {
	fields, ms := interfaces.GetFields(m, opts)
	c.Entry.WithFields(fields).Error(ms...)

	return c
}

func (c *Modules) Errorf(f string, opts ...interface{}) interfaces.Logger {
	fields, ms := interfaces.GetFields(nil, opts)
	c.Entry.WithFields(fields).Errorf(f, ms...)

	return c
}

func (c *Modules) Fatal(m interface{}, opts ...interface{}) interfaces.Logger {
	fields, ms := interfaces.GetFields(m, opts)
	c.Entry.WithFields(fields).Fatal(ms...)

	return c
}

func (c *Modules) Fatalf(f string, opts ...interface{}) interfaces.Logger {
	fields, ms := interfaces.GetFields(nil, opts)
	c.Entry.WithFields(fields).Fatalf(f, ms...)

	return c
}

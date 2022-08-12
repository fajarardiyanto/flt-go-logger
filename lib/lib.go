package lib

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fajarardiyanto/flt-go-logger/interfaces"
	"github.com/jwalton/gchalk"
	"io"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"time"
)

type Modules struct {
	namespace string
	level     interfaces.DebugLevel
	writer    io.Writer
}

func NewLib() interfaces.Logger {
	return &Modules{
		level: interfaces.DebugLevelVerbose,
	}
}

func (c *Modules) New() interfaces.Logger {
	return NewLib()
}

func (c *Modules) Init(namespace string) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	log.SetOutput(c)
	c.namespace = namespace
}

func (c *Modules) SetLogLevel(level interfaces.DebugLevel) {
	c.level = level
}

func (c *Modules) GetLogLevel() (level interfaces.DebugLevel) {
	return c.level
}

func (c *Modules) Debug(format interface{}, input ...interface{}) {
	fmt.Println(c.ParsingLog(c.createMsg(interfaces.LogLevelDebug, interfaces.GetCaller(2), format, input)))
}

func (c *Modules) Info(format interface{}, input ...interface{}) {
	fmt.Println(c.ParsingLog(c.createMsg(interfaces.LogLevelInfo, interfaces.GetCaller(2), format, input)))
}

func (c *Modules) Trace(format interface{}, input ...interface{}) {
	fmt.Println(c.ParsingLog(c.createMsg(interfaces.LogLevelTrace, interfaces.GetCaller(2), format, input)))
}

func (c *Modules) Warning(format interface{}, input ...interface{}) {
	fmt.Println(c.ParsingLog(c.createMsg(interfaces.LogLevelWarning, interfaces.GetCaller(2), format, input)))
}

func (c *Modules) Success(format interface{}, input ...interface{}) {
	fmt.Println(c.ParsingLog(c.createMsg(interfaces.LogLevelSuccess, interfaces.GetCaller(2), format, input)))
}

func (c *Modules) Error(format interface{}, input ...interface{}) interfaces.Logger {
	fmt.Println(c.ParsingLog(c.createMsg(interfaces.LogLevelError, interfaces.GetCaller(2), format, input)))
	return c
}

func (c *Modules) createMsg(
	level interfaces.LogLevel,
	caller interfaces.Caller,
	format interface{},
	input ...interface{}) (msg interfaces.LoggerMessage) {

	var inp []interface{}
	for _, s := range input {
		if val, ok := s.([]interface{}); ok {
			inp = append(inp, val...)
		}
	}

	var ffs string
	var msgs interface{}
	if val, ok := format.(string); ok {
		ffs = val
		msgs = fmt.Sprintf(ffs, inp...)
	} else if val, ok := format.(error); ok {
		ffs = val.Error()
		msgs = fmt.Sprintf(ffs, inp...)
	} else {
		msgs = format
	}

	return interfaces.LoggerMessage{
		Message:   msgs,
		LevelName: interfaces.GetLogLevelString(level),
		File:      caller.File,
		Line:      caller.Line,
		FuncName:  caller.Fname,
		Level:     level,
		Time:      time.Now(),
	}
}

func (c *Modules) ParsingLog(msg interfaces.LoggerMessage) (raw string) {
	mm := gchalk.WithBold()
	mmc := gchalk.WithBold()
	var ems string
	var vms string
	mmsg := fmt.Sprintf("%s", msg.Message)

	vv := reflect.TypeOf(msg.Message)
	if vv != nil {
		switch vv.Kind() {
		case reflect.String:
		default:
			if vals, err := json.Marshal(msg.Message); err == nil {
				mmsg = string(vals)
			}
		}
	} else {
		if vals, err := json.Marshal(msg.Message); err == nil {
			mmsg = string(vals)
		}
	}

	switch msg.Level {
	case interfaces.LogLevelTrace:
		ems = mm.BrightWhite(interfaces.GetLogLevelPrintString(msg.Level))
		vms = mmc.White(mmsg)
	case interfaces.LogLevelDebug:
		ems = mm.BrightBlue(interfaces.GetLogLevelPrintString(msg.Level))
		vms = mmc.White(mmsg)
	case interfaces.LogLevelNotice:
		ems = mm.BrightCyan(interfaces.GetLogLevelPrintString(msg.Level))
		vms = mmc.BrightCyan(mmsg)
	case interfaces.LogLevelInfo:
		ems = mm.BrightMagenta(interfaces.GetLogLevelPrintString(msg.Level))
		vms = mmc.BrightMagenta(mmsg)
	case interfaces.LogLevelWarning:
		ems = mm.Yellow(interfaces.GetLogLevelPrintString(msg.Level))
		vms = mmc.Yellow(mmsg)
	case interfaces.LogLevelError:
		ems = mm.Red(interfaces.GetLogLevelPrintString(msg.Level))
		vms = mmc.Red(mmsg)
	case interfaces.LogLevelSuccess:
		ems = mm.Green(interfaces.GetLogLevelPrintString(msg.Level))
		vms = mmc.Green(mmsg)

	}

	raw = fmt.Sprintf("[%s][%s][%s][%d] %s",
		gchalk.Magenta(msg.Time.Format("2006-01-02 15:04:05")),
		ems, gchalk.BrightCyan(filepath.Base(msg.FuncName)),
		msg.Line,
		vms)
	return raw
}

func (c *Modules) Quit() {
	os.Exit(0)
}

func (c *Modules) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewBuffer(p))
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) != 0 {
			c.Debug("%s", text)
		}
	}
	return len(p), nil
}

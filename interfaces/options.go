package interfaces

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

type Caller struct {
	File       string `json:"file"`
	Line       int    `json:"line"`
	Fname      string `json:"fname"`
	FnameShort string `json:"-"`
}

func (c Caller) String() string {
	bs, _ := json.Marshal(c)
	return string(bs)
}

func GetCaller(skip int) (cs Caller) {
	if skip == 0 {
		skip = 1
	}

	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		cs.File = file
		cs.Line = line
		fc := runtime.FuncForPC(pc)
		cs.Fname = fc.Name()
		drname := filepath.Dir(fc.Name())
		spsd := strings.Split(drname, "/")
		if len(spsd) >= 2 {
			spsd = spsd[len(spsd)-2:]
		}
		spsd = append(spsd, filepath.Base(fc.Name()))
		cs.FnameShort = filepath.Join(spsd...)
	}

	return cs
}

func GetFields(m interface{}, opts []interface{}) (logrus.Fields, []interface{}) {
	fields := logrus.Fields{}
	var interfaces []interface{}
	if m != nil {
		interfaces = append(interfaces, m)
	}

	if opts != nil {
		for _, values := range opts {
			if values == nil {
				continue
			}

			tValues := reflect.ValueOf(values)
			switch tValues.Kind() {
			case reflect.Map:
				for _, k := range tValues.MapKeys() {
					fields[k.String()] = tValues.MapIndex(k).Interface()
				}
			default:
				interfaces = append(interfaces, values)
			}
		}
	}

	return fields, interfaces
}

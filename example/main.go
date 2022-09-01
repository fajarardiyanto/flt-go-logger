package main

import (
	"github.com/fajarardiyanto/flt-go-logger/interfaces"
	"github.com/fajarardiyanto/flt-go-logger/lib"
)

func main() {
	logger := lib.NewLib()
	logger.Init("Testing modules")
	logger.SetOutputFormat(interfaces.OutputFormatJSON)
	loggerOutput(logger)
}

func loggerOutput(logger interfaces.Logger) {
	logger.Info("Lorem Ipsum is simply dummy text of the printing and typesetting %s.", "industry")
	logger.Debug("Lorem Ipsum is simply dummy text of the printing and typesetting %s.", "industry")
	logger.Info(map[string]interface{}{
		"name": "Flt Go Logger",
	})
}

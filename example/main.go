package main

import (
	"gitlab.com/fajardiyanto/flt-go-logger/interfaces"
	"gitlab.com/fajardiyanto/flt-go-logger/lib"
)

func main() {
	logger := lib.NewLib()
	logger.Init("Testing modules")
	loggerOutput(logger)
}

func loggerOutput(logger interfaces.Logger) {
	logger.Info("Lorem Ipsum is simply dummy text of the printing and typesetting %s.", "industry")
	logger.Debug("Lorem Ipsum is simply dummy text of the printing and typesetting %s.", "industry")
	logger.Info(map[string]interface{}{
		"name": "Flt Go Logger",
	})
}

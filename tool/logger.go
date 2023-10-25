package tool

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func NewLogger(filename string) (logger *logrus.Logger) {
	var output io.Writer
	if len(filename) == 0 {
		filename = os.Getenv("LOG_OUTPUT")
	}
	output, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		output = os.Stdout
	}

	logger = logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	logger.SetOutput(output)
	return
}

var Logger = NewLogger("")

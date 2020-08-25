package logger

import (
	"io"

	"github.com/sirupsen/logrus"
)

var baseLogger = logrus.New()
var baseEntry = logrus.NewEntry(baseLogger)

func Entry() *logrus.Entry {
	return baseEntry
}

func SetEntry(e *logrus.Entry) {
	baseEntry = e
}

func SetLevel(l logrus.Level) {
	baseLogger.SetLevel(l)
}

func SetOutput(output io.Writer) {
	baseLogger.SetOutput(output)
}

func SetFormatter(formatter logrus.Formatter) {
	baseLogger.SetFormatter(formatter)
}

package logger

import (
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

func SetFormatter(formatter logrus.Formatter) {
	baseLogger.SetFormatter(formatter)
}

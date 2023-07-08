package logger

import (
	"github.com/sirupsen/logrus"
)

var l logrus.Logger = func() logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)

	return *logger
}()

func Infof(format string, args ...interface{}) {
	l.Infof(format, args...)
}

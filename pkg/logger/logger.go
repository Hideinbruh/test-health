package logger

import "github.com/sirupsen/logrus"

func Logger(err error) {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
}

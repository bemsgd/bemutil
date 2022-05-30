package util

import (
	"os"

	"github.com/bemsgd/bemutil/logging"
)

var logger = logging.GetLogger()

func DieOnError(msg string, err error) {
	if err != nil {
		logger.Add(logging.Error, msg, err.Error())
		os.Exit(1)
	}
}

func LogOnError(msg string, err error) {
	if err != nil {
		logger.Add(logging.Error, msg, err.Error())
	}
}

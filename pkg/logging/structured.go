package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

const (
	FieldKeyMsg         = "msg"
	FieldKeyLevel       = "level"
	FieldKeyTime        = "time"
	FieldKeyLogrusError = "logrus_error"
	FieldKeyFunc        = "func"
	FieldKeyFile        = "file"
)

func New() *logrus.Logger {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(
		&logrus.JSONFormatter{
			TimestampFormat:   "",
			DisableTimestamp:  false,
			DisableHTMLEscape: false,
			DataKey:           "",
			FieldMap: logrus.FieldMap{
				FieldKeyMsg:         FieldKeyMsg,
				FieldKeyLevel:       FieldKeyLevel,
				FieldKeyTime:        FieldKeyTime,
				FieldKeyLogrusError: FieldKeyLogrusError,
				FieldKeyFunc:        FieldKeyFunc,
				FieldKeyFile:        FieldKeyFile,
			},
			CallerPrettyfier: nil,
			PrettyPrint:      false,
		},
	)
	return logger
}

func NewWithFields(fields map[string]interface{}) *logrus.Entry {
	logger := New()
	loggerWithFields := logger.WithFields(fields)
	return loggerWithFields
}

func NewWithID(id string) *logrus.Entry {
	logger := New()
	loggerWithID := logger.WithField("id", id)
	return loggerWithID
}

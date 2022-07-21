package module

import (
	"io"
	"os"
	"path"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

func NewLog(config *Config) (logger *logrus.Logger, err error) {
	level, err := logrus.ParseLevel(config.Log.Level)
	if err != nil {
		return
	}

	path := path.Dir(config.Log.File)
	err = os.MkdirAll(path, 0750)
	if err != nil {
		return
	}

	writers := []io.Writer{}
	file, err := os.OpenFile(config.Log.File, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	if err != nil {
		return
	}
	writers = append(writers, file)

	if config.Log.Stdout {
		writers = append(writers, os.Stderr)
	}

	logrus.SetOutput(io.MultiWriter(writers...))
	logrus.SetLevel(level)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&nested.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		NoColors:        true,
		HideKeys:        true,
		ShowFullLevel:   true,
	})

	logger = logrus.StandardLogger()
	return
}

package logger

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log *logrus.Logger

func Init() {
	Log = logrus.New()

	logFormat := os.Getenv("LOG_FORMAT")
	if logFormat == "json" {
		Log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	} else {
		Log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
			ForceColors:     true,
		})
	}

	logLevel := os.Getenv("LOG_LEVEL")
	switch logLevel {
	case "debug":
		Log.SetLevel(logrus.DebugLevel)
	case "info":
		Log.SetLevel(logrus.InfoLevel)
	case "warn":
		Log.SetLevel(logrus.WarnLevel)
	case "error":
		Log.SetLevel(logrus.ErrorLevel)
	default:
		Log.SetLevel(logrus.InfoLevel)
	}

	enableFileLog := os.Getenv("ENABLE_FILE_LOG")
	Log.Info("File logging enabled:", enableFileLog)
	if enableFileLog == "true" {
		logFile := &lumberjack.Logger{
			Filename:   filepath.Join("logs", "gateway.log"),
			MaxSize:    10,
			MaxBackups: 5,
			MaxAge:     30,
			Compress:   true,
		}

		Log.SetOutput(logFile)
	} else {
		Log.SetOutput(os.Stdout)
	}

	Log.Info("Logger initialized")
}

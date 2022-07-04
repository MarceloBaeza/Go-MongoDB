package util

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

var LOGGER *log.Logger

func ConfigureLog() {
	LOGGER = log.New()
	LOGGER.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		ForceColors:     true,
		TimestampFormat: "2006-01-02 15:04:05.0000",
	})
	LOGGER.SetOutput(os.Stdout)
	getLevelLog()
}

func getLevelLog() {
	leveLConfigurate := os.Getenv("LEVEL_LOG")
	switch strings.ToUpper(leveLConfigurate) {
	case "DEBUG":
		LOGGER.SetLevel(log.DebugLevel)
	case "INFO":
		LOGGER.SetLevel(log.InfoLevel)
	case "WARN":
		LOGGER.SetLevel(log.WarnLevel)
	case "ERROR":
		LOGGER.SetLevel(log.ErrorLevel)
	default:
		LOGGER.SetLevel(log.InfoLevel)
		LOGGER.Warnf("Home: invalid log level supplied: '%s'", leveLConfigurate)
	}
}

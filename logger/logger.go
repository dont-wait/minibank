package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
)

func NewLogger(level zerolog.Level) *zerolog.Logger {
	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stderr,
		NoColor:    false,
		TimeFormat: time.RFC3339,
		FormatLevel: func(i interface{}) string {
			return fmt.Sprintf("[%s]", i)
		},
		FormatCaller: func(i interface{}) string {
			return fmt.Sprintf("[%s]", i)
		},
		FormatMessage: func(i interface{}) string {
			return fmt.Sprintf("{%-20s}", i)
		},
	}
	logger := zerolog.New(consoleWriter).Level(level).With().Caller().Timestamp().Logger()
	return &logger
}

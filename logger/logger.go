package logger

import (
	"fmt"
	"os"
	"strings"
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
			path := fmt.Sprintf("%s", i)
			parts := strings.SplitN(path, "/", -1)
			if len(parts) > 2 {
				path = strings.Join(parts[len(parts)-3:], "/")
			}
			return fmt.Sprintf("[%s]", path)
		},
		FormatMessage: func(i interface{}) string {
			return fmt.Sprintf("{%-20s}", i)
		},
	}
	logger := zerolog.New(consoleWriter).Level(level).With().Caller().Timestamp().Logger()
	return &logger
}

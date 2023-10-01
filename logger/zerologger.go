package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

type ZeroLogger struct {
	zerolog *zerolog.Logger
}

func New() *ZeroLogger {
	zerolog.TimeFieldFormat = zerolog.TimestampFunc().String()
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	logger := zerolog.New(os.Stdout).With().Caller().Logger()
	return &ZeroLogger{zerolog: &logger}
}

func (l *ZeroLogger) Print(message string, properties map[string]any) {
	l.zerolog.Info().CallerSkipFrame(1).Fields(properties).Msg(message)
}

func (l *ZeroLogger) Error(err error, properties map[string]any) {
	l.zerolog.Error().CallerSkipFrame(1).Err(err).Fields(properties).Msg("error")
}

func (l *ZeroLogger) Fatal(err error, properties map[string]any) {
	l.zerolog.Fatal().CallerSkipFrame(1).Fields(properties).Msg(err.Error())
}

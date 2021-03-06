package testhelp

import (
	"strings"
	"testing"

	"github.com/cresta/zapctx"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ZapTestingLogger returns a logger that outputs to the testing Log method, T.Jog
func ZapTestingLogger(t *testing.T) *zapctx.Logger {
	return zapctx.New(zap.New(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
				MessageKey:     "msg",
				LevelKey:       "level",
				NameKey:        "logger",
				EncodeLevel:    zapcore.LowercaseLevelEncoder,
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				EncodeDuration: zapcore.StringDurationEncoder,
			}),
			zapcore.AddSync(&LogSync{t: t}),
			zap.DebugLevel),
	))
}

type LogSync struct {
	t *testing.T
}

func (l *LogSync) Write(p []byte) (n int, err error) {
	l.t.Log(strings.TrimSpace(string(p)))
	return len(p), nil
}

func (l *LogSync) Sync() error {
	return nil
}

var _ zapcore.WriteSyncer = &LogSync{}

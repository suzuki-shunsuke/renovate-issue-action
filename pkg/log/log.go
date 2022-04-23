package log

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewConfig() *zap.Config {
	logCfg := zap.NewProductionConfig()
	logCfg.Encoding = "console"
	logCfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	logCfg.EncoderConfig.ConsoleSeparator = " "
	logCfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return &logCfg
}

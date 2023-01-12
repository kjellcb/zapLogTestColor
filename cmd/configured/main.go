package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.ConsoleSeparator = " | "
	l, _ := config.Build()
	zap.ReplaceGlobals(l)
	defer l.Sync()
}

func main() {
	zap.L().Info("This is a test")
	zap.L().With(
		zap.Bool("bool", false),
		zap.Float64("float64", 100.10),
		zap.Int64("int64", 100),
	).Info("testing fields")
	zap.L().Error("error test")
	zap.L().Warn("warning test")
}

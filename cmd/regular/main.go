package main

import "go.uber.org/zap"

func init() {
	l, _ := zap.NewDevelopment()
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

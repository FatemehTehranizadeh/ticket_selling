package logging

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type NewLogger struct {
	Logger *zap.SugaredLogger
}

var Sugar *NewLogger

func InitLogger() {
	logfile, err := os.OpenFile("internal/logging/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		os.Exit(1)
	}

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"stderr",
			logfile.Name(),
		},
		ErrorOutputPaths: []string{
			"stderr",
			logfile.Name(),
		},
		InitialFields: map[string]interface{}{
			"pid": os.Getpid(),
		},
	}

	logger, err := config.Build()
	if err != nil {
		fmt.Println("Failed to build logger:", err)
		os.Exit(1)
	}

	Sugar = &NewLogger{
		Logger: logger.Sugar(),
	}
}

package logging

import (
	"encoding/json"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger
var sugar *zap.SugaredLogger

func getLogLevel() zapcore.Level {
	defaultLevel := zapcore.DebugLevel

	envLevel := os.Getenv("LOG_LEVEL")

	if envLevel == "" {
		return defaultLevel
	}
	if strings.ToLower(envLevel) == "debug" {
		return zapcore.DebugLevel
	}
	if strings.ToLower(envLevel) == "info" {
		return zapcore.InfoLevel
	}
	if strings.ToLower(envLevel) == "warn" {
		return zapcore.WarnLevel
	}
	if strings.ToLower(envLevel) == "error" {
		return zapcore.ErrorLevel
	}
	if strings.ToLower(envLevel) == "panic" {
		return zapcore.PanicLevel
	}
	return defaultLevel
}

func InitializeLogger() {
	rawJSON := []byte(`{
		"level": "debug",
		"encoding": "json",
		"outputPaths": ["/Users/donovan/code/mcdctl/mcdctl.log"],
		"errorOutputPaths": ["/Users/donovan/code/mcdctl/mcdctl-err.log"],
		"encoderConfig": {
			"messageKey": "message",
			"levelKey": "level",
			"levelEncoder": "lowercase"
		}
	}`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}

	cfg.Level.SetLevel(getLogLevel())

	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.TimeKey = "timestamp"

	logger := zap.Must(cfg.Build())
	defer logger.Sync()

	sugar = logger.Sugar()

	sugar.Info("Logging initialized")
}

func Debug(msg string) {
	if sugar == nil {
		InitializeLogger()
	}
	sugar.Debug(msg)
}

func Info(msg string) {
	if sugar == nil {
		InitializeLogger()
	}
	sugar.Info(msg)
}

func Warn(msg string) {
	if sugar == nil {
		InitializeLogger()
	}
	sugar.Warn(msg)
}

func Error(msg string) {
	if sugar == nil {
		InitializeLogger()
	}
	sugar.Error(msg)
}

func Panic(msg string) {
	if sugar == nil {
		InitializeLogger()
	}
	sugar.Panic(msg)
}

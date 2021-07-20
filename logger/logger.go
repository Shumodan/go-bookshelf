package logger

import (
	_ "embed"
	"fmt"
	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"gopkg.in/yaml.v2"
	"os"
)

// Config represents the setting for zap logger.
type Config struct {
	ZapConfig zap.Config        `json:"zap_config" yaml:"zap_config"`
	LogRotate lumberjack.Logger `json:"log_rotate" yaml:"log_rotate"`
}

// Logger is an alternative implementation of *gorm.Logger
type Logger struct {
	Zap *zap.SugaredLogger
}

// NewLogger create logger object for *gorm.DB from *echo.Logger
func NewLogger(env []byte) *Logger {
	fmt.Printf("read logger configuration: %s")
	configYaml := env
	var myConfig *Config
	if err := yaml.Unmarshal(configYaml, &myConfig); err != nil {
		fmt.Printf("Failed to read zap logger configuration: %s", err)
		os.Exit(2)
	}
	var zap *zap.Logger
	zap, err := build(myConfig)
	if err != nil {
		fmt.Printf("Failed to compose zap logger : %s", err)
		os.Exit(2)
	}
	sugar := zap.Sugar()
	// set package varriable logger.
	logger := &Logger{Zap: sugar}
	logger.Zap.Infof("Success to read zap logger configuration: zaplogger.")
	_ = zap.Sync()
	return logger
}

// GetZapLogger returns zapSugaredLogger
func (log *Logger) GetZapLogger() *zap.SugaredLogger {
	return log.Zap
}

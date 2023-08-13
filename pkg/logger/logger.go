package logger

import (
	"collector-service/config"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger methods interface
type Logger interface {
	InitLogger()
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	DPanic(args ...interface{})
	DPanicf(template string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
}

// Logger
type APIlogger struct {
	cfg         *config.Config
	sugarLogger *zap.SugaredLogger
}
 
// App Logger constructor
func NewAPIlogger(cfg *config.Config) *APIlogger {
	return &APIlogger{cfg: cfg}
}


// For mapping config logger to app logger levels
var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func (l *APIlogger) getLoggerLevel(cfg *config.Config) zapcore.Level {
	level, exist := loggerLevelMap[cfg.Logger.Level]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}

// Init logger
func (l *APIlogger) InitLogger() {
	logLevel := l.getLoggerLevel(l.cfg)

	// Open the log file for writing
	logFile, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		// Handle error opening the log file
		l.sugarLogger.Errorf("Error opening log file: %s", err)
	}

	// Create a core for writing to the log file
	logWriter := zapcore.AddSync(logFile)

	var encoderCfg zapcore.EncoderConfig
	// Check which env mode is in usage
	if l.cfg.Server.Mode == "Development" {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}

	//var encoder zapcore.Encoder
	encoderCfg.LevelKey = "LEVEL"
	encoderCfg.CallerKey = "CALLER"
	encoderCfg.TimeKey = "TIME"
	encoderCfg.NameKey = "NAME"
	encoderCfg.MessageKey = "MESSAGE"

	/* if l.cfg.Logger.Encoding == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	} */

	fileEncoder := zapcore.NewJSONEncoder(encoderCfg)
	consoleEncoder := zapcore.NewConsoleEncoder(encoderCfg)

	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	// Create a core for writing to the log file to file and console
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, logWriter, zap.NewAtomicLevelAt(logLevel)),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zap.NewAtomicLevelAt(logLevel)),
	)

	// core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(logLevel))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	l.sugarLogger = logger.Sugar()
	if err := l.sugarLogger.Sync(); err != nil {
		l.sugarLogger.Error(err)
	}


}

// Logger methods

func (l *APIlogger) Debug(args ...interface{}) {
	l.sugarLogger.Debug(args...)
}

func (l *APIlogger) Debugf(template string, args ...interface{}) {
	l.sugarLogger.Debugf(template, args...)
}

func (l *APIlogger) Info(args ...interface{}) {
	l.sugarLogger.Info(args...)
}

func (l *APIlogger) Infof(template string, args ...interface{}) {
	l.sugarLogger.Infof(template, args...)
}

func (l *APIlogger) Warn(args ...interface{}) {
	l.sugarLogger.Warn(args...)
}

func (l *APIlogger) Warnf(template string, args ...interface{}) {
	l.sugarLogger.Warnf(template, args...)
}

func (l *APIlogger) Error(args ...interface{}) {
	l.sugarLogger.Error(args...)
}

func (l *APIlogger) Errorf(template string, args ...interface{}) {
	l.sugarLogger.Errorf(template, args...)
}

func (l *APIlogger) DPanic(args ...interface{}) {
	l.sugarLogger.DPanic(args...)
}

func (l *APIlogger) DPanicf(template string, args ...interface{}) {
	l.sugarLogger.DPanicf(template, args...)
}

func (l *APIlogger) Panic(args ...interface{}) {
	l.sugarLogger.Panic(args...)
}

func (l *APIlogger) Panicf(template string, args ...interface{}) {
	l.sugarLogger.Panicf(template, args...)
}

func (l *APIlogger) Fatal(args ...interface{}) {
	l.sugarLogger.Fatal(args...)
}

func (l *APIlogger) Fatalf(template string, args ...interface{}) {
	l.sugarLogger.Fatalf(template, args...)
}

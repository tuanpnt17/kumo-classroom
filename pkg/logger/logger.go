package logger

import (
	"github.com/tuanpnt17/kumo-classroom-be/pkg/settings"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type ZapLogger struct {
	*zap.Logger
}

func NewZapLogger(loggerSetting settings.ZapLoggerSetting) *ZapLogger {

	level := getLogLevel(loggerSetting.Level)
	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename:   loggerSetting.LogPath,
		MaxSize:    loggerSetting.MaxSize, // megabytes
		MaxBackups: loggerSetting.MaxBackups,
		MaxAge:     loggerSetting.MaxAge, //days
		LocalTime:  loggerSetting.LocalTime,
		Compress:   loggerSetting.Compress, // disabled by default
	}
	syncer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook), zapcore.AddSync(zapcore.Lock(os.Stdout)))
	core := zapcore.NewCore(encoder, syncer, level)

	customLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	return &ZapLogger{
		Logger: customLogger,
	}
}

func getLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
func getEncoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	// config log time
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"

	// config log level
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// config caller
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewConsoleEncoder(encoderConfig)
}

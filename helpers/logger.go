package helpers

import (
	"fmt"
	"log"
	"log/syslog"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewSyslogCore creates a zap core that writes logs to rsyslog
func NewSyslogCore() zapcore.Core {
    // Connect to local rsyslog using /dev/log
    writer, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL0, "api-gateway")
    if err != nil {
        log.Fatalf("failed to connect to syslog: %v", err)
    }

    syslogWriter := zapcore.AddSync(writer)

    encoderConfig := zap.NewProductionEncoderConfig()
    encoderConfig.TimeKey = ""
    encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
    encoderConfig.LevelKey = "level"
    encoderConfig.NameKey = "logger"
    encoderConfig.CallerKey = "caller"
    encoderConfig.MessageKey = "msg"
    encoderConfig.StacktraceKey = "stacktrace"
    encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
    encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

    core := zapcore.NewCore(
        zapcore.NewConsoleEncoder(encoderConfig), // Use console encoder
        syslogWriter,
        zapcore.InfoLevel,
    )

    return core
}

func Logger() *zap.Logger {
    logger := zap.New(NewSyslogCore(), zap.AddCaller())
    defer logger.Sync()
    return logger
}

func LoggerMess(message string, method string, status int) string {
    return fmt.Sprintf("%s %s %d", message, method, status)
}
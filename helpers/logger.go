package helpers

import (
	"log"
	"log/syslog"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewSyslogCore creates a zap core that writes logs to rsyslog
func NewSyslogCore() zapcore.Core {
	filename := os.Getenv("LOG_FILENAME")
	
	// Connect to local rsyslog using /dev/log
	writer, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL0, filename)
	if err != nil {
		log.Fatalf("failed to connect to syslog: %v", err)
	}

	syslogWriter := zapcore.AddSync(writer)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // Use JSON encoder
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

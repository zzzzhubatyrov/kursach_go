package pkg

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"gorm.io/gorm/logger"
)

type CustomLogger struct {
	Logger logger.Interface
	Level  logger.LogLevel
	File   *os.File
}

func (l *CustomLogger) Info(ctx context.Context, str string, data ...interface{}) {
	if l.Level >= logger.Info && strings.Contains(str, "CREATE TABLE") {
		l.Logger.Info(ctx, str, data...)
		l.logToFile("INFO", str)
	}
}

func (l *CustomLogger) Warn(ctx context.Context, str string, data ...interface{}) {
	if l.Level >= logger.Warn {
		l.Logger.Warn(ctx, str, data...)
		l.logToFile("WARN", str)
	}
}

func (l *CustomLogger) Error(ctx context.Context, str string, data ...interface{}) {
	if l.Level >= logger.Error {
		l.Logger.Error(ctx, str, data...)
		l.logToFile("ERROR", str)
	}
}

func (l *CustomLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.Level >= logger.Silent {
		sql, rows := fc()
		if strings.Contains(sql, "CREATE TABLE") {
			l.logToFile("TRACE", fmt.Sprintf("SQL: %s | Rows affected: %d", sql, rows))
		}
	}
}

func (l *CustomLogger) LogMode(Level logger.LogLevel) logger.Interface {
	l.Level = Level
	return l
}

func (l *CustomLogger) logToFile(Level, msg string) {
	logEntry := fmt.Sprintf("[%s] %s\n", Level, msg)
	l.File.WriteString(logEntry)
}

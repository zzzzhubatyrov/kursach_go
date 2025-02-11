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
	logger logger.Interface
	level  logger.LogLevel
	file   *os.File
}

func (l *CustomLogger) Info(ctx context.Context, str string, data ...interface{}) {
	if l.level >= logger.Info && strings.Contains(str, "CREATE TABLE") {
		l.logger.Info(ctx, str, data...)
		l.logToFile("INFO", str)
	}
}

func (l *CustomLogger) Warn(ctx context.Context, str string, data ...interface{}) {
	if l.level >= logger.Warn {
		l.logger.Warn(ctx, str, data...)
		l.logToFile("WARN", str)
	}
}

func (l *CustomLogger) Error(ctx context.Context, str string, data ...interface{}) {
	if l.level >= logger.Error {
		l.logger.Error(ctx, str, data...)
		l.logToFile("ERROR", str)
	}
}

func (l *CustomLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.level >= logger.Silent {
		sql, rows := fc()
		if strings.Contains(sql, "CREATE TABLE") {
			l.logToFile("TRACE", fmt.Sprintf("SQL: %s | Rows affected: %d", sql, rows))
		}
	}
}

func (l *CustomLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.level = level
	return l
}

func (l *CustomLogger) logToFile(level, msg string) {
	logEntry := fmt.Sprintf("[%s] %s\n", level, msg)
	l.file.WriteString(logEntry)
}

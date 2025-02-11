package storage

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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

func SQLiteStorageInit() (*gorm.DB, error) {
	file, err := os.OpenFile("gorm.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %v", err)
	}

	newLogger := logger.New(
		log.New(file, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      false,
		},
	)

	customLogger := &CustomLogger{logger: newLogger, level: logger.Info, file: file}

	db, err := gorm.Open(sqlite.Open("kurs.db"), &gorm.Config{
		Logger: customLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("err: %v", err)
	}
	return db, err
}

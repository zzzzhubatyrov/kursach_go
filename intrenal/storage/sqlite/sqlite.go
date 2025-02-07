package storage

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// CustomLogger - пользовательский логгер для фильтрации CREATE TABLE запросов
type CustomLogger struct {
	logger logger.Interface
	level  logger.LogLevel
}

func (l *CustomLogger) Info(ctx context.Context, str string, data ...interface{}) {
	if l.level >= logger.Info && strings.Contains(str, "CREATE TABLE") {
		l.logger.Info(ctx, str, data...)
	}
}

func (l *CustomLogger) Warn(ctx context.Context, str string, data ...interface{}) {
	if l.level >= logger.Warn {
		l.logger.Warn(ctx, str, data...)
	}
}

func (l *CustomLogger) Error(ctx context.Context, str string, data ...interface{}) {
	if l.level >= logger.Error {
		l.logger.Error(ctx, str, data...)
	}
}

func (l *CustomLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.level >= logger.Silent {
		l.logger.Trace(ctx, begin, fc, err)
	}
}

func (l *CustomLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.level = level
	return l
}

func SQLiteStorageInit() (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	customLogger := &CustomLogger{logger: newLogger}

	db, err := gorm.Open(sqlite.Open("kurs.db"), &gorm.Config{
		Logger: customLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("err: %v", err)
	}
	return db, err
}

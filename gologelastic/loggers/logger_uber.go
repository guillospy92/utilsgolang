package loggers

import (
	"errors"
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"math/rand"
	"os"
	"time"
)

// where is all information
// https://programmer.help/blogs/zap-log-base-practice.html practice full good
// https://pmihaylov.com/go-service-with-elk/ configure with container golang

type User struct {
	Name     string
	LastName string
	Age      int
}

func (f *User) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("name", f.Name)
	enc.AddString("last_name", f.LastName)
	enc.AddInt("age", f.Age)
	return nil
}

func AddLoggerLogStash() {
	reasons := []string{
		"good",
		"bad",
		"error_logic",
		"not_found",
	}
	a := User{
		Name:     "Guillermo Rom",
		LastName: "Rom Rom",
		Age:      25,
	}

	encoderConfig := ecszap.NewDefaultEncoderConfig()
	// more information https://programmer.help/blogs/zap-log-base-practice.html practice full good
	// more information second https://github.com/uber-go/zap/blob/master/example_test.go
	// multiples output

	getLogWriter := func() zapcore.WriteSyncer {
		lumberJackLogger := &lumberjack.Logger{
			Filename:   "./logs/logstash/logs.log",
			MaxSize:    1,
			MaxBackups: 3,
			MaxAge:     1,
			Compress:   false,
		}
		return zapcore.AddSync(lumberJackLogger)
	}

	core := ecszap.NewCore(encoderConfig, zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), getLogWriter()), zap.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	logger = logger.
		With(zap.String("app", "goMyapp")).
		With(zap.String("environment", "psm")).
		With(zap.String("coders", "coders"))
	count := 0
	for {
		time.Sleep(1 * time.Second)
		var status int
		n := rand.Int() % len(reasons)
		if reasons[n] == "good" {
			status = 200
		}
		if reasons[n] == "bad" {
			status = 400
		}
		if reasons[n] == "error_logic" {
			status = 280
		}
		if reasons[n] == "not_found" {
			status = 404
		}

		if rand.Float32() > 0.8 {
			logger.Error("oops...something is wrong",
				zap.Int("count", count),
				zap.Int("status", status),
				zap.Error(errors.New("error details")),
			)
			logger.With(zap.Object("ctx", &a))
		} else {
			logger.Info(
				"everything is fine",
				zap.Int("count", count),
				zap.Int("status", status),
				zap.Object("ctx", &a),
			)
			logger.With(zap.Object("ctx", &a))
		}
		count++
		//time.Sleep(time.Millisecond * 10)
	}
}

func NewLoggerInstanceWitOutLogStash() {
	getLogWriter := func() zapcore.WriteSyncer {
		lumberJackLogger := &lumberjack.Logger{
			Filename:   "./logs/normal/logs.log",
			MaxSize:    1,
			MaxBackups: 3,
			MaxAge:     1,
			Compress:   false,
		}
		return zapcore.AddSync(lumberJackLogger)
	}
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.DebugLevel)

	encoder := func() zapcore.Encoder {
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.TimeKey = "timestamp"
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		return zapcore.NewJSONEncoder(encoderConfig)
	}

	core := zapcore.NewCore(
		encoder(),
		zapcore.NewMultiWriteSyncer(getLogWriter(), zapcore.AddSync(os.Stdout)),
		atomicLevel,
	)
	field := zap.Fields(
		zap.String("appName", "app_golang"),
		zap.String("name", "Guillermo"),
	)

	development := zap.Development()
	logger := zap.New(core, zap.AddCaller(), development, field).Sugar()

	logger.Info("everything is fine", zap.Int("count", 1))
	var count = 0
	for {
		if rand.Float32() > 0.8 {
			logger.Error(
				"oops...something is wrong",
				zap.Int("count", count),
				zap.Int("code", 404),
				zap.Error(errors.New("error details")),
			)
		} else {
			logger.Info(
				"everything is fine",
				zap.Int("count", count),
				zap.Int("code", 280),
			)
		}
		time.Sleep(time.Second * 1)
		count++
	}
}

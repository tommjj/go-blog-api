package logger

import (
	"os"

	"github.com/tommjj/go-blog-api/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logLevelMap = map[string]zapcore.Level{
	"Debug":  zap.DebugLevel,
	"Info":   zap.InfoLevel,
	"Warn":   zap.WarnLevel,
	"Error":  zap.ErrorLevel,
	"DPanic": zap.DPanicLevel,
	"Panic":  zap.PanicLevel,
	"Fatal":  zap.FatalLevel,
}

var L *zap.Logger

func Set(conf config.Logger) error {
	level, ok := logLevelMap[conf.Level]
	if !ok {
		return nil
	}

	var encoder zapcore.EncoderConfig
	if conf.Encoder == "production" {
		encoder = zap.NewProductionEncoderConfig()
	} else {
		encoder = zap.NewDevelopmentEncoderConfig()
	}

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   conf.FileName,
		MaxSize:    conf.MaxSize, // megabytes
		MaxBackups: conf.MaxBackups,
		MaxAge:     conf.MaxAge, // days
	})

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoder),
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(w),
			zapcore.AddSync(os.Stderr),
		),
		level,
	)

	L = zap.New(core)

	return nil
}

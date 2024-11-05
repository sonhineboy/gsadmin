package initialize

import (
	"fmt"
	"github.com/sonhineboy/gsadmin/service/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

func ZapInit(c *config.Config) *zap.SugaredLogger {
	core := zapcore.NewCore(enc(), ws(c), zap.NewAtomicLevel())
	logger := zap.New(core, zap.AddStacktrace(zap.ErrorLevel))
	return logger.Sugar()
}

func ZapSync(zap *zap.SugaredLogger) {
	err := zap.Sync()
	if err != nil {
		return
	}
}

func enc() zapcore.Encoder {
	cfg := zap.NewProductionEncoderConfig()
	cfg.TimeKey = "time"
	cfg.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	return zapcore.NewJSONEncoder(cfg)
}

func ws(c *config.Config) zapcore.WriteSyncer {
	logFileName := fmt.Sprintf("./%s/web-%v.log", c.Logger.Path, time.Now().Format("2006-01-02"))
	lumberjackLogger := &lumberjack.Logger{
		Filename: logFileName,
		MaxSize:  c.Logger.Size,
		//MaxBackups: c.Logger.MaxAge,
		MaxAge:   c.Logger.MaxAge,
		Compress: false,
	}

	if c.Logger.StdOut {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberjackLogger), zapcore.AddSync(os.Stdout))
	} else {
		return zapcore.AddSync(lumberjackLogger)
	}
}

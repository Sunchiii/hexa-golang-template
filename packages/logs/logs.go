package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger


func init() {
  config := zap.NewProductionConfig()
  config.EncoderConfig.TimeKey = "timestamp"
  config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
  config.EncoderConfig.StacktraceKey = ""
  
  var err error
  log, err = config.Build(zap.AddCallerSkip(1))
  if err != nil {
    panic(err)
  }
}

func Info(msg string, fields ...zap.Field){
  log.Info(msg, fields...)
}
func Debug(msg string, fields ...zap.Field){
  log.Debug(msg, fields...)
}

func Error(message interface{}, fields ...zap.Field){
  switch message.(type) {
    case string:
      log.Error(message.(string), fields...)
    case error:
      log.Error(message.(error).Error(), fields...)
  }
}


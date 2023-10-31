package log

import (
	"bill/modules/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var (
	logger      *zap.Logger
	sugarLogger *zap.SugaredLogger
)

func Get() *zap.Logger {
	return logger
}

func GetSugar() *zap.SugaredLogger {
	return sugarLogger
}

func init() {
	logger = initComplexLogger()
	sugarLogger = logger.Sugar()
}

func initComplexLogger() *zap.Logger {

	dir := setting.GetProjectPath() + "/logs"

	normal := &lumberjack.Logger{
		Filename:   dir + "/normal.log",
		MaxSize:    5,
		MaxAge:     30,
		MaxBackups: 10,
		Compress:   false,
	}
	critical := &lumberjack.Logger{
		Filename:   dir + "/critical.log",
		MaxSize:    5,
		MaxAge:     30,
		MaxBackups: 10,
		Compress:   false,
	}
	// 分割error以及以上信息到critical.log，其余在normal.log

	// First, define our level-handling logic.
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	normalSync := zapcore.AddSync(normal)
	criticalSyncer := zapcore.AddSync(critical)
	consoleSyncer := zapcore.Lock(os.Stderr)

	consoleEncoder := getEncoderForConsole()
	fileEncoder := getEncoderForFile()

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleSyncer, zapcore.InfoLevel),
		zapcore.NewCore(fileEncoder, normalSync, lowPriority),
		zapcore.NewCore(fileEncoder, criticalSyncer, highPriority),
	)
	return zap.New(core).WithOptions(zap.AddCaller(), zap.AddStacktrace(highPriority))
}

func getEncoderForFile() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)
	return jsonEncoder
}

func getEncoderForConsole() zapcore.Encoder {
	consoloEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	return consoloEncoder
}

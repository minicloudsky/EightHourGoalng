package zaplog

import (
	"encoding/json"
	"fmt"
	"github.com/minicloudsky/golang-in-action/util/file"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// ZapLogger get zap logger
func ZapLogger(logLevel, logPath string) (*zap.Logger, error) {
	if !file.IsExist(logPath) {
		fmt.Println("logPath: ", logPath)
		file, err := os.Create(logPath)
		if err != nil {
			panic(err)
		}
		fmt.Println("log path: ", file.Name())
	}
	rawJSON := []byte(fmt.Sprintf(`{
	  "level": "%s",
	  "encoding": "json",
	  "outputPaths": ["stdout", "%s"],
	  "errorOutputPaths": ["stderr"],
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  },
      "EncodeTime":"zapcore.ISO8601TimeEncoder"
	}`, logLevel, logPath))

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	return logger, err
}

// InitLogger logpath 日志文件路径
// loglevel 日志级别
func InitLogger(logpath, loglevel, serviceName string) *zap.Logger {
	// 日志分割
	hook := lumberjack.Logger{
		Filename:   logpath, // 日志文件路径，默认 os.TempDir()
		MaxSize:    10,      // 每个日志文件保存10M，默认 100M
		MaxBackups: 30,      // 保留30个备份，默认不限
		MaxAge:     7,       // 保留7天，默认不限
		Compress:   true,    // 是否压缩，默认不压缩
	}
	write := zapcore.AddSync(&hook)
	// 设置日志级别
	// debug 可以打印出 info debug warn
	// info  级别可以打印 warn info
	// warn  只能打印 warn
	// debug->info->warn->error
	var level zapcore.Level
	switch loglevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseColorLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,         // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,     //
		EncodeCaller:   zapcore.ShortCallerEncoder,         // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)
	core := zapcore.NewCore(
		// zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewJSONEncoder(encoderConfig),
		// zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&write)), // 打印到控制台和文件
		write,
		level,
	)
	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段,如：添加一个服务器名称
	filed := zap.Fields(zap.String("serviceName", serviceName))
	// 构造日志
	logger := zap.New(core, caller, development, filed)
	//logger.Info("DefaultLogger init success")
	return logger
}

// InitZapLogger 支持error、warn、info多种日志级别
func InitZapLogger() *zap.Logger {
	var coreArr []zapcore.Core
	//获取编码器
	encoderConfig := zap.NewProductionEncoderConfig()            //NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder        //指定时间格式
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder //按级别显示不同颜色，不需要的话取值zapcore.CapitalLevelEncoder就可以了
	//encoderConfig.EncodeCaller = zapcore.FullCallerEncoder      	//显示完整文件路径
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	//日志级别
	highPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool { //error级别
		return level >= zap.ErrorLevel
	})
	mediumPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { //warn级别
		return lev < zap.ErrorLevel && lev >= zap.WarnLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { //info和debug级别,debug级别是最低的
		return lev < zap.ErrorLevel && lev >= zap.DebugLevel
	})

	//info文件writeSyncer
	infoFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./log/info.log", //日志文件存放目录，如果文件夹不存在会自动创建
		MaxSize:    2,                //文件大小限制,单位MB
		MaxBackups: 100,              //最大保留日志文件数量
		MaxAge:     30,               //日志文件保留天数
		Compress:   false,            //是否压缩处理
	})
	infoFileCore := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(infoFileWriteSyncer, zapcore.AddSync(os.Stdout)),
		lowPriority) //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
	//warn文件writeSyncer
	warnFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./log/warn.log", //日志文件存放目录，如果文件夹不存在会自动创建
		MaxSize:    2,                //文件大小限制,单位MB
		MaxBackups: 100,              //最大保留日志文件数量
		MaxAge:     30,               //日志文件保留天数
		Compress:   false,            //是否压缩处理
	})
	warnFileCore := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(warnFileWriteSyncer, zapcore.AddSync(os.Stdout)),
		mediumPriority) //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志

	//error文件writeSyncer
	errorFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./log/error.log", //日志文件存放目录
		MaxSize:    1,                 //文件大小限制,单位MB
		MaxBackups: 5,                 //最大保留日志文件数量
		MaxAge:     30,                //日志文件保留天数
		Compress:   false,             //是否压缩处理
	})
	errorFileCore := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(errorFileWriteSyncer, zapcore.AddSync(os.Stdout)),
		highPriority) //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志

	coreArr = append(coreArr, infoFileCore)
	coreArr = append(coreArr, warnFileCore)
	coreArr = append(coreArr, errorFileCore)
	logger := zap.New(zapcore.NewTee(coreArr...), zap.AddCaller()) //zap.AddCaller()为显示文件名和行号，可省略
	return logger
}
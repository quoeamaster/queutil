package queutil

import (
    "fmt"
)

// TODO: add ILogFormatter interface for formatting logs (if necessary)






type FlexLoggerConfig struct {
    LogFile string
    LogFileMaxSizeMb int
    LogFileMaxBackups int
    LogFileMaxDaysForRetention int
    LogFileBackupCompress bool
}

func NewFlexLoggerConfig(file string, maxFileSize, maxFileBackups, maxRetentionDays int, compressLogFile bool) *FlexLoggerConfig {
    c := new(FlexLoggerConfig)

    c.LogFile = file
    if maxFileSize > 0 {
        c.LogFileMaxSizeMb = maxFileSize
    }
    if maxFileBackups > 0 {
        c.LogFileMaxBackups = maxFileBackups
    }
    if maxRetentionDays > 0 {
        c.LogFileMaxDaysForRetention = maxRetentionDays
    }
    c.LogFileBackupCompress = compressLogFile

    return c
}


type FlexLogger struct {
    // map of ILogger instances for logging. Key is the name of the logger,
    // value is the ILogger implementation
    Loggers map[string]ILogger



    // lumberjack logger is for roll-able file logging
    //lumberjackLogger *lumberjack.Logger

    // console / stdout logger
    //consoleLogger *ConsoleLogger
}

func NewFlexLogger() *FlexLogger {
    l := new(FlexLogger)

    l.Loggers = make(map[string]ILogger)

    // TODO: no more initialization here... instead call AddLogger(*ILogger) instead
    /*
    if loggerConfig != nil {
        l.lumberjackLogger = &lumberjack.Logger{
            Filename: loggerConfig.LogFile,
            MaxSize: loggerConfig.LogFileMaxSizeMb,
            MaxBackups: loggerConfig.LogFileMaxBackups,
            MaxAge: loggerConfig.LogFileMaxDaysForRetention,
            Compress: loggerConfig.LogFileBackupCompress,
        }
    }
    l.consoleLogger = new(ConsoleLogger)
    */
    return l
}

// method add the given logger implementation
func (f *FlexLogger) AddLogger(logger ILogger) {
    if logger != nil {
        f.Loggers[logger.Name()] = logger
    }
}

// similar to calling WriteWithOptions([]byte, nil);
// which means all available Logger(s) would log the given message
func (f *FlexLogger) Write(p []byte) (n int, err error) {
    return f.WriteWithOptions(p, nil)
}

// log message based on the options provided; if options is nil then all
// logger(s) will log the given []byte, else would need to
// check if a "true" is associated with logger's name in which a "true"
// indicates the logger to log the message
func (f *FlexLogger) WriteWithOptions(p []byte, options map[string]bool) (n int, err error) {
    // force log for all available logger(s)
    forceLog := false

    if options == nil {
        forceLog = true
    }

    for loggerName, logger := range f.Loggers {
        if forceLog == true || options[loggerName] == true {
            _, err := logger.Write(p)
            if err != nil {
                // TODO: should let it pass through or return error?
                return 0, err
            }
        }
    }
    return 0, nil
}


// release resources before the instance is removed
func (f *FlexLogger) Release(optionalParam map[string]interface{}) error {
    var firstErr error

    for _, logger := range f.Loggers {
        err := logger.Release(optionalParam)
        if err != nil {
            if firstErr == nil {
                firstErr = err
            }
            fmt.Printf("[%v] release error => %v\n", logger.Name(), err)
        }
    }
    return firstErr
}


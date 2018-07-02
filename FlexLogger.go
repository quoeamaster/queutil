package queutil

import (
    "gopkg.in/natefinch/lumberjack.v2"
    "fmt"
)

// TODO: add ILogFormatter interface for formatting logs (if necessary)
// TODO: add switches to control which underlying logger should log the message (e.g. prevent consoleLogger to log)

// Logger for console / stdout logging
type ConsoleLogger struct {

}

// simple just write to the stdout
func (c *ConsoleLogger) Write(p []byte) (n int, err error) {
    return fmt.Println(string(p))
}


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
    // lumberjack logger is for roll-able file logging
    lumberjackLogger *lumberjack.Logger

    // console / stdout logger
    consoleLogger *ConsoleLogger
}

func NewFlexLogger(loggerConfig *FlexLoggerConfig) *FlexLogger {
    l := new(FlexLogger)

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

    return l
}

func (f *FlexLogger) Write(p []byte) (n int, err error) {
    // write to console (usually ok)
    f.consoleLogger.Write(p)

    // write to file (lumberjack)
    iWrote, err := f.lumberjackLogger.Write(p)
    if err != nil {
        return iWrote, err
    }
    _, err = f.lumberjackLogger.Write([]byte("\n"))
    if err != nil {
        return 1, err
    }
    return iWrote, nil
}

// release resources before the instance is removed
func (f *FlexLogger) Release(optionalParam map[string]interface{}) error {
    // consoleLogger... anything to release???

    return f.lumberjackLogger.Close()
}


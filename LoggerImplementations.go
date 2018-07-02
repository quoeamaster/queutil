package queutil

import (
    "gopkg.in/natefinch/lumberjack.v2"
    "fmt"
)

// -------------------------
// *** RollingFileLogger ***
// -------------------------

type RollingFileLogger struct {
    // the actual logger is lumberjack...
    logger *lumberjack.Logger
}
// method to create a lumberjackLogger instance. LumberjackLogger is an
// implementation for rolling file logger
func NewRollingFileLogger(filename string, maxFileSize, maxBackup, maxRetentionDays int, compress bool) *RollingFileLogger {
    l := new(RollingFileLogger)
    l.logger = &lumberjack.Logger {
        Filename: filename,
        MaxSize: maxFileSize,
        MaxBackups: maxBackup,
        MaxAge: maxRetentionDays,
        Compress: compress,
    }
    return l
}

func (l *RollingFileLogger) Write(p []byte) (n int, err error) {
    return l.logger.Write(p)
}
func (l *RollingFileLogger) Name() string {
    return "rollingFileLogger"
}
func (l *RollingFileLogger) Release(optionalParam map[string]interface{}) error {
    if l.logger != nil {
        return l.logger.Close()
    }
    return nil
}


// ---------------------
// *** ConsoleLogger ***
// ---------------------

// Logger for console / stdout logging
type ConsoleLogger struct {

}
// method to create a console logger
func NewConsoleLogger() *ConsoleLogger {
    l := new(ConsoleLogger)
    return l
}

// simple just write to the stdout
func (l *ConsoleLogger) Write(p []byte) (n int, err error) {
    return fmt.Print(string(p))
}
func (l *ConsoleLogger) Name() string {
    return "consoleLogger"
}
func (l *ConsoleLogger) Release(optionalParam map[string]interface{}) error {
    // release resource (if any)
    return nil
}
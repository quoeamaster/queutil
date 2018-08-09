package queutil

import (
    "fmt"
    "bytes"
    "time"
)

const LogLevelDebug = 1
const LogLevelInfo  = 2
const LogLevelWarn  = 4
const LogLevelErr   = 8

type FlexLogger struct {
    // map of ILogger instances for logging. Key is the name of the logger,
    // value is the ILogger implementation
    Loggers map[string]ILogger

    // the log level for the logger(s)
    LogLevel int
}

func NewFlexLogger() *FlexLogger {
    l := new(FlexLogger)

    // no more initialization here... instead call AddLogger(*ILogger) instead
    l.Loggers = make(map[string]ILogger)

    // log level default is info
    l.LogLevel = LogLevelInfo

    return l
}

// method add the given logger implementation
func (f *FlexLogger) AddLogger(logger ILogger) *FlexLogger {
    if logger != nil {
        f.Loggers[logger.Name()] = logger
    }
    return f
}

// return the configured ILogger(s)'s name
func (f *FlexLogger) GetLoggerNames() []string {
    lNames := make([]string, 0)
    for name := range f.Loggers {
        lNames = append(lNames, name)
    }
    return lNames
}

// similar to calling WriteWithOptions([]byte, nil);
// which means all available Logger(s) would log the given message;
// plus using the default logLevel (which is info if not set)
func (f *FlexLogger) Write(p []byte) (n int, err error) {
    return f.WriteWithOptions(p, nil, f.LogLevel)
}

// log message based on the options provided; if options is nil then all
// logger(s) will log the given []byte, else would need to
// check if a "true" is associated with logger's name in which a "true"
// indicates the logger to log the message;
func (f *FlexLogger) WriteWithOptions(p []byte, options map[string]bool, logLevel int) (n int, err error) {
    // check if the given logLevel should be logged or not
    if f.LogLevel > logLevel {
        return 0, nil
    }
    // force log for all available logger(s)
    forceLog := false

    if options == nil {
        forceLog = true
    }

    // add timestamp and log-level headers for the message
    finalBytes := f.addLogHeader(p, logLevel, nil)

    for loggerName, logger := range f.Loggers {
        if forceLog == true || options[loggerName] == true {
            _, err := logger.Write(finalBytes)
            if err != nil {
                // TODO: should let it pass through or return error?
                return 0, err
            }
        }
    }
    return 0, nil
}

// a method to add the LogHeader (private method)
func (f *FlexLogger) addLogHeader(p []byte, logLevel int, optionals map[string]interface{}) []byte {
    var buf bytes.Buffer
    var logLevelString = "info"

    now := time.Now().UTC()
    // 2006-01-02T15:04:05.999999-07:00 => hk +08:00 but since we are using UTC... must be 00:00:00
    buf.WriteString(fmt.Sprintf("[%v]", now.Format("2006-01-02T15:04:05.999999 UTC")))
    switch logLevel {
    case LogLevelDebug:
        logLevelString = "DEBUG"
    case LogLevelInfo:
        logLevelString = "INFO "
    case LogLevelWarn:
        logLevelString = "WARN "
    case LogLevelErr:
        logLevelString = "ERR  "
    default:
        logLevelString = "INFO "
    }
    buf.WriteString(fmt.Sprintf("[%v] ", logLevelString))
    buf.Write(p)

    // additional params
    if optionals != nil {
        for key, value := range optionals {
            buf.WriteString(fmt.Sprintf("[%v -> %v]", key, value))
        }
    }
    // final linebreak
    // buf.WriteString("\n")

    return buf.Bytes()
}

func (f *FlexLogger) Debug(p []byte) (int, error) {
    return f.WriteWithOptions(p, nil, LogLevelDebug)
}
func (f *FlexLogger) DebugString(p string) (int, error) {
    return f.WriteWithOptions([]byte(p), nil, LogLevelDebug)
}
func (f *FlexLogger) DebugWithOptions(p []byte, options map[string]bool) (int, error) {
    return f.WriteWithOptions(p, options, LogLevelDebug)
}

func (f *FlexLogger) Info(p []byte) (int, error) {
    return f.WriteWithOptions(p, nil, LogLevelInfo)
}
func (f *FlexLogger) InfoString(p string) (int, error) {
    return f.WriteWithOptions([]byte(p), nil, LogLevelInfo)
}
func (f *FlexLogger) InfoWithOptions(p []byte, options map[string]bool) (int, error) {
    return f.WriteWithOptions(p, options, LogLevelInfo)
}

func (f *FlexLogger) Warn(p []byte) (int, error) {
    return f.WriteWithOptions(p, nil, LogLevelWarn)
}
func (f *FlexLogger) WarnString(p string) (int, error) {
    return f.WriteWithOptions([]byte(p), nil, LogLevelWarn)
}
func (f *FlexLogger) WarnWithOptions(p []byte, options map[string]bool) (int, error) {
    return f.WriteWithOptions(p, options, LogLevelWarn)
}

func (f *FlexLogger) Err(p []byte) (int, error) {
    return f.WriteWithOptions(p, nil, LogLevelErr)
}
func (f *FlexLogger) ErrString(p string) (int, error) {
    return f.WriteWithOptions([]byte(p), nil, LogLevelErr)
}
func (f *FlexLogger) ErrWithOptions(p []byte, options map[string]bool) (int, error) {
    return f.WriteWithOptions(p, options, LogLevelErr)
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


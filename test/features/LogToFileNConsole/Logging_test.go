package LogToFileNConsole

import (
    "github.com/DATA-DOG/godog"
    "queutil"
    "os"
    "fmt"
    "strings"
    "github.com/DATA-DOG/godog/gherkin"
    "path/filepath"
    "strconv"
)

var logger *queutil.FlexLogger
var currentWd = ""
var logFolder = ""

// scenario 1

// step 1
func gotLogFolder(folder string) error {
    logFolder = folder

    return nil
}

// step 2
func gotLogPattern(logFilePattern string) error {
    logFilepath := currentWd + queutil.GetFilepathSeparator() + logFolder + queutil.GetFilepathSeparator() + logFilePattern

    logger = queutil.NewFlexLogger()
    // added 2 loggers
    logger.AddLogger(queutil.NewRollingFileLogger(
        logFilepath, 1, 2, 2, true))
    logger.AddLogger(queutil.NewConsoleLogger())

    return nil
}

// step 3
func log(message string) error {
    iWrote, err := logger.Write([]byte(message + "\n"))
    if err != nil {
        return err
    }
    fmt.Printf("success in logging, wrote '%v' bytes of data\n", iWrote)

    return nil
}

// step 4
func verifyLogResults(logFilePattern, message string) error {
    logFilepath := currentWd + queutil.GetFilepathSeparator() + logFolder + queutil.GetFilepathSeparator() + logFilePattern
    // tail the last entry of the log file
    bytesArr, err := queutil.ReadFileContent(logFilepath)
    if err != nil {
        return nil
    }
    lines := strings.Split(strings.TrimSpace(string(bytesArr)), "\n")
    if lines != nil {
        lastLine := ""

        if len(lines) == 1 {
            lastLine = lines[0]
        } else if len(lines) > 1 {
            lastLine = lines[len(lines)-1]
        } else {
            return fmt.Errorf("no logs available~~~~")
        }

        if strings.Index(lastLine, message) >= 0 {
            return nil
        } else {
            return fmt.Errorf("mismatch~ expected [%v] but got [%v]", message, lastLine)
        }
    }
    defer func() {
        err := logger.Release(nil)
        if err != nil {
            // can panic or just log
            fmt.Println("***", err)
        }
    }()

    return fmt.Errorf("no lines available")
}

// ** scenario 2

var finalLogFolder = ""
var rollingLogger *queutil.FlexLogger

func gotRotationLogFile(rollingFile, logFolder string) error {
    finalLogFolder = currentWd + queutil.GetFilepathSeparator() + logFolder + queutil.GetFilepathSeparator() + rollingFile
    fmt.Println("rolling file location =>", finalLogFolder)

    // setup the logger
    rollingLogger = queutil.NewFlexLogger()
    // add only RollingFileLogger
    rollingLogger.AddLogger(queutil.NewRollingFileLogger(
        finalLogFolder, 1, 2, 2, false))

    return nil
}

func logNTimes(message string, numOfTimes int) error {
    for i := 0; i < numOfTimes; i++ {
        _, err := rollingLogger.Write([]byte(message + "\n"))
        if err != nil {
            return err
        }
    }
    return nil
}

func validateNumberOfRollingFiles(logFolder string, minRollingFiles int, filenamePrefix string) error {
    defer func() {
        rollingLogger.Release(nil)
    }()

    rollingFilenames := make([]string, 0)
    finalLogFolder = currentWd + queutil.GetFilepathSeparator() + logFolder

    filepath.Walk(finalLogFolder,
        func(path string, info os.FileInfo, err error) error {
            if err != nil {
                return err
            }
            if !info.IsDir() && strings.Index(info.Name(), filenamePrefix) == 0 {
                rollingFilenames = append(rollingFilenames, info.Name())
            }
            return nil
        })
    if len(rollingFilenames) >= minRollingFiles {
        fmt.Println("matched files => \n\t", rollingFilenames)
        return nil
    } else {
        return fmt.Errorf("mismatch on the expected no.of files (%v) matching the given pattern [%v]", minRollingFiles, filenamePrefix)
    }
}

// *** scenario 3
var finalLogWithOptionFile = ""
var loggerWithOptions *queutil.FlexLogger

func gotLogOptionFile(file string) error {
    finalLogWithOptionFile = currentWd + queutil.GetFilepathSeparator() + "logs" + queutil.GetFilepathSeparator() + file
    fmt.Println("file with option =>", finalLogWithOptionFile)

    return nil
}

func addLoggersForOptionTesting(loggerList string) error {
    loggerListComponents := strings.Split(loggerList, ",")

    loggerWithOptions = queutil.NewFlexLogger()
    for _, component := range loggerListComponents {
        switch component {
        case "consoleLogger":
            loggerWithOptions.AddLogger(queutil.NewConsoleLogger())
        case "rollingFileLogger":
            loggerWithOptions.AddLogger(queutil.NewRollingFileLogger(
                finalLogWithOptionFile, 1, 2, 2, true))
        default:
            return fmt.Errorf("non supoorted logger type [%v]", component)
        }
    }
    return nil
}

func logWithOptions(message, loggerName, boolOptionInString string) error {
    optionMap := make(map[string]bool)
    optionMap["consoleLogger"] = true
    optionMap["rollingFileLogger"] = true
    boolOption, err := strconv.ParseBool(boolOptionInString)
    if err != nil {
        return err
    }

    switch loggerName {
    case "consoleLogger":
        optionMap["consoleLogger"] = boolOption
    case "rollingFileLogger":
        optionMap["rollingFileLogger"] = boolOption
    }
    // log
    _, err = loggerWithOptions.WriteWithOptions([]byte(message + "\n"), optionMap, queutil.LogLevelInfo)

    return err
}

func checkIfLogFileContentsWithOptions(_, message string) error {
    defer func() {
        loggerWithOptions.Release(nil)
    }()

    bytesArr, err := queutil.ReadFileContent(finalLogWithOptionFile)
    if err != nil {
        return err
    }
    content := strings.TrimSpace(string(bytesArr))
    if strings.Index(content, message) >= 0 {
        return nil
    } else {
        return fmt.Errorf("mismatch~ expected to contain [%v], actual content is [%v]", message, content)
    }
}

// scenario 4 (log level)

var loggerForLevelTest *queutil.FlexLogger

func gotLogLevelFile(filename string) error {
    finalLogLevelFile := currentWd + queutil.GetFilepathSeparator() + "logs" + queutil.GetFilepathSeparator() + filename

    loggerForLevelTest = queutil.NewFlexLogger()
    loggerForLevelTest.AddLogger(queutil.NewRollingFileLogger(
        finalLogLevelFile, 1, 1, 1,
        true))

    return nil
}

func logMessageWithLogLevel(logLevelMethod, msg, loggerLevel string) error {
    var err error

    // logger level
    switch loggerLevel {
    case "info":
        loggerForLevelTest.LogLevel = queutil.LogLevelInfo
    case "debug":
        loggerForLevelTest.LogLevel = queutil.LogLevelDebug
    case "warn":
        loggerForLevelTest.LogLevel = queutil.LogLevelWarn
    case "err":
        loggerForLevelTest.LogLevel = queutil.LogLevelErr
    default:
        loggerForLevelTest.LogLevel = queutil.LogLevelInfo
    }

    // log method
    switch logLevelMethod {
    case "debug":
        _, err = loggerForLevelTest.Debug([]byte(msg + "\n"))
    case "info":
        _, err = loggerForLevelTest.Info([]byte(msg + "\n"))
    case "warn":
        _, err = loggerForLevelTest.Warn([]byte(msg + "\n"))
    case "err":
        _, err = loggerForLevelTest.Err([]byte(msg + "\n"))
    default:
        _, err = loggerForLevelTest.Info([]byte(msg + "\n"))
    }

    if err != nil {
        return err
    }

    return nil
}

func checkIfLogFileContentsMatchWithLevel(file, msg string) error {
    bArr, err := queutil.ReadFileContent(currentWd + queutil.GetFilepathSeparator() + "logs" + queutil.GetFilepathSeparator() + file)

    if err != nil {
        return err
    }
    lines := strings.Split(string(bArr), "\n")
    for _, line := range lines {
        if strings.Index(line, msg) >= 0 {
            return nil
        }
    }
    return fmt.Errorf("expected the file contents to contain [%v], but NOT found", msg)
}

func checkIfLogFileContentsNotMatchWithLevel(file, msg string) error {
    bArr, err := queutil.ReadFileContent(currentWd + queutil.GetFilepathSeparator() + "logs" + queutil.GetFilepathSeparator() + file)
    matched := false

    if err != nil {
        return err
    }
    lines := strings.Split(string(bArr), "\n")
    for _, line := range lines {
        if strings.Index(line, msg) >= 0 {
            matched = true
        }
    }
    if matched == true {
        return fmt.Errorf("expected the file contents NOT to contain [%v], but FOUND", msg)
    } else {
        return nil
    }
}


// ### Setup and execution ###

func FeatureContext(s *godog.Suite) {
    s.BeforeSuite(func() {
        pwd, err := os.Getwd()
        if err != nil {
            fmt.Println("could be get the current working directory", err)
            return
        }
        currentWd = pwd
    })

    s.BeforeScenario(func(i interface{}) {
        // cast back to *gherkin.Scenario
        scenario := i.(*gherkin.Scenario)
        if strings.Index(scenario.Name, "2)") == 0 {
            // clean up the log folder
            err := filepath.Walk(currentWd + queutil.GetFilepathSeparator() + "logs",
                func(path string, info os.FileInfo, err error) error {
                    if err != nil {
                        return err
                    }
                    if !info.IsDir() {
                        /*
                        fmt.Println(path)
                        fmt.Println(info.Name())
                        */
                        return os.Remove(path)
                    }
                    return nil
                })
            if err != nil {
                fmt.Println("something wrong when walking the logs folder =>", err)
            }
        }   // only clean based on certain scenario(s)
    })

    s.Step(`^a log folder named "([^"]*)"$`, gotLogFolder)
    s.Step(`^a logger is created with log file patterns as follows "([^"]*)"$`, gotLogPattern)
    s.Step(`^logging a message "([^"]*)"$`, log)
    s.Step(`^stdout would display the message PLUS the log file "([^"]*)" would also contain this entry "([^"]*)" as its last line$`, verifyLogResults)

    s.Step(`^a log file named "([^"]*)" under folder "([^"]*)"$`, gotRotationLogFile)
    s.Step(`^logging "([^"]*)" for (\d+) times$`, logNTimes)
    s.Step(`^the "([^"]*)" folder should contain at least (\d+) logs with prefix "([^"]*)"$`, validateNumberOfRollingFiles)

    s.Step(`^a log file named "([^"]*)"$`, gotLogOptionFile)
    s.Step(`^a logger is created with the following loggers "([^"]*)"$`, addLoggersForOptionTesting)
    s.Step(`^logging a message "([^"]*)" with options "([^"]*)" => "([^"]*)"$`, logWithOptions)
    s.Step(`^the console should have no log\(s\) whilst the "([^"]*)" file contains "([^"]*)"$`, checkIfLogFileContentsWithOptions)

    s.Step(`^a log file named "([^"]*)" for loglevel test$`, gotLogLevelFile)
    s.Step(`^logging a "([^"]*)" message "([^"]*)" with logLevel set to "([^"]*)"$`, logMessageWithLogLevel)
    s.Step(`^the "([^"]*)" file contains "([^"]*)"$`, checkIfLogFileContentsMatchWithLevel)
    s.Step(`^the console should have no log\(s\) whilst the "([^"]*)" file DOES NOT contains "([^"]*)"$`, checkIfLogFileContentsNotMatchWithLevel)
}

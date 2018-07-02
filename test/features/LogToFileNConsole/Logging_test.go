package LogToFileNConsole

import (
    "github.com/DATA-DOG/godog"
    "queutil"
    "os"
    "fmt"
    "strings"
    "github.com/DATA-DOG/godog/gherkin"
    "path/filepath"
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

    logger = queutil.NewFlexLogger(&queutil.FlexLoggerConfig {
        LogFileBackupCompress: true,
        LogFileMaxDaysForRetention: 2,
        LogFileMaxBackups: 2,
        LogFileMaxSizeMb: 1,
        LogFile: logFilepath,
    })

    return nil
}

// step 3
func log(message string) error {
    iWrote, err := logger.Write([]byte(message))
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

        if strings.Compare(lastLine, message) == 0 {
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
    rollingLogger = queutil.NewFlexLogger(queutil.NewFlexLoggerConfig(finalLogFolder,
        1, 2, 2, true))

    return nil
}

func logNTimes(message string, numOfTimes int) error {
    for i := 0; i < numOfTimes; i++ {
        _, err := rollingLogger.Write([]byte(message))
        if err != nil {
            return err
        }
    }
    return nil
}

func validateNumberOfRollingFiles(logFolder string, minRollingFiles int, filenamePrefix string) error {
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
}

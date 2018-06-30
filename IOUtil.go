package queutil

import (
    "strings"
    "os"
    "fmt"
)

// check if the given file exists or not
func IsFileExists(path string) bool {
    // a) is path valid?
    if len(strings.TrimSpace(path)) == 0 {
        return false
    }
    // b) check existence
    _, err := os.Stat(path)
    if os.IsExist(err) {
        return true
    }
    fmt.Println("** error?")
    fmt.Println(err)

    return false
}


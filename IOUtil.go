package queutil

import (
    "strings"
    "os"
)

// check if the given file exists or not
func IsFileExists(path string) bool {
    // a) is path valid?
    if len(strings.TrimSpace(path)) == 0 {
        return false
    }
    // b) check existence
    if _, err := os.Stat(path); os.IsExist(err) {
        return true
    }
    return false
}


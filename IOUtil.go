package queutil

import (
    "strings"
    "os"
    "os/user"
)

// check if the given file exists or not
func IsFileExists(path string) bool {
    // a) is path valid?
    if len(strings.TrimSpace(path)) == 0 {
        return false
    }
    // b) check existence
    fileInfo, err := os.Stat(path)
    if os.IsNotExist(err) {
        return false
    } else if fileInfo.IsDir() == true {
        return false
    }
    return true
}

// return the current user's home directory (string)
func GetCurrentUserHomeDir() (string, error) {
    userPtr, err := user.Current()
    if err != nil {
        return "", err
    }
    return userPtr.HomeDir, nil
}


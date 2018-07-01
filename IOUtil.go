package queutil

import (
    "strings"
    "os"
    "os/user"
    "io/ioutil"
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

// read the given file; assume the file exist check has passed before
// calling this method
func ReadFileContent(file string) ([]byte, error) {
    // assume file exist check has passed
    filePtr, err := os.OpenFile(file, os.O_RDONLY, 0755)
    if err != nil {
        return nil, err
    }
    byteArr, err := ioutil.ReadAll(filePtr)
    if err != nil {
        return nil, err
    }
    return byteArr, nil
}



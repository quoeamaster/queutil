package queutil

import (
    "strings"
    "os"
    "os/user"
    "io/ioutil"
    "syscall"
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

// write the given "content" to the "file" path
func WriteStringToFile(file string, content string) error {
    return ioutil.WriteFile(file, []byte(content), 0755)
}

// rename the given file to the target destination.
// Assume the source passed the exists check
func RenameFile(file string, targetFile string) error {
    return os.Rename(file, targetFile)
}

// method to lock a file (exclusive lock and non-blocking)
func LockFile(file string) (*os.File, error) {
    // must open the file (to gain basic access lock)
    filePtr, err := os.Open(file)
    if err != nil {
        return nil, err
    }
    err = syscall.Flock(int(filePtr.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
    if err != nil {
        return nil, err
    }
    return filePtr, nil
}

// method to unlock the given file
func UnlockFile(filePtr *os.File) error {
    if filePtr != nil {
        err := filePtr.Close()
        if err != nil {
            return err
        }
        return syscall.Flock(int(filePtr.Fd()), syscall.LOCK_UN)
    }
    return nil
}


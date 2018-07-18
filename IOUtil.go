package queutil

import (
    "strings"
    "os"
    "os/user"
    "io/ioutil"
    "github.com/theckman/go-flock"
    "fmt"
    "runtime"
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
    filePtr, err := os.OpenFile(file, os.O_RDONLY, 0444)
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
    return ioutil.WriteFile(file, []byte(content), 0444)
}

// rename the given file to the target destination.
// Assume the source passed the exists check
func RenameFile(file string, targetFile string, permission os.FileMode) error {
    err := os.Rename(file, targetFile)
    if err != nil {
        return err
    }
    // chmod (change to -r--r--r-- => 0444)
    return os.Chmod(targetFile, permission)
}

// map of lock(s)
var fileLockMap = make(map[string]*flock.Flock, 0)

// method to lock a file (exclusive lock and non-blocking)
func LockFile(file string) error {
    // simply change the mode to -r--r--r--
    flocked := flock.NewFlock(file)
    if flocked != nil {
        ok, err := flocked.TryLock()
        if err != nil {
            return err
        }
        if !ok {
            return fmt.Errorf("could not acquire file lock for [%v]", file)
        }
        fileLockMap[file] = flocked
    }
    return nil
}

// method to unlock the given file
func UnlockFile(file string) error {
    // reference on unix file mode => https://www.tutorialspoint.com/unix/unix-file-permission.htm
    //return os.Chmod(file, 0755)
    flocked := fileLockMap[file]
    if flocked != nil {
        return flocked.Unlock()
    }
    return nil
}

// method to return the correct filepath separator based on the OS
func GetFilepathSeparator() string {
    if strings.Compare(runtime.GOOS, "windows") == 0 {
        return "\\"
    } else {
        return "/"
    }
}

// method to write the given []byte to a file
func WriteByteArrayToFile (p []byte, file string) error {
    return ioutil.WriteFile(file, p, 0444)
}


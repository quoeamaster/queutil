package FileLocks

import (
    "github.com/DATA-DOG/godog"
    "os"
    "queutil"
    "fmt"
    "syscall"
    "io/ioutil"
    "strings"
    "github.com/theckman/go-flock"
)

var targetFilePtr *os.File
var targetFile string
var flocked *flock.Flock

// ****************
// ** scenario 1 **
// ****************

// step 1
func gotLocalFile(file string) error {
    if queutil.IsFileExists(file) == false {
        return fmt.Errorf("file doesn't exists~ [%v]", file)
    }
    ptr, err := os.OpenFile(file, syscall.O_RDONLY, 0444)
    if err != nil {
        return err
    }
    targetFilePtr = ptr
    targetFile = file

    return nil
}

// step 2
func contentOfFileCheck(content string) error {
    if targetFilePtr != nil {
        byteArr, err := ioutil.ReadAll(targetFilePtr)
        if err != nil {
            return err
        }
        rContent := strings.TrimSpace(string(byteArr))
        if strings.Compare(rContent, content) == 0 {
            return nil
        } else {
            return fmt.Errorf("content values mismatch, expecte [%v] but got [%v]", content, rContent)
        }
    }
    return fmt.Errorf("file pointer is nil")
}

// step 3
func lockFileWithPermission(permission string) error {
    //return queutil.LockFile(targetFile)

    // lock using library
    /*
    f := flock.NewFlock(targetFile)
    flocked = f
    ok, err := f.TryLock()
    if err != nil {
        return err
    }
    if !ok {
        return fmt.Errorf("lock could not be acquired, possibly some process already locked the file")
    }
    return nil
    */
    return queutil.LockFile(targetFile)
}

// step 4
func contentOfFileCheckAfterLock(content string) error {
    // ***** TODO could not re-use the previous file Pointer as it doesn't make sense after chmod~~~~
    ptr, err := os.OpenFile(targetFile, syscall.O_RDONLY, 0444)
    if err != nil {
        return err
    }
    byteArr, err := ioutil.ReadAll(ptr)
    if err != nil {
        return err
    }
    rContent := strings.TrimSpace(string(byteArr))
    if strings.Compare(rContent, content) == 0 {
        return nil
    }
    return fmt.Errorf("mismatch~ [%v] vs [%v]", content, rContent)
}

// step 4
func filePermissionCheck(permission string) error {
    fileInfo, err := os.Stat(targetFile)
    if err != nil {
        return err
    }
    fileMode := fileInfo.Mode()
    if strings.Compare("-r--r--r--", fileMode.String()) == 0 {
        return nil
    } else {
        return fmt.Errorf("the permission mismatched~ expected [%v] got [%v]", "-r--r--r--", fileMode.String())
    }
}

// step 5
func lockFileAgain() error {
    err := lockFileWithPermission("")
    if err != nil {
        fmt.Println (err)
        return nil
    }
    return fmt.Errorf("something is wrong... should not be able to acquire the lock again")
}



func FeatureContext(s *godog.Suite) {
    s.Step(`^a local file named "([^"]*)"$`, gotLocalFile)
    s.Step(`^read the content of the file and should return "([^"]*)"$`, contentOfFileCheck)
    s.Step(`^lock the file with rights "([^"]*)"$`, lockFileWithPermission)
    s.Step(`^now the file is still readable and content is still "([^"]*)"$`, contentOfFileCheckAfterLock)
    s.Step(`^the file\'s permission is "([^"]*)"$`, filePermissionCheck)
    s.Step(`^try to acquire the lock on the same file again would got exception$`, lockFileAgain)
}

package queutil

import (
    "strings"
    "fmt"
)

// method to create an error interface / object based on the given string
func CreateErrorWithString (err string) error {
    return fmt.Errorf(fmt.Sprintf("%v", err))
}

// method to check if Client.Timeout has occured
func IsHttpClientTimeoutError (err error) bool {
    if err != nil {
        idx := strings.Index(strings.ToLower(err.Error()), "client.timeout")
        if idx >= 0 {
            return true
        }
    }
    return false
}

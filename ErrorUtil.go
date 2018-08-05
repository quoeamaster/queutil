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

// a method for checking if it is an EOF error; when REST api contains a body,
// there might be chances to receive a "EOF" even though it is valid in the content
func IsHttpRequestValidEOFError (err error, httpRequestContentLength int) bool {
    if err == nil {
        return false
    }
    if httpRequestContentLength > 0 && strings.Index(err.Error(), "EOF") == 0 {
        return true
    }
    return false
}
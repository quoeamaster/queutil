package queutil

import "strings"

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

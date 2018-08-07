package queutil

import "strings"

// method to check if the given string is empty or not
func IsStringEmpty (value string) bool {
    return len(strings.TrimSpace(value)) == 0
}

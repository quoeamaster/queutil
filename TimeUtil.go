package queutil

import "time"

// create a time.Duration struct based on the given string;
// must have parsing error.
// valid strings could be => "5s", "100ms"
func CreateTimeoutByString (timeInString string) (time.Duration, error) {
    timeout, err := time.ParseDuration(timeInString)
    if err != nil {
        return *new(time.Duration), err
    }
    return timeout, nil
}

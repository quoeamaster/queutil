package queutil

import "github.com/kjk/betterguid"

// return a generated UUID based on betterguid library
// (timestamp + 72 bits of random characters)
func GenerateUUID() string {
    return betterguid.New()
}


package queutil

import (
    "bytes"
    "fmt"
)

// builds the api url based on the given info.
// PS. note that it only builds the url not the http request itself;
// you still to interact with a Http.Client instance
func BuildGenericApiUrl (seedAddr, protocol, apiEndpoint string) string {
    var b bytes.Buffer

    // if empty... assume normal http://
    if IsStringEmpty(protocol) {
        b.WriteString("http://")
    } else {
        b.WriteString(protocol)
    }
    b.WriteString(fmt.Sprintf("%v/%v", seedAddr, apiEndpoint))

    return b.String()
}

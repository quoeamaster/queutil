package queutil

import (
    "net/http"
    "time"
)

// * duration and client creation sample code
// timeout, err := time.ParseDuration("5s")
//    if err != nil {
//        return nil
//    }
//    client = queutil.GenerateHttpClient(timeout, nil, nil, nil)
//
//    res, err := client.Get(url)
//
// * cleanup
// defer func() {
//    res.Body.Close()
// }()

// method to create an HttpClient
func GenerateHttpClient (
    timeout time.Duration,
    transport http.RoundTripper,
    checkRedirectFunc func(req *http.Request, via []*http.Request) error,
    cookies http.CookieJar) *http.Client {

    client := http.Client{}
    client.Timeout = timeout
    if transport != nil {
        client.Transport = transport
    }
    if checkRedirectFunc != nil {
        client.CheckRedirect = checkRedirectFunc
    }
    if cookies != nil {
        client.Jar = cookies
    }
    return &client
}




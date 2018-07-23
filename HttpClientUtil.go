package queutil

import (
    "net/http"
    "time"
)

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




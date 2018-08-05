package queutil

import (
    "net/http"
    "time"
    "fmt"
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

// method to create an HttpClient.
// The important thing is that the caller should close the response body
// after processing to prevent memory / resource leak
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

// method to get back the httpRequest content; could be a valid []byte or
// an empty []byte or
// nil + error
func GetHttpRequestContent (req *http.Request) ([]byte, error) {
    if req == nil {
        return nil, CreateErrorWithString(fmt.Sprintf("http request is not valid~ [%v]", req))
    }
    contentLen := int(req.ContentLength)
    if contentLen > 0 {
        bArr := make([]byte, contentLen)
        _, err := req.Body.Read(bArr)

        // is it a valid EOF exception?
        if IsHttpRequestValidEOFError(err, contentLen) {
            return bArr, nil
        } else {
            return nil, err
        }
    } else {
        // really empty entry; hence return empty []byte
        return []byte {}, nil
    }
}




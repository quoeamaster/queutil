package HttpClient

import (
    "github.com/DATA-DOG/godog"
    "strings"
    "fmt"
    "net/http"
    "queutil"
    "io/ioutil"
)

type webserviceParams struct {
    Key string
    Value string
}

var srvUrl string
var srvParams []webserviceParams
var client *http.Client
var responseString string

func getWebServiceUrl(url string) error {
    srvUrl = url
    srvParams = make([]webserviceParams, 0)

    return nil
}

func getParams(paramsInString string) error {
    paramS := paramsInString[1:len(paramsInString)-1]
    paramKVs := strings.Split(paramS,",")

    if len(paramKVs) <= 0 {
        return fmt.Errorf("should have at least one pair of params; but got this %v", paramKVs)
    }
    for _, kv := range paramKVs {
        kvArr := strings.Split(kv, ":")
        if len(kvArr) == 2 {
            param := webserviceParams{}
            param.Key = kvArr[0]
            param.Value = kvArr[1]

            srvParams = append(srvParams, param)
        }
    }
    // create httpClient  ʕ·ᴥ·ʔ
    timeout, err := queutil.CreateTimeoutByString("5s")
    if err != nil {
        return err
    }
    // timeout, err := time.ParseDuration("5s")

    client = queutil.GenerateHttpClient(timeout, nil, nil, nil)

    return nil
}

func getJsonResponse() error {
    // build the final url
    finalParams := ""
    for idx, param := range srvParams {
        if idx == 0 {
            finalParams = fmt.Sprintf("%v?%v=%v", finalParams, param.Key, param.Value)
        } else {
            finalParams = fmt.Sprintf("%v&%v=%v", finalParams, param.Key, param.Value)
        }
    }
    finalUrl := fmt.Sprintf("%v%v", srvUrl, finalParams)
    // get result
    res, err := client.Get(finalUrl)
    if err != nil {
        if queutil.IsHttpClientTimeoutError(err) {
            fmt.Println("** timeout on request **")
        }
        return err
    }

    byteArr, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return err
    }
    responseString = string(byteArr)

    // TODO: MUST cleanup
    defer func() {
        res.Body.Close()
    }()

    return nil
}

func isWordFoundInResponse(word string) error {
    // no brain-er approach, indexOf
    idx := strings.Index(responseString, word)
    if idx >= 0 {
        fmt.Printf("found at %v\n", idx)
        return nil
    }
    return fmt.Errorf("word not found in the result! Expected [%v] within [%v]", word, responseString)
}

func FeatureContext(s *godog.Suite) {
    // cleanup
    s.BeforeScenario(func(i interface{}) {
        srvUrl = ""
        srvParams = nil
        client = nil
    })

    s.Step(`^a webservice url "([^"]*)"$`, getWebServiceUrl)
    s.Step(`^parameters provided as "([^"]*)"$`, getParams)
    s.Step(`^calling the api would resulted a Json response$`, getJsonResponse)
    s.Step(`^"([^"]*)" is found within the response$`, isWordFoundInResponse)
}

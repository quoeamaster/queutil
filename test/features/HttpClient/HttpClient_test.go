package HttpClient

import "github.com/DATA-DOG/godog"

func getWebServiceUrl(url string) error {
    return godog.ErrPending
}

func getParams(paramsInString string) error {
    return godog.ErrPending
}

func getJsonResponse() error {
    return godog.ErrPending
}

func isWordFoundInResponse(word string) error {
    return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
    s.Step(`^a webservice url "([^"]*)"$`, getWebServiceUrl)
    s.Step(`^parameters provided as "([^"]*)"$`, getParams)
    s.Step(`^calling the api would resulted a Json response$`, getJsonResponse)
    s.Step(`^"([^"]*)" is found within the response$`, isWordFoundInResponse)
}

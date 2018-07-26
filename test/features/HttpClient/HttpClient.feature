Feature: Http Client features (able to send and handle http request; the result is a Json response in this case)
    some software components would need to access the internet for data input
    or for call a webservice to fulfill a request

    Assumptions for the feature test:
    - request could be in any format; however in this test case, only Json request is handled
    - response could be in any format as well; however again, only Json response in handled here

    Major use cases:
    - make a request to a website / webservice in json and get back the response in json
    - able to parse the result as well

    api url:
    - glosbe.com
    - https://glosbe.com/gapi/tm?from=pol&dest=eng&format=json&phrase=borsuk&page=1&pretty=true

    Scenario: 1) basic request
        Given a webservice url "https://glosbe.com/gapi/tm"
        When parameters provided as "[from:eng,dest:eng,format:json,phrase:soccer,page:1,pretty:true]"
        Then calling the api would resulted a Json response
        And "football" is found within the response


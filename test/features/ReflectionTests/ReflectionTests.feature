Feature:
    to test reflection related scenarios

    Assumptions for the feature test:
    - nil

    Major use cases:
    - invoke a normal method with a non-nil instance, no parameter(s) required and return no value
    - invoke a normal method with a non-nil instance, no parameter(s) required and return a single value
    - invoke a normal method with a non-nil instance, no parameter(s) required and return an array of values

    - invoke a normal method with a non-nil instance, a list of parameters and return no value
    - invoke a normal method with a non-nil instance, a list of parameters and return a single value
    - invoke a normal method with a non-nil instance, a list of parameters and return an array of values

    - invoke a missing method with a non-nil instance, no parameter(s) required and return no value
    - invoke a normal method with a nil instance, no parameter(s) required and return no value
    - invoke a non-public (e.g. private) method with a non-nil instance, no parameter(s) required and return no value

    - invoke a normal method with a MAP object with string key (method name)
        and function value (pointing to some valid function pointer), no parameter(s) required and return no value
    - invoke a normal method with a MAP object with string key (method name)
        and function value (pointing to some valid function pointer), a list of parameters and return an array of values


    Scenario: 1) non-nil instance, no parameter(s) required, return no value
        Given a method named "PrintHelloToConsole"
        When an instance of "MethodTester_1" is provided
        Then triggering the method would yield 0 returned values
        Then the method execution has no error

    Scenario: 2) non-nil instance, no parameter(s) required, return 1 value
        Given a method named "GetSingleValue"
        When an instance of "MethodTester_1" is provided
        Then triggering the method would yield 1 returned values
        Then the method execution has no error

    Scenario: 3) non-nil instance, no parameter(s) required, return array of values
        Given a method named "GetMultiValue"
        When an instance of "MethodTester_1" is provided
        Then triggering the method would yield 4 returned values
        Then the method execution has no error

    Scenario: 4) non-nil instance, parameter provided, return no value
        Given a method named "PrintHelloToConsoleWithParams"
        When parameters are provided [{ "key":"firstName","value":"wong","type":"string" } ,{ "key":"lastName","value":"jason","type":"string" } ,{ "key":"yearsInES","value":4,"type":"int" }]
        And an instance of "MethodTester_2" is provided
        Then triggering the method would yield 0 returned values
        Then the method execution has no error

    Scenario: 5) non-nil instance, parameter provided, return 1 value
        Given a method named "GetSingleValue"
        When parameters are provided [{ "key":"firstName","value":"wong","type":"string" } ,{ "key":"lastName","value":"jason","type":"string" } ,{ "key":"yearsInES","value":4,"type":"int" }]
        And an instance of "MethodTester_2" is provided
        Then triggering the method would yield 1 returned values
        Then the method execution has no error

    Scenario: 6) non-nil instance, parameter provided, return array of values
        Given a method named "GetMultiValue"
        When parameters are provided [{ "key":"firstName","value":"wong","type":"string" } ,{ "key":"lastName","value":"jason","type":"string" } ,{ "key":"yearsInES","value":4,"type":"int" }]
        And an instance of "MethodTester_2" is provided
        Then triggering the method would yield 2 returned values
        Then the method execution has no error

    Scenario: 7) non-nil instance, missing method
        Given a method named "UnknownMethod"
        When parameters are provided [{ "key":"firstName","value":"wong","type":"string" } ,{ "key":"lastName","value":"jason","type":"string" } ,{ "key":"yearsInES","value":4,"type":"int" }]
        And an instance of "MethodTester_2" is provided
        Then the method execution has error

    Scenario: 8) nil instance
        Given a method named "GetSingleValue"
        And an instance of "null" is provided
        Then the method execution has error

    Scenario: 9) non-nil instance, no parameter(s) required, return no value (private method)
        Given a method named "printInPrivate"
        When an instance of "MethodTester_1" is provided
        Then triggering the method would yield 0 returned values
        Then the method execution has error

    # Scenario: 10) map instance, normal execution (key => method name; value => func pointer)
    #    Given a method named "TestFuncWrappingByMap"
    #    When an instance of "map" is provided
    #    Then the method execution has no error
    #    Then triggering the method would yield 1 returned values

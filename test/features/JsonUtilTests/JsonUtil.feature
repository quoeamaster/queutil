Feature:
    to test json creation utility; though golang already have a json
    parser and creator, however you would need to create a struct with
    annotations which might not be a very good solution when the struct's
    contents are ever changing (you need to change code simply).

    JsonUtil.go tries to create the corresponding json based on a "smooth" API,
    which means you can create different json structures based on different
    situations.

    Assumptions for the feature test:
    - nil

    Major use cases:
    - create a json (string) that is parse-able (which means it is valid) based on different parameters provided


    Scenario: 1) create a json (normal use case)
        Given a scenario to start with
        When key "firstname" and string value "Huang" is given
        Then key "lastname" and string value "Json" is given
        Then key "age" and int value 26 is given
        Then key "member" and bool value true is given
        Then key "loyalty" and float32 value 7.5 is given
        Then key "satisfaction" and float64 value 9.0 is given
        Then key "hobby" and array value [{"x": 1, "y": 2} ,{"x": 3, "y": 4}] is given
        Then close the scenario
        And result of the json created should be {"firstname": "Huang","lastname": "Json","age": 26,"member": true,"loyalty": 7.5,"satisfaction": 9,"hobby": [{"X": 1,"Y": 2},{"X": 3,"Y": 4}]}

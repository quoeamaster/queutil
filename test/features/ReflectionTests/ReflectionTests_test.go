package ReflectionTests

import (
    "github.com/DATA-DOG/godog"
    "fmt"
    "queutil"
    "reflect"
    "strings"
    "encoding/json"
)

var methodName string
var generalErr error
var generalReturnValues []reflect.Value

// ****************
// ** scenario 1 **
// ****************

func gotMethodName(method string) error {
    methodName = method
    return nil
}

func gotTestingObjectName(testId string) error {
    // TODO: create the instance based on the value of testId
    switch testId {
    case "MethodTester_1":
        testMethodTester_1()
    case "MethodTester_2":
        testMethodTester_2()
    case "null": {
        generalReturnValues, generalErr = queutil.InvokeMethodCallByReflection(
            methodName, nil, nil)
    }
    case "map": {
        testMethodTester_map()
    }
    default:
        return fmt.Errorf("unknown tester [%v]", testId)
    }

    if generalErr != nil {
        //return generalErr
    }
    return nil
}

func testMethodTester_1 () {
    instancePtr := new(MethodTester_1)

    generalReturnValues, generalErr = queutil.InvokeMethodCallByReflection(
        methodName, instancePtr, nil)
}
func testMethodTester_2 () {
    instancePtr := new(MethodTester_2)

    generalReturnValues, generalErr = queutil.InvokeMethodCallByReflection(
        methodName, instancePtr, generalInParams)
}
func testMethodTester_map() {
    instancePtr := new(MethodTester_1)
    // map of default func (empty param)
    mapPtr := make(map[string]func() int)
    mapPtr["TestFuncWrappingByMap"] = instancePtr.GetSingleValue

    generalReturnValues, generalErr = queutil.InvokeMethodCallByReflection(
        methodName, mapPtr, generalInParams)

}

func invokeMethodAndCheckNumOfReturnValues(num int) error {
    // check the number of return values
    if num == len(generalReturnValues) {
        if num > 0 {
            fmt.Printf("** number of return values [%v] => ", num)
            for idx, val := range generalReturnValues {
                fmt.Printf(" { idx: %v, val: %v } ", idx, reflect.Indirect(val))
            }
            fmt.Printf("\n")
        }
        return nil
    } else {
        return fmt.Errorf(
            "expected return values to be [%v] but got [%v]; {%v}",
            num, len(generalReturnValues), generalReturnValues)
    }
}

func noErrorExecution() error {
    if generalErr != nil {
        return generalErr
    }
    return nil
}

// ****************
// ** scenario 4 **
// ****************

var generalInParams []reflect.Value

func gotMethodParams(paramString string) error {
    generalInParams = make([]reflect.Value, 0)
    kvArray := strings.Split(paramString, " ,")
    for _, kv := range kvArray {
        jmap := make(map[string]interface{})
        err := json.Unmarshal([]byte(kv), &jmap)
        if err != nil {
            return err
        }
        // create a new reflect.Value for the given parameter
        generalInParams = append(generalInParams, reflect.ValueOf(jmap["value"]))
    }
    return nil
}

// ** scenario 7 **

func gotErrorExecution() error {
    if generalErr != nil {
        fmt.Printf("## error was found which is expected result => %v\n", generalErr)
        return nil
    }
    return fmt.Errorf("should HAVE error... somehow... all good")
}


func FeatureContext(s *godog.Suite) {
    // lifecycle triggers
    s.BeforeScenario(func(i interface{}) {
        methodName = ""
        generalErr = nil
        generalReturnValues = nil
        generalInParams = nil
    })

    s.Step(`^a method named "([^"]*)"$`, gotMethodName)
    s.Step(`^an instance of "([^"]*)" is provided$`, gotTestingObjectName)
    s.Step(`^triggering the method would yield (\d+) returned values$`, invokeMethodAndCheckNumOfReturnValues)
    s.Step(`^the method execution has no error$`, noErrorExecution)

    s.Step(`^parameters are provided \[(.*)\]$`, gotMethodParams)

    s.Step(`^the method execution has error$`, gotErrorExecution)
}


// *********************
// ** Testing objects **
// *********************

type MethodTester_1 struct {}
func (m *MethodTester_1) PrintHelloToConsole () {
    fmt.Println("** MethodTester_1 => PrintHelloToConsole")
}
func (_ *MethodTester_1) GetSingleValue () int {
    return 1001
}
func (m *MethodTester_1) GetMultiValue () (int, string, []int, map[string]string) {
    map1 := make(map[string]string)
    map1["key1"] = "value1"
    map1["key2"] = "whatever ~"

    return 2001, "helloWorld~", []int { 1, 2, 3 }, map1
}
func (_ *MethodTester_1) printInPrivate () {
    fmt.Println("** NEVER could reach and print in here **")
}

type MethodTester_2 struct {}
func (_ *MethodTester_2) PrintHelloToConsoleWithParams (firstName string, lastName string, yearsInES float64) {
    fmt.Printf("%v %v has %v years experience in ES.\n", firstName, lastName, yearsInES)
}
func (_ *MethodTester_2) GetSingleValue (firstName string, lastName string, yearsInES float64) string {
    return fmt.Sprintf("%v %v has %v years experience in ES.", lastName, firstName, yearsInES)
}
func (_ *MethodTester_2) GetMultiValue (firstName string, lastName string, yearsInES float64) (string, int) {
    return fmt.Sprintf("%v %v has %v years experience in ES.", lastName, firstName, yearsInES), 3
}

type IPrinter interface {
    PrintToConsole (msg string) (numOfCharsPrinted int)
}
type MethodTesterA struct {}
func (_ *MethodTesterA) PrintToConsole (msg string) int {
    fmt.Println("## inside A ::", msg)
    return len(msg)
}
type MethodTesterB struct {}
func (_ *MethodTesterB) PrintToConsole (msg string) int {
    fmt.Println("* inside B")
    fmt.Println(msg)
    return len(msg)
}



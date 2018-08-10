package JsonUtilTests

import (
    "testing"
    "github.com/DATA-DOG/godog"
    "os"
    "queutil"
    "strconv"
    "strings"
    "encoding/json"
    "fmt"
    "bytes"
)

var buf bytes.Buffer

func init() {
    //godog.BindFlags("godog.", flag.CommandLine, &opt)
}
func TestMain(m *testing.M) {
    status := godog.RunWithOptions("godog", func(s *godog.Suite) {
        FeatureContext(s)
    }, godog.Options{
        Format:    "pretty",
        Paths:     []string{"./"},
        // Randomize: time.Now().UTC().UnixNano(), // randomize scenario execution order
    })

    if st := m.Run(); st > status {
        status = st
    }
    os.Exit(status)
}

func scenarioStarts() error {
    buf = queutil.BeginJsonStructure(buf)
    return nil
}

func addStringValue(key, value string) error {
    buf = queutil.AddStringToJsonStructure(buf, key, value)
    return nil
}

func addIntValue(key string, value int) error {
    buf = queutil.AddIntToJsonStructure(buf, key, value)
    return nil
}

func addBoolValue(key string, value string) error {
    bVal, err := strconv.ParseBool(value)
    if err != nil {
        return err
    }
    buf = queutil.AddBoolToJsonStructure(buf, key, bVal)
    return nil
}

func addFloatValue(key string, floatBase int, value string) error {
    if floatBase == 32 {
        fVal, err := strconv.ParseFloat(value, 32)
        if err != nil {
            return err
        }
        buf = queutil.AddFloat32ToJsonStructure(buf, key, float32(fVal))

    } else if floatBase == 64 {
        fVal, err := strconv.ParseFloat(value, 64)
        if err != nil {
            return err
        }
        buf = queutil.AddFloat64ToJsonStructure(buf, key, fVal)
    }

    return nil
}

func addArrayValue(key, jsonString string) error {
    // parse json string back to interface{}
    objList := strings.Split(jsonString, " ,")
    //XYPairList := make([]XYPair, 0)
    XYPairList := make([]interface{}, 0)

    for _, objString := range objList {
        xyPair := new(XYPair)
        err := json.Unmarshal([]byte(objString), xyPair)
        if err != nil {
            return err
        }
        XYPairList = append(XYPairList, xyPair)
    }
    buf = queutil.AddArrayToJsonStructure(buf, key, XYPairList)

    return nil
}

func scenarioEnds() error {
    buf = queutil.EndJsonStructure(buf)
    return nil
}

func verifyResults(result string) error {
    if strings.Compare(buf.String(), result) == 0 {
        return nil
    }
    return fmt.Errorf("expected [%v] but got [%v]", result, buf.String())
}


// ** scenario 2 **

var user User

func objectScenarioStarts() error {
    user = *new(User)
    return nil
}

func setFieldTo(fieldName, value string) error {
    switch fieldName {
    case "FirstName":
        user.FirstName = value
    case "LastName":
        user.LastName = value
    default:
        return fmt.Errorf("unknown field %v", fieldName)
    }
    return nil
}
func setFieldToNumber(fieldName string, value int) error {
    switch fieldName {
    case "Age":
        user.Age = value
    default:
        return fmt.Errorf("unknown field %v", fieldName)
    }
    return nil
}
func doneWithFieldSetup() error {
    buf = queutil.ConvertInterfaceToJsonStructure(buf, user)

    return nil
}

func FeatureContext(s *godog.Suite) {
    s.BeforeScenario(func(i interface{}) {
        buf.Reset()
    })

    s.Step(`^a scenario to start with$`, scenarioStarts)
    s.Step(`^key "([^"]*)" and string value "([^"]*)" is given$`, addStringValue)
    s.Step(`^key "([^"]*)" and int value (\d+) is given$`, addIntValue)
    s.Step(`^key "([^"]*)" and bool value (.*) is given$`, addBoolValue)
    s.Step(`^key "([^"]*)" and float(\d+) value (.*) is given$`, addFloatValue)
    s.Step(`^close the scenario$`, scenarioEnds)
    s.Step(`^key "([^"]*)" and array value \[(.*)\] is given$`, addArrayValue)
    s.Step(`^result of the json created should be (.*)$`, verifyResults)

    s.Step(`^an Object with several fields set-able$`, objectScenarioStarts)
    s.Step(`^set field "([^"]*)" to "([^"]*)"$`, setFieldTo)
    s.Step(`^set field "([^"]*)" to number (\d+)$`, setFieldToNumber)
    s.Step(`^done with field setup$`, doneWithFieldSetup)
}


type XYPair struct {
    X int
    Y int
}
func (xy *XYPair) JsonString () string {
    var b bytes.Buffer

    b = queutil.BeginJsonStructure(b)
    b = queutil.AddIntToJsonStructure(b, "X", xy.X)
    b = queutil.AddIntToJsonStructure(b, "Y", xy.Y)
    b = queutil.EndJsonStructure(b)

    return b.String()
}
func (xy *XYPair) String() string {
    return xy.JsonString()
}

type User struct {
    FirstName string `json:"FirstName"`
    LastName string `json:"LastName"`
    Age int `json:"Age"`
}
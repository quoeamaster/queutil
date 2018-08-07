package JsonUtilTests

import (
    "github.com/DATA-DOG/godog"
    "bytes"
    "queutil"
    "strconv"
    "strings"
    "fmt"
)

var buf bytes.Buffer

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
    s.Step(`^result of the json created should be (.*)$`, verifyResults)
}
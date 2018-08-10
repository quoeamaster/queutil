package queutil

import (
    "bytes"
    "strings"
    "fmt"
    "reflect"
)

// appends the "{" to the buffer
func BeginJsonStructure (buf bytes.Buffer) bytes.Buffer {
    buf.WriteString("{")
    return buf
}

// appends the "}" to the buffer
func EndJsonStructure (buf bytes.Buffer) bytes.Buffer {
    newBuf := removeTrailingCommaFromJsonStructure(buf)
    newBuf.WriteString("}")
    // kill the old buffer
    buf.Reset()

    return newBuf
}

// appends the given key, value pair to the buffer
func AddStringToJsonStructure (buf bytes.Buffer, key, value string) bytes.Buffer {
    if IsStringEmpty(key) {
        return buf
    }
    // key
    buf.WriteString("\"")
    buf.WriteString(key)
    buf.WriteString("\": ")
    // value
    if IsStringEmpty(value) {
        buf.WriteString("null")
    } else {
        buf.WriteString("\"")
        buf.WriteString(value)
        buf.WriteString("\",")
    }
    return buf
}

func AddIntToJsonStructure (buf bytes.Buffer, key string, value int) bytes.Buffer {
    if IsStringEmpty(key) {
        return buf
    }
    // key and value
    buf.WriteString("\"")
    buf.WriteString(key)
    buf.WriteString("\": ")
    buf.WriteString(fmt.Sprintf("%v,", value))
    return buf
}

func AddFloat32ToJsonStructure (buf bytes.Buffer, key string, value float32) bytes.Buffer {
    if IsStringEmpty(key) {
        return buf
    }
    // key and value
    buf.WriteString("\"")
    buf.WriteString(key)
    buf.WriteString("\": ")
    buf.WriteString(fmt.Sprintf("%v,", value))
    return buf
}

func AddFloat64ToJsonStructure (buf bytes.Buffer, key string, value float64) bytes.Buffer {
    if IsStringEmpty(key) {
        return buf
    }
    // key and value
    buf.WriteString("\"")
    buf.WriteString(key)
    buf.WriteString("\": ")
    buf.WriteString(fmt.Sprintf("%v,", value))
    return buf
}

func AddBoolToJsonStructure (buf bytes.Buffer, key string, value bool) bytes.Buffer {
    if IsStringEmpty(key) {
        return buf
    }
    // key and value
    buf.WriteString("\"")
    buf.WriteString(key)
    buf.WriteString("\": ")
    buf.WriteString(fmt.Sprintf("%v,", value))
    return buf
}

// assume the interface{} String() returns a valid json as well...
// or else it won't work
func AddArrayToJsonStructure (buf bytes.Buffer, key string, values []interface{}) bytes.Buffer {
    if IsStringEmpty(key) {
        return buf
    }
    // key and value
    buf.WriteString("\"")
    buf.WriteString(key)
    buf.WriteString("\": [")

    for idx, val := range values {
        if idx > 0 {
            buf.WriteString(",")
        }
        // cast back to IJsonStringAble or Stringer
        // ref: https://stackoverflow.com/questions/27803654/explanation-of-checking-if-value-implements-interface
        ijson, ok := val.(fmt.Stringer)
        if ok {
            buf.WriteString(ijson.String())

        } else {
            fmt.Println("unknown...", ijson)
        }
    }
    buf.WriteString("]")
    return buf
}


// TODO: add array syntax and map (object) syntax too

func BeginObjectJsonStructure (buf bytes.Buffer, key string) bytes.Buffer {
    if IsStringEmpty(key) {
        return buf
    }
    // add the key and open "{"
    buf.WriteString("\"")
    buf.WriteString(key)
    buf.WriteString("\": {")

    return buf
}

func EndObjectJsonStructure (buf bytes.Buffer) bytes.Buffer {
    // add the key and open "{"
    buf.WriteString("}")

    return buf
}

// helper method to append a "," if necessary
func removeTrailingCommaFromJsonStructure (buf bytes.Buffer) bytes.Buffer {
    lastChar := string(buf.Bytes()[buf.Len()-1])
    if strings.Compare(lastChar, ",") == 0 {
        newBuf := bytes.NewBuffer(buf.Bytes()[0:buf.Len()-1])
        return *newBuf
    }
    return buf
}

// experimental method to convert a given interface to a json string.
// Only support simple, basic type(s) { int, float32, float64, string, bool }
func ConvertInterfaceToJsonStructure (buf bytes.Buffer, value interface{}) bytes.Buffer {
    var b bytes.Buffer
    if value != nil {
        valType := reflect.TypeOf(value)
        valValue := reflect.ValueOf(value)

        b = BeginJsonStructure(b)
        for i := 0; i < valType.NumField(); i++ {
            fieldMeta := valType.Field(i)
            targetFieldName := fieldMeta.Tag.Get("json")
            fieldData := valValue.Field(i)

            switch fieldData.Kind() {
            case reflect.String:
                b = AddStringToJsonStructure(b, targetFieldName, fieldData.String())
            case reflect.Int:
                b = AddIntToJsonStructure(b, targetFieldName, int(fieldData.Int()))
            case reflect.Float32:
                b = AddFloat32ToJsonStructure(b, targetFieldName, float32(fieldData.Float()))
            case reflect.Float64:
                b = AddFloat64ToJsonStructure(b, targetFieldName, fieldData.Float())
            case reflect.Bool:
                b = AddBoolToJsonStructure(b, targetFieldName, fieldData.Bool())
            default:
                // including slice (due to the lack of ability to cast value to a static string listed type value)
                fmt.Printf("non supported data type [%v]\n", fieldData.Kind())
            }
        }
        b = EndJsonStructure(b)
    }
    return b
}

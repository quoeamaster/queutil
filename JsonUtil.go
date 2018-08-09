package queutil

import (
    "bytes"
    "strings"
    "fmt"
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
    if isStringEmpty(key) {
        return buf
    }
    // key
    buf.WriteString("\"")
    buf.WriteString(key)
    buf.WriteString("\": ")
    // value
    if isStringEmpty(value) {
        buf.WriteString("null")
    } else {
        buf.WriteString("\"")
        buf.WriteString(value)
        buf.WriteString("\",")
    }
    return buf
}

func AddIntToJsonStructure (buf bytes.Buffer, key string, value int) bytes.Buffer {
    if isStringEmpty(key) {
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
    if isStringEmpty(key) {
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
    if isStringEmpty(key) {
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
    if isStringEmpty(key) {
        return buf
    }
    // key and value
    buf.WriteString("\"")
    buf.WriteString(key)
    buf.WriteString("\": ")
    buf.WriteString(fmt.Sprintf("%v,", value))
    return buf
}

// TODO: add array syntax and map (object) syntax too

// assume the interface{} String() returns a valid json as well...
// or else it won't work
func AddArrayToJsonStructure (buf bytes.Buffer, key string, values []interface{}) bytes.Buffer {
    if isStringEmpty(key) {
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

func BeginObjectJsonStructure (buf bytes.Buffer, key string) bytes.Buffer {
    if isStringEmpty(key) {
        return buf
    }
    // add the key and open "{"
    buf.WriteString("\"")
    buf.WriteString(key)
    buf.WriteString("\": {")

    return buf
}

func EndObjectJsonStructure (buf bytes.Buffer, key string) bytes.Buffer {
    if isStringEmpty(key) {
        return buf
    }
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

// method to check if the given string is empty or not (trim + len == 0 check)
func isStringEmpty (val string) bool {
    if len(strings.TrimSpace(val)) == 0 {
        return true
    }
    return false
}
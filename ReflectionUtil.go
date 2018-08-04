package queutil

import (
    "reflect"
    "fmt"
)

// invoke an instance's method through the given parameter(s)
func InvokeMethodCallByReflection (
    methodName string,
    instance interface{},
    methodParams []reflect.Value) (returnValues []reflect.Value, err error) {

    // final try / catch
    defer func() {
        r := recover()
        if r != nil {
            err = CreateErrorWithString(fmt.Sprintf("%v", r))
        }
    }()

    // instance must be valid
    if instance == nil {
        err = CreateErrorWithString(fmt.Sprintf("instance to search for the given method (%v) is nil", methodName))
        return nil, err
    }

    // get back the method reference
    methodRef := reflect.ValueOf(instance).MethodByName(methodName)
    if methodRef.IsNil() {
        err = CreateErrorWithString(fmt.Sprintf("[%v] is not available for the given object [%v]", methodName, instance))
        return nil, err
    }

    // invoke the method simply
    returnValues = methodRef.Call(methodParams)

    return returnValues, err
}

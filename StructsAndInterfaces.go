package queutil

// acts as a common interface for "releasable" object(s)
type IReleasable interface {
    // operations to be done before closing the associated object.
    // This method is good for releasing expensive resources
    // to prevent memory leak
    Release(optionalParam map[string]interface{}) error
}

// acts as a common interface for logging implementations
type ILogger interface {
    // implements the io.Writer interface; simply to be able to
    // "write" out []byte to the target stream
    Write(p []byte) (n int, err error)

    // return the name of the logger implementation
    Name() string

    // releasable implementation, so that resources could be released
    // when necessary
    Release(optionalParam map[string]interface{}) error
}

// interface indicating the implementation is able to be represented in json format
type IJsonStringAble interface {
    JsonString() string
}
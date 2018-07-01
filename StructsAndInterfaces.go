package queutil

// acts as a common interface for "releasable" object(s)
type IReleasable interface {
    // operations to be done before closing the associated object.
    // This method is good for releasing expensive resources
    // to prevent memory leak
    Release(optionalParam map[string]interface{}) error
}

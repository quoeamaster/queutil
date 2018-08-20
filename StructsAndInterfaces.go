package queutil

import "encoding/json"

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

// ***********************************
// **   Value Object definition(s)  **
// **   sharable containers for     **
// **   common data sharing between **
// **   api calls or processes      **
// ***********************************


// BrokerSeed related information
type BrokerSeedVO struct {
    CanJoin bool                        `json:"CanJoin"`
    IsMasterReady bool                  `json:"IsMasterReady"`
    IsDataReady bool                    `json:"IsDataReady"`

    // is the broker an active master already (no election needed)
    IsActiveMaster bool                 `json:"IsActiveMaster"`

    ClusterName string                  `json:"ClusterName"`
    BrokerName string                   `json:"BrokerName"`
    BrokerCommunicationAddr string      `json:"BrokerCommunicationAddr"`
    BrokerId string                     `json:"BrokerId"`

    DiscoveryBrokerSeeds []string       `json:"DiscoveryBrokerSeeds"`

    SecurityScheme string               `json:"SecurityScheme"`
}
// method to provide a json-fy string based on the contents of the current instance
func (b *BrokerSeedVO) JsonString() string {
    buf, err := json.Marshal(b)
    if err != nil {
        panic(err)
    }
    return string(buf)
}
func (b *BrokerSeedVO) String() string {
    return b.JsonString()
}


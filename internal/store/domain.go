package store

import (
	"time"
    "net/netip"
)

type Request struct {
	Id string
    //Endpoint_id string
    Method string
    Path string
    Query string
    Headers map[string][]string
    Body []byte
    Body_size_bytes int
    Body_truncated bool
    Content_type string
    Source_ip netip.Addr
    Received_at time.Time
}

type Endpoint struct {
    
}

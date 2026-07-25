package store

import (
	"net/netip"
	"time"
)

type Request struct {
	Id              string              `json:"id"`
	Method          *string             `json:"method"`
	Path            *string             `json:"path"`
	Query           *string             `json:"query"`
	Headers         map[string][]string `json:"headers"`
	Body            []byte              `json:"body"`
	Body_size_bytes int                 `json:"body_size_bytes"`
	Body_truncated  bool                `json:"body_truncated"`
	Content_type    string              `json:"content_type"`
	Source_ip       netip.Addr          `json:"source_ip"`
	Received_at     time.Time           `json:"received_at"`
}

type Endpoint struct {
}

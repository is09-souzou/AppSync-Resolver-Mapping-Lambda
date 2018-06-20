package define

import (
	"encoding/json"
)

// Identity Payload Identity struct
type Identity struct {
	Sub                 string          `json:"sub"`
	Issuer              string          `json:"issuer"`
	UserName            string          `json:"username"`
	Claims              json.RawMessage `json:"claims"`
	SourceIP            []string        `json:"sourceIp"`
	DefaultAuthStrategy string          `json:"defaultAuthStrategy"`
}

// Payload request Payload struct
type Payload struct {
	Field     string          `json:"field"`
	Arguments json.RawMessage `json:"arguments"`
	Identity  Identity        `json:"identity"`
}

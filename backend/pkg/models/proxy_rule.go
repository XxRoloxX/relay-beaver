package models

// import (
// 	UUID "github.com/google/uuid"
// )

type ProxyRule struct {
	Id               string  `bson:"_id,omitempty" json:"id,omitempty"`
	Destination      Address `bson:"destination"`
	Targets          []Address
	LoadBalancer     LoadBalancer          `bson:"load_balancer"`
	RequestModifiers []RequestModification `bson:"request_modifiers"`
	// ExternalHooks    []ExternalHook
}

type ExternalHook interface{}

type RequestModification struct {
	Type         string
	Modification interface{}
}

type HeaderModification struct {
	ModifiedHeaders map[string]string
}

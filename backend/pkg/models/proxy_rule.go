package models

// import (
// 	UUID "github.com/google/uuid"
// )

type ProxyRule struct {
	Destination Address `bson:"destination"`
	Targets     []Address
	// LoadBalancer     LoadBalancer
	// RequestModifiers []RequestModification
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

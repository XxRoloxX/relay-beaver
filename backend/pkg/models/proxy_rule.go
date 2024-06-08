package models

type ProxyRule struct {
	Id           string       `bson:"_id,omitempty" json:"id,omitempty"`
	Host         string       `bson:"host" json:"host"`
	Targets      []Address    `bson:"targets" json:"targets"`
	Headers      []Header     `bson:"headers "json:"headers"`
	LoadBalancer LoadBalancer `bson:"load_balancer "json:"load_balancer"`
	// RequestModifiers []RequestModification `bson:"request_modifiers"`
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

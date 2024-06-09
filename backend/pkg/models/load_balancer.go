package models

type LoadBalancer struct {
	Id     string            `bson:"_id,omitempty" json:"id,omitempty"`
	Name   string            `bson:"name"`
	Params map[string]string `bson:"params"`
	//Balancer interface{}
}

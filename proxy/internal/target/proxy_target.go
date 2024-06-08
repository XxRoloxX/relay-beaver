package target

import "strconv"

type HostAddress struct {
	Host string `json:"host" bson:"host"`
	Port int    `json:"port" bson:"port"`
}

func (p *HostAddress) GetURL() string {
	return p.Host + ":" + strconv.Itoa(p.Port)
}

func NewHostAddress(host string, port int) *HostAddress {
	return &HostAddress{
		Host: host,
		Port: port,
	}
}

package target

import (
	"strconv"
	"strings"
)

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
func HostAddressFromString(host string) (HostAddress, error) {
	hostParts := strings.Split(host, ":")

	if len(hostParts) < 2 {
		return HostAddress{
			Host: host,
			Port: 80,
		}, nil
	}

	hostname := strings.Join(hostParts[0:len(hostParts)-1], "")
	port, err := strconv.Atoi(hostParts[len(hostParts)-1])

	if err != nil {
		port = 80
	}

	return HostAddress{
		Host: hostname,
		Port: port,
	}, nil
}

package target

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

type LoadBalancer interface {
	NextTarget(targets []HostAddress) (HostAddress, error)
}

type RoundRobinLoadBalancer struct {
	idx int
}

func (r *RoundRobinLoadBalancer) NextTarget(targets []HostAddress) (HostAddress, error) {
	if len(targets) == 0 {
		log.Error().Msg(fmt.Sprintf("Target cannot be selected, no targets to select from: %s", targets))
		return HostAddress{}, fmt.Errorf("no targets for host")
	}
	return targets[r.idx%len(targets)], nil
}

type LoadBalancerConfiguration struct {
	Name   string
	Params map[string]string
}

func LoadBalancerFactory(lb LoadBalancerConfiguration) LoadBalancer {
	switch lb.Name {
	case "round_robin":
		return &RoundRobinLoadBalancer{}
	default:
		return &RoundRobinLoadBalancer{}
	}
}

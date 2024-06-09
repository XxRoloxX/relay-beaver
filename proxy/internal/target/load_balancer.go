package target

type LoadBalancer interface {
	NextTarget(targets []HostAddress) HostAddress
}

type RoundRobinLoadBalancer struct {
	idx int
}

func (r *RoundRobinLoadBalancer) NextTarget(targets []HostAddress) HostAddress {
	return targets[r.idx%len(targets)]
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

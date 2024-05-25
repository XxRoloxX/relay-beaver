package target

type LoadBalancer interface {
	NextTarget(targets []ProxyTarget) ProxyTarget
}

type RoundRobinLoadBalancer struct {
	idx int
}

func (r *RoundRobinLoadBalancer) NextTarget(targets []ProxyTarget) ProxyTarget {
	return targets[r.idx%len(targets)]
}

package fetch

type RoundRobinLoadBalancer struct {
	indexMap map[string]int
}

func NewRoundRobinLoadBalancer() RoundRobinLoadBalancer {
	return RoundRobinLoadBalancer{make(map[string]int)}
}

func (rrlb RoundRobinLoadBalancer) Next(s Service) Address {
	addr := s.Addresses[rrlb.indexMap[s.Name]]

	if (rrlb.indexMap[s.Name] + 1) >= len(s.Addresses) {
		rrlb.indexMap[s.Name] = 0
	} else {
		rrlb.indexMap[s.Name]++
	}

	return addr
}

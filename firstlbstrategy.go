package fetch

type FirstLoadBalancer struct {
}

func (flb FirstLoadBalancer) Next(s Service) Address {
	return s.Addresses[0]
}

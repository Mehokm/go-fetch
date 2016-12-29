package fetch

type firstLoadBalancer struct {
}

func (flb firstLoadBalancer) Next(s Service) Address {
	return s.Addresses[0]
}

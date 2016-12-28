package fetch

import (
	"math/rand"
	"time"
)

type RandomLoadBalancer struct {
	rander *rand.Rand
}

func NewRandomLoadBalancer() RandomLoadBalancer {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	return RandomLoadBalancer{rand}
}

func (rlb RandomLoadBalancer) Next(s Service) Address {
	return s.Addresses[rlb.rander.Intn(len(s.Addresses))]
}

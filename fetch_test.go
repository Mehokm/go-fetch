package fetch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type stubAdapter struct{}

func (sa stubAdapter) GetService(s string) (Service, error) {
	svc := Service{
		Name: s,
		Addresses: []Address{
			Address{"localhost" + s, "10001"},
			Address{"localhost2" + s, "10002"},
		},
	}

	return svc, nil
}

type stubLbStrategy struct{}

func (slb stubLbStrategy) Next(s Service) Address {
	return s.Addresses[len(s.Addresses)-1]
}

func TestFetchGetHost(t *testing.T) {
	Init(stubAdapter{})

	host, _ := GetHost("ONE")

	assert.Equal(t, "localhostONE", host)
}

func TestFetchGetPort(t *testing.T) {
	Init(stubAdapter{})

	port, _ := GetPort("TWO")

	assert.Equal(t, "10001", port)
}

func TestFetchGetHostAndPort(t *testing.T) {
	Init(stubAdapter{})

	hp, _ := GetHostAndPort("THREE")

	assert.Equal(t, "localhostTHREE:10001", hp)
}

func TestFetchUseStrategy(t *testing.T) {
	Init(stubAdapter{})
	Use(stubLbStrategy{})

	hp, _ := GetHostAndPort("FOUR")

	assert.Equal(t, "localhost2FOUR:10002", hp)
}

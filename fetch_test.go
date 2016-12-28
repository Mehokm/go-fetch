package fetch

import (
	"fmt"
	"testing"
)

type StubAdapter struct{}

func (sa StubAdapter) GetService(s string) (Service, error) {
	svc := Service{
		Name: s,
		Addresses: []Address{
			Address{"localhost", "10001"},
		},
	}

	return svc, nil
}

func TestFetchGetHost(t *testing.T) {
	Init(StubAdapter{})

	host, err := GetHostAndPort("ABCDEF")

	fmt.Println(host, err)
}

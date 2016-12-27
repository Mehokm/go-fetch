package fetch

import (
	"fmt"
	"testing"
)

type StubAdapter struct{}

func (s StubAdapter) GetService(svc string) (Service, error) {
	return Service{svc, "localhost" + svc, "10001"}, nil
}

func TestFetchGetHost(t *testing.T) {
	Init(StubAdapter{})

	host, err := GetHostAndPort("ABCDEF")

	fmt.Println(host, err)
}

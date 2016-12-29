package fetch

import "fmt"

var adapter Adapter
var strategy LbStrategy

type Adapter interface {
	GetService(string) (Service, error)
}

type LbStrategy interface {
	Next(Service) Address
}

type Service struct {
	Name      string
	Addresses []Address
}

type Address struct {
	Host, Port string
}

func init() {
	strategy = firstLoadBalancer{}
}

func Init(a Adapter) {
	adapter = a
}

func Use(lb LbStrategy) {
	strategy = lb
}

func GetHost(s string) (string, error) {
	svc, err := adapter.GetService(s)

	addr := strategy.Next(svc)

	return addr.Host, err
}

func GetPort(s string) (string, error) {
	svc, err := adapter.GetService(s)

	addr := strategy.Next(svc)

	return addr.Port, err
}

func GetHostAndPort(s string) (string, error) {
	svc, err := adapter.GetService(s)

	addr := strategy.Next(svc)

	return fmt.Sprintf("%v:%v", addr.Host, addr.Port), err
}

package fetch

import "fmt"

var adapter Adapter

type Adapter interface {
	GetService(string) (Service, error)
}

type Service struct {
	Name, Host, Port string
}

func Init(a Adapter) {
	adapter = a
}

func GetHost(s string) (string, error) {
	svc, err := adapter.GetService(s)
	return svc.Host, err
}

func GetPort(s string) (string, error) {
	svc, err := adapter.GetService(s)
	return svc.Port, err
}

func GetHostAndPort(s string) (string, error) {
	svc, err := adapter.GetService(s)
	return fmt.Sprintf("%v:%v", svc.Host, svc.Port), err
}

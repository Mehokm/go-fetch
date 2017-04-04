package fetch

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	envReg = regexp.MustCompile("SVC_[A-Z0-9_]+_ADDR")
	hosts  = make(map[string][]string)
	ports  = make(map[string][]string)
)

const (
	stringSep   = "_"
	stringDash  = "-"
	stringColon = ":"
)

type EnvAdpater struct {
	hosts map[string][]string
	ports map[string][]string
}

func loadEnvVariables() {
	envs := os.Environ()

	for _, env := range envs {
		eqIndex := strings.Index(env, "=")

		key := env[:eqIndex]
		value := env[eqIndex+1:]

		if envReg.MatchString(key) {
			name := parseServiceName(key)

			hps := strings.Fields(value)

			for _, hp := range hps {
				host, port := parseHostAndPort(hp)

				hosts[name] = append(hosts[name], host)
				ports[name] = append(ports[name], port)
			}
		}
	}
}

func parseServiceName(s string) string {
	return strings.ToLower(strings.Replace(s[strings.Index(s, stringSep)+1:strings.LastIndex(s, stringSep)], stringSep, stringDash, -1))
}

func parseHostAndPort(s string) (string, string) {
	commaIndex := strings.Index(s, stringColon)

	if commaIndex < 0 {
		return s, "80"
	}

	return s[:commaIndex], s[commaIndex+1:]
}

func NewEnvAdapter() EnvAdpater {
	loadEnvVariables()

	return EnvAdpater{hosts, ports}
}

func (ea EnvAdpater) GetService(s string) (Service, error) {
	svc := Service{}
	var err error

	hosts, ok1 := ea.hosts[s]
	ports, ok2 := ea.ports[s]

	if ok1 && ok2 {
		svc.Name = s

		var addrs []Address
		for i := 0; i < len(hosts) && i < len(ports); i++ {
			addrs = append(addrs, Address{Host: hosts[i], Port: ports[i]})
		}

		svc.Addresses = addrs
	} else {
		err = fmt.Errorf("envadapter: cannot find service by name '%s'", s)
	}

	return svc, err
}

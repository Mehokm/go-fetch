package fetch

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

var (
	envReg = regexp.MustCompile("SVC_[A-Z0-9_]+_ADDR")
)

const (
	stringSep   = "_"
	stringDash  = "-"
	stringColon = ":"
)

type EnvAdpater struct {
	svcMap map[string]Service
}

func NewEnvAdapter() EnvAdpater {
	return EnvAdpater{parseEnvVariables()}
}

func (ea EnvAdpater) GetService(s string) (Service, error) {
	if svc, ok := ea.svcMap[s]; ok {
		return svc, nil
	}

	return Service{}, fmt.Errorf("envadapter: cannot find service by name '%s'", s)
}

func parseEnvVariables() map[string]Service {
	svcMap := make(map[string]Service)

	envs := os.Environ()
	for _, env := range envs {
		eqIndex := strings.Index(env, "=")

		key := env[:eqIndex]
		value := env[eqIndex+1:]

		if envReg.MatchString(key) {
			hps := strings.Split(removeWhitespace(value), ",")
			name := parseServiceName(key)

			svc := Service{}
			svc.Name = name
			svc.Addresses = make([]Address, len(hps))

			for i := 0; i < len(hps); i++ {
				host, port := parseHostAndPort(hps[i])

				svc.Addresses[i].Host = host
				svc.Addresses[i].Port = port
			}
			svcMap[name] = svc
		}
	}

	return svcMap
}

func parseServiceName(s string) string {
	return strings.ToLower(strings.Replace(s[strings.Index(s, stringSep)+1:strings.LastIndex(s, stringSep)], stringSep, stringDash, -1))
}

func parseHostAndPort(s string) (string, string) {
	commaIndex := strings.Index(s, stringColon)

	if commaIndex < 0 {
		return s, "80"
	} else if commaIndex == 0 && len(s) == 1 {
		return "localhost", "80"
	} else if commaIndex == 0 {
		return "localhost", s[commaIndex+1:]
	} else if commaIndex == len(s)-1 {
		return s[:commaIndex], "80"
	}

	return s[:commaIndex], s[commaIndex+1:]
}

func removeWhitespace(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

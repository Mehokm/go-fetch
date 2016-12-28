package fetch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstLoadBalancer(t *testing.T) {
	svc := Service{
		Name: "test_svc",
		Addresses: []Address{
			Address{"host1", "8080"},
			Address{"host2", "8081"},
			Address{"host3", "8082"},
		},
	}

	flb := FirstLoadBalancer{}

	addr := flb.Next(svc)

	assert.Equal(t, addr.Host, "host1")
	assert.Equal(t, addr.Port, "8080")
}

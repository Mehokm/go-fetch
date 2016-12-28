package fetch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomLoadBalancer(t *testing.T) {
	svc := Service{
		Name: "test_svc",
		Addresses: []Address{
			Address{"host1", "8080"},
			Address{"host2", "8081"},
			Address{"host3", "8082"},
		},
	}

	rlb := NewRandomLoadBalancer()

	addr1 := rlb.Next(svc)
	addr2 := rlb.Next(svc)
	addr3 := rlb.Next(svc)

	assert.NotNil(t, addr1)
	assert.NotNil(t, addr2)
	assert.NotNil(t, addr3)
}

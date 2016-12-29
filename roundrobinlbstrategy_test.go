package fetch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoundRobinLoadBalancer(t *testing.T) {
	svc1 := Service{
		Name: "test_svc",
		Addresses: []Address{
			Address{"host1", "8080"},
			Address{"host2", "8081"},
			Address{"host3", "8082"},
		},
	}

	svc2 := Service{
		Name: "test_svc2",
		Addresses: []Address{
			Address{"host1", "9080"},
			Address{"host2", "9081"},
			Address{"host3", "9082"},
		},
	}

	rrlb := NewRoundRobinLoadBalancer()

	addr1 := rrlb.Next(svc1)
	addr2 := rrlb.Next(svc1)
	addr3 := rrlb.Next(svc2)
	addr4 := rrlb.Next(svc1)
	addr5 := rrlb.Next(svc2)
	addr6 := rrlb.Next(svc1)

	assert.Equal(t, addr1, Address{"host1", "8080"})
	assert.Equal(t, addr2, Address{"host2", "8081"})
	assert.Equal(t, addr3, Address{"host1", "9080"})
	assert.Equal(t, addr4, Address{"host3", "8082"})
	assert.Equal(t, addr5, Address{"host2", "9081"})
	assert.Equal(t, addr6, Address{"host1", "8080"})
}

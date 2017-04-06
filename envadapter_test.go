package fetch

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvAdapter_CanReadFromEnv(t *testing.T) {
	// given
	os.Setenv("SVC_TEST_SVC_ONE_ADDR", "test123:8080")

	// when
	svc, _ := NewEnvAdapter().GetService("test-svc-one")

	// then
	assert.Equal(t, "test-svc-one", svc.Name)
	assert.Equal(t, []Address{Address{Host: "test123", Port: "8080"}}, svc.Addresses)

	os.Unsetenv("SVC_TEST_SVC_ONE_ADDR")
}

func TestEnvAdapter_CanReadFromEnv_MultipleAddresses(t *testing.T) {
	// given
	os.Setenv("SVC_TEST_SVC_ONE_ADDR", "test123:8080, test123:8081, test321:9090")

	// when
	svc, _ := NewEnvAdapter().GetService("test-svc-one")

	// then
	assert.Equal(t, "test-svc-one", svc.Name)
	assert.Equal(t, []Address{
		Address{Host: "test123", Port: "8080"},
		Address{Host: "test123", Port: "8081"},
		Address{Host: "test321", Port: "9090"}},
		svc.Addresses,
	)

	os.Unsetenv("SVC_TEST_SVC_ONE_ADDR")
}

func TestEnvAdapter_CanReadFromEnv_TrailingComma(t *testing.T) {
	// given
	os.Setenv("SVC_TEST_SVC_ONE_ADDR", "test123:")

	// when
	svc, _ := NewEnvAdapter().GetService("test-svc-one")

	// then
	assert.Equal(t, "test-svc-one", svc.Name)
	assert.Equal(t, []Address{Address{Host: "test123", Port: "80"}}, svc.Addresses)

	os.Unsetenv("SVC_TEST_SVC_ONE_ADDR")
}

func TestEnvAdapter_CanReadFromEnv_LeadingComma(t *testing.T) {
	// given
	os.Setenv("SVC_TEST_SVC_ONE_ADDR", ":8080")

	// when
	svc, _ := NewEnvAdapter().GetService("test-svc-one")

	// then
	assert.Equal(t, "test-svc-one", svc.Name)
	assert.Equal(t, []Address{Address{Host: "localhost", Port: "8080"}}, svc.Addresses)

	os.Unsetenv("SVC_TEST_SVC_ONE_ADDR")
}

func TestEnvAdapter_CanReadFromEnv_OnlyComma(t *testing.T) {
	// given
	os.Setenv("SVC_TEST_SVC_ONE_ADDR", ":")

	// when
	svc, _ := NewEnvAdapter().GetService("test-svc-one")

	// then
	assert.Equal(t, "test-svc-one", svc.Name)
	assert.Equal(t, []Address{Address{Host: "localhost", Port: "80"}}, svc.Addresses)

	os.Unsetenv("SVC_TEST_SVC_ONE_ADDR")
}

func TestEnvAdapter_CanReadFromEnv_MultipleSvcEnvs(t *testing.T) {
	// given
	os.Setenv("SVC_TEST_SVC_ONE_ADDR", "test567:8080")
	os.Setenv("SVC_TEST_SVC_TWO_ADDR", "test890:9090")

	// when
	ea := NewEnvAdapter()

	svc1, _ := ea.GetService("test-svc-one")
	svc2, _ := ea.GetService("test-svc-two")

	// then
	assert.Equal(t, "test-svc-one", svc1.Name)
	assert.Equal(t, []Address{Address{Host: "test567", Port: "8080"}}, svc1.Addresses)

	assert.Equal(t, "test-svc-two", svc2.Name)
	assert.Equal(t, []Address{Address{Host: "test890", Port: "9090"}}, svc2.Addresses)

	os.Unsetenv("SVC_TEST_SVC_ONE_ADDR")
	os.Unsetenv("SVC_TEST_SVC_TWO_ADDR")
}

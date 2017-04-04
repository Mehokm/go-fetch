package fetch

import (
	"fmt"
	"os"
	"testing"
)

func TestEnvAdapter_CanReadFromEnv(t *testing.T) {
	// given
	os.Setenv("SVC_TEST_SVC_ONE_ADDR", "test123:8080 test123:8081 test321:9090 test321")

	loadEnvVariables()

	ea := NewEnvAdapter()

	// when
	svc, err := ea.GetService("test-svc-one")

	// then

	fmt.Println(svc, err)
	// assert.Equal(t, data, fa.services)

	os.Unsetenv("SVC_TEST_SVC_ONE_ADDR")
}

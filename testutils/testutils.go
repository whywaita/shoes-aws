package testutils

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/localstack"
)

var (
	testEndpoint = ""
)

// IntegrationTestRunner is all integration test
func IntegrationTestRunner(m *testing.M) int {
	ctx := context.Background()

	container, err := localstack.RunContainer(ctx,
		testcontainers.WithImage("localstack/localstack:1.4.0"),
		testcontainers.WithEnv(map[string]string{
			"SERVICES": "ec2",
		}),
	)
	if err != nil {
		log.Fatalf("localstack.StartContainer(): %s", err)
	}

	provider, err := testcontainers.NewDockerProvider()
	if err != nil {
		log.Fatalf("testcontainers.NewDockerProvider(): %s", err)
	}

	host, err := provider.DaemonHost(ctx)
	if err != nil {
		log.Fatalf("provider.DaemonHost(): %s", err)
	}

	mappedPort, err := container.MappedPort(ctx, nat.Port("4566/tcp"))
	if err != nil {
		log.Fatalf("container.MappedPort(): %s", err)
	}

	testEndpoint = fmt.Sprintf("http://%s:%d", host, mappedPort.Int())

	testingEnv := map[string]string{
		"AWS_RESOURCE_TYPE_MAPPING": `{"nano": "c5a.large", "micro": "c5a.xlarge"}`,
	}

	// rewrite to T.Setenv after Go 1.17
	prevEnv := map[string]string{}
	for k, v := range testingEnv {
		prev := os.Getenv(k)
		if err := os.Setenv(k, v); err != nil {
			log.Fatalf("Could not set environment value (key: %q): %+v", k, err)
		}
		prevEnv[k] = prev
	}

	code := m.Run()

	if err := container.Terminate(ctx); err != nil {
		log.Fatalf("failed to terminate container: %s", err)
	}

	return code
}

// GetTestEndpoint get endpoint for test
func GetTestEndpoint() string {
	if testEndpoint == "" {
		panic("testEndpoint is blank, not initialized yet")
	}

	return testEndpoint
}

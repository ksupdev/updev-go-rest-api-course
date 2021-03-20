// +build e2e

package test

import (
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
)

func TestHealthEndpoint(t *testing.T) {
	fmt.Println(" Running E2E test for health check endpoint")

	client := resty.New()
	resp, err := client.R().Get("http://localhost:8088/api/health")

	if err != nil {
		t.Fail()
	}

	fmt.Println(resp.StatusCode())

}

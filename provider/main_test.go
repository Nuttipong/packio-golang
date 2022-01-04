package main

import (
	// "errors"
	// "fmt"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"testing"
)

func TestServerPact_BrokerVerification(t *testing.T) {
	pact := dsl.Pact{
		Provider: "demo-server",
	}

	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		BrokerURL: "http://localhost:9292/",
		// BrokerToken:     "<API TOKEN>",
		ProviderBaseURL: "http://127.0.0.1:5000",
		ProviderVersion: "1.0.0",
		ConsumerVersionSelectors: []types.ConsumerVersionSelector{
			{
				Consumer: "demo-client",
				Tag:      "1.0.0",
			},
		},
		PublishVerificationResults: true, // publish results of verification to PACT broker
	})

	if err != nil {
		t.Log(err)
	}
}

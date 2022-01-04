package main

import (
	"errors"
	"fmt"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"testing"
)

func TestConsumer(t *testing.T) {

	// initialize PACT DSL
	pact := dsl.Pact{
		Consumer: "demo-client",
		Provider: "demo-server",
	}

	// setup a PACT Mock Server
	pact.Setup(true)

	t.Run("Get Product By Id", func(t *testing.T) {
		productID := 35

		pact.
			AddInteraction().
			Given("Product id 35 already exists.").
			UponReceiving("Product 'id' is requested").
			WithRequest(dsl.Request{
				Method: "GET",
				Path:   dsl.Term("/api/products/35", "/api/products/[0-9]+"),
			}).
			WillRespondWith(dsl.Response{
				Status: 200,
				Body: dsl.Like(Product{
					ProductID: productID,
					Sku:       "r2Y361mak",
				}),
			})

		// verify interaction on client side
		err := pact.Verify(func() error {
			// specify host and post of PACT Mock Server as actual server
			host := fmt.Sprintf("http://%s:%d", pact.Host, pact.Server.Port)

			fmt.Println(host)

			// execute function
			product, err := CallGet(host, productID)
			if err != nil {
				return errors.New("error is not expected")
			}

			// check if actual user is equal to expected
			if product == nil || product.ProductID != productID {
				return fmt.Errorf("expected product with ID %d but got %v", productID, product)
			}

			return err
		})

		if err != nil {
			t.Fatal(err)
		}
	})

	// write Contract into file
	if err := pact.WritePact(); err != nil {
		t.Fatal(err)
	}

	publisher := dsl.Publisher{}
	err := publisher.Publish(types.PublishRequest{
		PactURLs:   []string{"./pacts/"},
		PactBroker: "http://localhost:9292", // PACT broker  http://10.198.105.17:9292
		// BrokerToken:     "<API TOKEN>",   // API token for PACT broker
		ConsumerVersion: "1.0.0",
		Tags:            []string{"1.0.0", "latest"},
	})
	if err != nil {
		t.Fatal(err)
	}

	// stop PACT mock server
	pact.Teardown()
}

func TestServerPact_Verification(t *testing.T) {
	// initialize PACT DSL
	pact := dsl.Pact{
		Provider: "demo-server",
	}

	// verify Contract on server side
	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL: "http://127.0.0.1:5000",
		PactURLs:        []string{"../consumer/pacts/demo-client-demo-server.json"},
	})

	if err != nil {
		t.Log(err)
	}
}

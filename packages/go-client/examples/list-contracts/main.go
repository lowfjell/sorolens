// Command list-contracts prints every contract tracked by a Sorolens API,
// following next_cursor until the last page.
//
//	SOROLENS_API_KEY=sl_... go run ./examples/list-contracts
//
// SOROLENS_API_URL defaults to https://api.sorolens.xyz. Set it to
// http://localhost:8080 to use a local API started with `make api`.
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	client "github.com/sorolens/sorolens/packages/go-client"
)

func main() {
	server := os.Getenv("SOROLENS_API_URL")
	if server == "" {
		server = "https://api.sorolens.xyz"
	}

	// The key is sent as "Authorization: Bearer <key>" on every request.
	// Leave SOROLENS_API_KEY unset to make anonymous requests.
	c, err := client.New(server, os.Getenv("SOROLENS_API_KEY"), &http.Client{Timeout: 30 * time.Second})
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	limit := client.Limit(100)
	params := &client.ListContractsParams{Limit: &limit}

	total := 0
	for {
		resp, err := c.ListContracts(ctx, params)
		if err != nil {
			log.Fatal(err) // transport error: DNS, timeout, connection refused
		}
		if resp.JSON200 == nil {
			// Non-2xx: the status line and raw body describe the API error.
			log.Fatalf("list contracts: %s: %s", resp.Status(), resp.Body)
		}

		for _, contract := range resp.JSON200.Contracts {
			fmt.Printf("%s\t%s\t%s\n", contract.Id, contract.Network, contract.Status)
		}
		total += len(resp.JSON200.Contracts)

		// An empty next_cursor means this was the last page.
		if resp.JSON200.NextCursor == "" {
			break
		}
		next := resp.JSON200.NextCursor
		params.Cursor = &next
	}
	fmt.Printf("%d contracts\n", total)
}

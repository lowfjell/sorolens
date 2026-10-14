package client_test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"

	client "github.com/sorolens/sorolens/packages/go-client"
)

// pagedServer serves three contracts, two per page, and rejects requests
// that do not carry the expected bearer token.
func pagedServer() *httptest.Server {
	ids := []string{"CONTRACT_A", "CONTRACT_B", "CONTRACT_C"}
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Header.Get("Authorization") != "Bearer sl_example" {
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(map[string]any{"error": map[string]string{"code": "UNAUTHORIZED", "message": "missing key"}})
			return
		}
		start := 0
		if r.URL.Query().Get("cursor") == "page2" {
			start = 2
		}
		end, next := start+2, "page2"
		if end >= len(ids) {
			end, next = len(ids), ""
		}
		var contracts []map[string]any
		for _, id := range ids[start:end] {
			contracts = append(contracts, map[string]any{
				"id": id, "network": "testnet", "status": "active", "added_at": "2026-09-01T00:00:00Z",
			})
		}
		_ = json.NewEncoder(w).Encode(map[string]any{"contracts": contracts, "next_cursor": next})
	}))
}

// Page through every contract by feeding next_cursor back as the cursor until
// the API returns an empty one.
func ExampleApiClient_ListContracts_pagination() {
	server := pagedServer()
	defer server.Close()

	c, err := client.New(server.URL, "sl_example", nil)
	if err != nil {
		log.Fatal(err)
	}

	limit := client.Limit(2)
	params := &client.ListContractsParams{Limit: &limit}
	for {
		resp, err := c.ListContracts(context.Background(), params)
		if err != nil {
			log.Fatal(err)
		}
		if resp.JSON200 == nil {
			log.Fatalf("list contracts: %s", resp.Status())
		}
		for _, contract := range resp.JSON200.Contracts {
			fmt.Println(contract.Id, contract.Status == client.ContractStatusActive)
		}
		if resp.JSON200.NextCursor == "" {
			break
		}
		next := client.Cursor(resp.JSON200.NextCursor)
		params.Cursor = &next
	}
	// Output:
	// CONTRACT_A true
	// CONTRACT_B true
	// CONTRACT_C true
}

// Without an API key the client sends no Authorization header, so endpoints
// that require a key answer with a non-2xx status and JSON200 stays nil.
func ExampleNew_anonymous() {
	server := pagedServer()
	defer server.Close()

	c, err := client.New(server.URL, "", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := c.ListContracts(context.Background(), &client.ListContractsParams{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.StatusCode(), resp.JSON200 == nil)
	// Output:
	// 401 true
}

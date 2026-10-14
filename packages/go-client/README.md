# go-client

Go client for the Sorolens API.

The HTTP client is generated from [`docs/openapi.yaml`](../../docs/openapi.yaml)
with `oapi-codegen` into `internal/client`. The package at the module root wraps
it with an API-key-aware constructor and typed methods, and re-exports the
request and response types those methods use, so you never import generated
code directly.

## Install

The module lives in a subdirectory of the monorepo, so its path includes it:

```sh
go get github.com/sorolens/sorolens/packages/go-client@latest
```

Releases are tagged `packages/go-client/vX.Y.Z` (the Go convention for a module
in a subdirectory); `@latest` picks the newest such tag. To track unreleased
changes, use `@main` instead.

The package name is `client`. The import path ends in `go-client`, which is not
a valid identifier, so name the import explicitly:

```go
import client "github.com/sorolens/sorolens/packages/go-client"
```

Requires Go 1.26 or newer (see `go.mod`).

## Authentication

Pass an API key to `client.New`. It is sent as `Authorization: Bearer <key>`
on every request:

```go
c, err := client.New("https://api.sorolens.xyz", os.Getenv("SOROLENS_API_KEY"), nil)
if err != nil {
	log.Fatal(err)
}
```

With an empty key the client sends no `Authorization` header. Public endpoints
answer anonymous requests; endpoints that need a key return `401`.

The third argument is an optional `*http.Client`. Pass one to set a timeout or a
custom transport; `nil` uses a plain `&http.Client{}`, which has no timeout:

```go
c, err := client.New(server, apiKey, &http.Client{Timeout: 30 * time.Second})
```

## Making requests

Each method returns the generated response type. `err` is non-nil only for
transport failures (DNS, timeouts, refused connections). An HTTP error status is
not a Go error: check `JSON200`, or the field for the status you expect, and
fall back to `StatusCode()`, `Status()`, and `Body`:

```go
resp, err := c.GetContract(ctx, "CA...")
if err != nil {
	return err
}
switch {
case resp.JSON200 != nil:
	fmt.Println(resp.JSON200.Label, resp.JSON200.Status)
case resp.StatusCode() == http.StatusNotFound:
	fmt.Println("contract is not tracked")
default:
	return fmt.Errorf("get contract: %s: %s", resp.Status(), resp.Body)
}
```

The typed methods are `ListContracts`, `GetContract`, `GetGlobalStats`,
`GetContractForecast`, `ListContractEvents`, `ListContractInvocations`,
`ListContractStorage`, `ListWatchdogAlerts` and `GetContractSnapshot`. For any
other endpoint, `c.Generated()` returns the full generated client. Its methods
follow the OpenAPI spec and may change when the client is regenerated; the
methods on `ApiClient` are the stable surface.

## Pagination

List endpoints return a page of results and a `NextCursor`. Pass it back as
`Cursor` to get the next page; an empty `NextCursor` means you have the last
page. `Limit` sets the page size; for contracts the API defaults to 50 and
caps it at 200.

```go
limit := client.Limit(100)
params := &client.ListContractsParams{Limit: &limit}
for {
	resp, err := c.ListContracts(ctx, params)
	if err != nil {
		return err
	}
	if resp.JSON200 == nil {
		return fmt.Errorf("list contracts: %s: %s", resp.Status(), resp.Body)
	}
	for _, contract := range resp.JSON200.Contracts {
		fmt.Println(contract.Id, contract.Status)
	}
	if resp.JSON200.NextCursor == "" {
		break
	}
	next := resp.JSON200.NextCursor
	params.Cursor = &next
}
```

`ListContractEvents`, `ListContractInvocations`, `ListContractStorage` and
`ListWatchdogAlerts` paginate the same way.

## Example

[`examples/list-contracts`](examples/list-contracts/main.go) is a complete
program that pages through every tracked contract:

```sh
cd packages/go-client
SOROLENS_API_KEY=sl_your_api_key go run ./examples/list-contracts
# against a local API started with `make api`:
SOROLENS_API_URL=http://localhost:8080 go run ./examples/list-contracts
```

CI (`go-client-external`) builds this example from a separate module outside
the repo, the same way a project that ran `go get` would.

## Regenerating

```sh
make client-go
```

The target builds `oapi-codegen` with the pinned Go 1.26 toolchain
(`GOTOOLCHAIN=go1.26.0`). `oapi-codegen` embeds a gzipped copy of the spec in
`client.gen.go`, and gzip output differs between Go releases, so regenerating
with another toolchain produces a diff that CI rejects even when the client is
otherwise in sync.

CI (`go-client-stale`) fails on any PR where the checked-in generated code
differs from the spec, so the client is always in sync with `docs/openapi.yaml`.

If a regeneration changes or removes a type re-exported in `types.go`, the
module stops compiling; update `types.go` to match.

## Releasing

Push a tag with the module's directory as a prefix:

```sh
git tag packages/go-client/v0.1.0
git push origin packages/go-client/v0.1.0
```

The Go module proxy picks it up on the first `go get`; nothing else is
published.

## Layout

- `client.go` — public wrapper: `client.New`, `ApiClient` and its typed methods.
- `types.go` — re-exports of the generated types used by `ApiClient`.
- `examples/list-contracts/` — runnable example (auth and pagination).
- `internal/client/client.gen.go` — generated types and HTTP client (do not edit).

## Testing

```sh
cd packages/go-client && go test -race ./...
```

The tests run against an `httptest` stub of the API. `example_test.go` holds
runnable examples (pagination and anonymous access) that `go test` executes and
`go doc` shows.

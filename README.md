# bs-euro

bs-euro prices European options with the Black-Scholes model. Given spot,
strike, time to expiry, risk-free rate, volatility, and a call or put flag,
it returns the price plus d1 and d2. A second endpoint returns the first-order
greeks: delta, gamma, vega, theta, and rho.

The kernel uses the standard error function for the normal CDF, so option
values and Greeks are computed with full library precision rather than a
coarse quadrature. The implementation keeps put-call parity closed with the
q=0 convention: C-P = S - K*e^{-rT}.

## Run a bundled example

```bash
go run . price --table example/atm-1y.json
```

The ATM one-year call uses S=K=100, r=3%, sigma=20%. Its price is near
0.4*S*sigma*sqrt(T).

## CLI

```bash
go run . price example/atm-1y.json
go run . greeks example/atm-1y.json
```

Invalid input is written to standard error and exits nonzero.

## HTTP service

The default command starts a service on port 8080:

```bash
go run . serve --http :8080
```

Endpoints:

```text
POST /api/price
POST /api/greeks
GET  /health
```

Example request:

```bash
curl -s -X POST http://127.0.0.1:8080/api/price \
  -H 'Content-Type: application/json' \
  -d '{"s":100,"k":100,"t":1,"r":0.03,"sigma":0.2,"flag":"call"}'
```

Errors are returned as JSON with a machine-readable `code` and a human
message, for example `invalid_spot`, `invalid_volatility`, or
`invalid_flag`.

## Build and test

```bash
go build ./...
go test ./...
```

The pricing kernel lives under `internal/`; `main.go` only wires the CLI and
HTTP entry points.

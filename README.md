# Bhole - ConnectRPC / REST Gateway & SDK for LaunchLibrary v2

Bhole is a high-performance, type-safe, and declarative ConnectRPC / REST integration layer and SDK for the LaunchLibrary v2 (LL2) API. It features a modular **Go-kit style architecture** for both client and server layers, along with a unified **TypeScript Web Client** for frontend integration.

---

## 🏗️ Architecture & Component Layout

The codebase separates concerns between pure business domains, generic transportation logic, and wire serialization:

```
├── client/                 # Client Packages
│   ├── {feature}/          # Feature Client (e.g. client/launch, client/agency)
│   │   ├── endpoint.go     # Declares the business service endpoints struct
│   │   └── connectrpc_transport.go # Declares ConnectRPC transport constructors & mappings
│   ├── ts/                 # TypeScript client package (Connect-ES v2.x)
│   └── transport/          # REST-mapping ConnectRPC HTTP client (RESTClient)
│
├── {feature}/              # Feature Package at the module root (e.g. launch/, agency/)
│   ├── {feature}.go        # Domain structs (request/response and resource models)
│   ├── service.go          # Service interface implemented by clients and servers alike
│   ├── endpoint.go         # Go-kit server endpoints
│   └── connectrpc_server.go # ConnectRPC service handler, encoders/decoders & mappers
│
├── cmd/                    # Binaries
│   ├── example/            # Go client demonstration calling the public LL2 REST API
│   └── server/             # ConnectRPC gateway hosting all 18 services, backed by LL2 REST
│
├── web/                    # Vite demo web app using the TypeScript client
├── proto/                  # Protobuf Schemas (Buf module)
└── scripts/                # Code generation engines
```

Domain models and `Service` interfaces live in top-level packages (`github.com/pobochiigo/bhole/launch`, `.../agency`, ...) rather than under `internal/`, so other Go modules can import the client SDK and construct requests.

---

## 🛠️ Code Generation Workflow

Due to the massive size of the LL2 API (18 primary resources, 50+ schemas, hundreds of fields), both Go client/server architectures are fully automated via generation engines. Every generator formats its output with `gofmt`, so the tree stays format-clean after regeneration.

### 1. Generate Go Client SDK
Generates protos, domain models, service interfaces, client endpoint files and the REST-mapping transport:
```bash
python3 scripts/generate_client.py
```

### 2. Generate Go Server Handlers
Generates endpoint bindings, encoders, decoders, and business-to-proto mappers inside each `{feature}/` package:
```bash
python3 scripts/generate_server.py
```

### 3. Generate API Server Command
Generates `cmd/server/main.go` registering all services on top of the REST-backed clients:
```bash
python3 scripts/generate_server_cmd.py
```

### 4. Compile Protobuf (Go & TypeScript)
Uses `buf` to compile schemas into Go client/server bindings and TypeScript definitions:
```bash
buf generate
```

Optional HTML reference docs for the schemas (requires `protoc-gen-doc` on `PATH`, output is gitignored):
```bash
buf generate --template buf.gen.docs.yaml
```

---

## 🚀 Getting Started

### Run the Go Client Example
The Go client uses the `RESTClient` in `client/transport`, which intercepts ConnectRPC requests and maps them directly to the public REST API, requiring **no local running server**:
```bash
go run ./cmd/example
```

The REST client propagates the caller's `context.Context` (cancellation, deadlines, Connect timeouts) and forwards an `Authorization` header when one is set on the Connect request, so LL2 API tokens work unchanged. Upstream HTTP errors surface as `*connect.Error` values with matching codes (`404` → `CodeNotFound`, `429` → `CodeResourceExhausted`, and so on).

### Run the ConnectRPC API Gateway
Spins up a local HTTP/1.1 + cleartext HTTP/2 (h2c) server hosting all 18 ConnectRPC service handlers with CORS support. Each handler is backed by the REST-mapping client, so responses contain real LL2 data:
```bash
go run ./cmd/server
```
The server listens at `http://localhost:8080` by default and shuts down gracefully on `SIGINT`/`SIGTERM`.

| Variable       | Default                            | Purpose                       |
|----------------|------------------------------------|-------------------------------|
| `LL2_BASE_URL` | `https://lldev.thespacedevs.com`   | Upstream LL2 REST base URL    |
| `BHOLE_ADDR`   | `:8080`                            | Listen address                |

### Use the Go SDK from another module
```go
import (
    "github.com/pobochiigo/bhole/client/transport"
    launchclient "github.com/pobochiigo/bhole/client/launch"
    "github.com/pobochiigo/bhole/launch"
)

rest := transport.NewRESTClient("https://lldev.thespacedevs.com", nil)
svc := launchclient.NewLaunchClient(rest, "https://lldev.thespacedevs.com")
resp, err := svc.ListLaunches(ctx, &launch.ListLaunchesRequest{Limit: 5})
```

### Build the TypeScript Web Client
Compile the TypeScript/ES ConnectRPC client package:
```bash
cd client/ts
npm install
npm run build
```
Build assets, typing definitions, and source maps will be generated in `client/ts/dist/`.

---

## 🧪 Running Tests & Checks
```bash
go test ./...
gofmt -l .            # must print nothing
golangci-lint run     # config in .golangci.yml
```

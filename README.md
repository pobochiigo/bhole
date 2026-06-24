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
│   └── transport/          # REST Interceptor & HTTP transport helper
│
├── internal/               # Server & Business Logic
│   ├── {feature}/          # Feature Business Logic (e.g. internal/launch)
│   │   ├── endpoint.go     # Declares Go-kit server endpoints
│   │   └── connectrpc_server.go # ConnectRPC service handler, encoders/decoders & mappers
│   └── transport/          # Generic server-side transport handlers
│
├── cmd/                    # Binaries
│   ├── example/            # Go client demonstration calling public LL2 REST API
│   └── server/             # API gateway server hosting all 18 ConnectRPC handlers
│
├── proto/                  # Protobuf Schemas (Buf module)
└── scripts/                # Code generation engines
```

---

## 🛠️ Code Generation Workflow

Due to the massive size of the LL2 API (18 primary resources, 50+ schemas, hundreds of fields), both Go client/server architectures are fully automated via generation engines:

### 1. Generate Go Client SDK
Generates type-safe client endpoint files and legacy file cleanup:
```bash
python3 scripts/generate_client.py
```

### 2. Generate Go Server Handlers
Generates endpoint bindings, encoders, decoders, and business-to-proto mappers inside `internal/{feature}`:
```bash
python3 scripts/generate_server.py
```

### 3. Generate API Server Command
Generates `cmd/server/main.go` registering all services and stubs:
```bash
python3 scripts/generate_server_cmd.py
```

### 4. Compile Protobuf (Go & TypeScript)
Uses `buf` to compile schemas into Go client/server bindings and TypeScript definitions:
```bash
buf generate
```

---

## 🚀 Getting Started

### Run the Go Client Example
The Go client utilizes a custom [RESTClient](file:///Users/28soft/GitLab/pobochiigo/bhole/client/transport/rest_client.go) that intercepts ConnectRPC requests and maps them directly to the public REST API, requiring **no local running server**:
```bash
go run cmd/example/main.go
```

### Run the ConnectRPC API Server
Spins up a local HTTP/2 cleartext (h2c) server hosting all 18 ConnectRPC service handlers with built-in CORS support:
```bash
go run cmd/server/main.go
```
The server will listen at `http://localhost:8080`.

### Build the TypeScript Web Client
Compile the compiled TypeScript/ES ConnectRPC client package:
```bash
cd client/ts
npm install
npm run build
```
Build assets, typing definitions, and source maps will be generated in `client/ts/dist/`.

---

## 🧪 Running Tests
All mock-based client/server integration tests can be run using the standard Go test command:
```bash
go test ./...
```

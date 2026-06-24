# Bhole TypeScript ConnectRPC Client

This package contains the TypeScript/ES ConnectRPC client generated from the project's Protobuf definitions.

## 🚀 How to Generate & Build the TS Client

If you modify the Protobuf files (`.proto`) and need to regenerate the TypeScript client, follow these steps:

### 1. Install Node Dependencies
Ensure you have installed the code generator plugin and dependencies:
```bash
cd client/ts
npm install
```

### 2. Generate the Client Code
From the **root directory** of the project, run the `buf` code generator:
```bash
# Run from the repository root
buf generate
```
This runs the local `@bufbuild/protoc-gen-es` plugin and outputs the generated files into `client/ts/src/gen`.

### 3. Build the Package
Compile the TypeScript code to JavaScript and type definitions:
```bash
cd client/ts
npm run build
```
This outputs compiled JavaScript files, declaration maps, and typing files to `client/ts/dist/`.

---

## 📦 Usage in a Web App

You can integrate this client into any frontend web application (React, Vue, Svelte, Next.js, etc.).

### 1. Initialize Transport & Client
```typescript
import { createClient, createConnectTransport, Launch } from "bhole-client";

// Create the ConnectRPC web transport pointing to your API gateway/server
const transport = createConnectTransport({
  baseUrl: "http://localhost:8080",
});

// Instantiate the service client
const client = createClient(Launch.LaunchService, transport);

// Call an endpoint
const response = await client.listLaunches({ limit: 10 });
console.log("Total Launches:", response.count);
```

### 2. Service Namespaces
All 18 service definitions and schemas are grouped under distinct namespaces to prevent naming clashes:
* `Agency`
* `Astronaut`
* `CelestialBody`
* `DockingEvent`
* `Event`
* `Expedition`
* `Landing`
* `Launcher`
* `LauncherConfiguration`
* `Launch`
* `Location`
* `Pad`
* `Payload`
* `Program`
* `SpaceStation`
* `Spacecraft`
* `Spacewalk`
* `Update`

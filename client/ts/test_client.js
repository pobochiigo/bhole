import { createClient, createConnectTransport, Launch } from "./dist/index.js";

console.log("ConnectRPC TypeScript client loaded successfully!");
console.log("LaunchService available:", !!Launch.LaunchService);
console.log("Launch schema available:", !!Launch.LaunchSchema);

// Setup a transport and client
const transport = createConnectTransport({
  baseUrl: "https://lldev.thespacedevs.com",
});

const client = createClient(Launch.LaunchService, transport);
console.log("Client successfully created:", !!client);
console.log("ListLaunches method available:", typeof client.listLaunches === "function");

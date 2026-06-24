import { createClient, createConnectTransport, Launch } from "./dist/index.js";

const transport = createConnectTransport({
  baseUrl: "http://localhost:8080",
});

const client = createClient(Launch.LaunchService, transport);

console.log("Sending ListLaunches request to local Go server on http://localhost:8080...");
try {
  const resp = await client.listLaunches({});
  console.log("Response received successfully!");
  console.log("Count:", resp.count);
  console.log("Results length:", resp.results.length);
  process.exit(0);
} catch (err) {
  console.error("Error calling server:", err);
  process.exit(1);
}

import { createConnectTransport } from "@connectrpc/connect-web";
import { createClient, Client } from "@connectrpc/connect";

// Re-export transport and client creation helpers
export { createConnectTransport, createClient };
export type { Client };

// Export generated services and messages grouped by namespaces to avoid conflicts
export * as Agency from "./gen/proto/agency/v1/agency_pb.js";
export * as Astronaut from "./gen/proto/astronaut/v1/astronaut_pb.js";
export * as CelestialBody from "./gen/proto/celestial_body/v1/celestial_body_pb.js";
export * as DockingEvent from "./gen/proto/docking_event/v1/docking_event_pb.js";
export * as Event from "./gen/proto/event/v1/event_pb.js";
export * as Expedition from "./gen/proto/expedition/v1/expedition_pb.js";
export * as Landing from "./gen/proto/landing/v1/landing_pb.js";
export * as Launcher from "./gen/proto/launcher/v1/launcher_pb.js";
export * as LauncherConfiguration from "./gen/proto/launcher_configuration/v1/launcher_configuration_pb.js";
export * as Launch from "./gen/proto/launch/v1/launch_pb.js";
export * as Location from "./gen/proto/location/v1/location_pb.js";
export * as Pad from "./gen/proto/pad/v1/pad_pb.js";
export * as Payload from "./gen/proto/payload/v1/payload_pb.js";
export * as Program from "./gen/proto/program/v1/program_pb.js";
export * as SpaceStation from "./gen/proto/space_station/v1/space_station_pb.js";
export * as Spacecraft from "./gen/proto/spacecraft/v1/spacecraft_pb.js";
export * as Spacewalk from "./gen/proto/spacewalk/v1/spacewalk_pb.js";
export * as Update from "./gen/proto/update/v1/update_pb.js";

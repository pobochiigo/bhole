package main

import (
	"context"
"fmt"
"log"
"net/http"
"golang.org/x/net/http2"
"golang.org/x/net/http2/h2c"
agency "com.gitlab/pobochiigo/bhole/internal/agency"
agencyv1connect "com.gitlab/pobochiigo/bhole/proto/agency/v1/agencyv1connect"
astronaut "com.gitlab/pobochiigo/bhole/internal/astronaut"
astronautv1connect "com.gitlab/pobochiigo/bhole/proto/astronaut/v1/astronautv1connect"
celestial_body "com.gitlab/pobochiigo/bhole/internal/celestial_body"
celestial_bodyv1connect "com.gitlab/pobochiigo/bhole/proto/celestial_body/v1/celestial_bodyv1connect"
docking_event "com.gitlab/pobochiigo/bhole/internal/docking_event"
docking_eventv1connect "com.gitlab/pobochiigo/bhole/proto/docking_event/v1/docking_eventv1connect"
event "com.gitlab/pobochiigo/bhole/internal/event"
eventv1connect "com.gitlab/pobochiigo/bhole/proto/event/v1/eventv1connect"
expedition "com.gitlab/pobochiigo/bhole/internal/expedition"
expeditionv1connect "com.gitlab/pobochiigo/bhole/proto/expedition/v1/expeditionv1connect"
landing "com.gitlab/pobochiigo/bhole/internal/landing"
landingv1connect "com.gitlab/pobochiigo/bhole/proto/landing/v1/landingv1connect"
launch "com.gitlab/pobochiigo/bhole/internal/launch"
launchv1connect "com.gitlab/pobochiigo/bhole/proto/launch/v1/launchv1connect"
launcher "com.gitlab/pobochiigo/bhole/internal/launcher"
launcherv1connect "com.gitlab/pobochiigo/bhole/proto/launcher/v1/launcherv1connect"
launcher_configuration "com.gitlab/pobochiigo/bhole/internal/launcher_configuration"
launcher_configurationv1connect "com.gitlab/pobochiigo/bhole/proto/launcher_configuration/v1/launcher_configurationv1connect"
location "com.gitlab/pobochiigo/bhole/internal/location"
locationv1connect "com.gitlab/pobochiigo/bhole/proto/location/v1/locationv1connect"
pad "com.gitlab/pobochiigo/bhole/internal/pad"
padv1connect "com.gitlab/pobochiigo/bhole/proto/pad/v1/padv1connect"
payload "com.gitlab/pobochiigo/bhole/internal/payload"
payloadv1connect "com.gitlab/pobochiigo/bhole/proto/payload/v1/payloadv1connect"
program "com.gitlab/pobochiigo/bhole/internal/program"
programv1connect "com.gitlab/pobochiigo/bhole/proto/program/v1/programv1connect"
space_station "com.gitlab/pobochiigo/bhole/internal/space_station"
space_stationv1connect "com.gitlab/pobochiigo/bhole/proto/space_station/v1/space_stationv1connect"
spacecraft "com.gitlab/pobochiigo/bhole/internal/spacecraft"
spacecraftv1connect "com.gitlab/pobochiigo/bhole/proto/spacecraft/v1/spacecraftv1connect"
spacewalk "com.gitlab/pobochiigo/bhole/internal/spacewalk"
spacewalkv1connect "com.gitlab/pobochiigo/bhole/proto/spacewalk/v1/spacewalkv1connect"
update "com.gitlab/pobochiigo/bhole/internal/update"
updatev1connect "com.gitlab/pobochiigo/bhole/proto/update/v1/updatev1connect"
)

// CORS Middleware to support browser-based TypeScript clients
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Connect-Protocol-Version, Connect-Timeout")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}

// --- Agency Stub Service ---
type stubAgencyService struct{}

func (s *stubAgencyService) ListAgencies(ctx context.Context, req *agency.ListAgenciesRequest) (*agency.ListAgenciesResponse, error) {
	return &agency.ListAgenciesResponse{
		Count: 1,
		Results: []agency.Agency{},
	}, nil
}

func (s *stubAgencyService) GetAgency(ctx context.Context, req *agency.GetAgencyRequest) (*agency.Agency, error) {
	return &agency.Agency{}, nil
}

// --- Astronaut Stub Service ---
type stubAstronautService struct{}

func (s *stubAstronautService) ListAstronauts(ctx context.Context, req *astronaut.ListAstronautsRequest) (*astronaut.ListAstronautsResponse, error) {
	return &astronaut.ListAstronautsResponse{
		Count: 1,
		Results: []astronaut.Astronaut{},
	}, nil
}

func (s *stubAstronautService) GetAstronaut(ctx context.Context, req *astronaut.GetAstronautRequest) (*astronaut.Astronaut, error) {
	return &astronaut.Astronaut{}, nil
}

// --- CelestialBody Stub Service ---
type stubCelestialBodyService struct{}

func (s *stubCelestialBodyService) ListCelestialBodies(ctx context.Context, req *celestial_body.ListCelestialBodiesRequest) (*celestial_body.ListCelestialBodiesResponse, error) {
	return &celestial_body.ListCelestialBodiesResponse{
		Count: 1,
		Results: []celestial_body.CelestialBody{},
	}, nil
}

func (s *stubCelestialBodyService) GetCelestialBody(ctx context.Context, req *celestial_body.GetCelestialBodyRequest) (*celestial_body.CelestialBody, error) {
	return &celestial_body.CelestialBody{}, nil
}

// --- DockingEvent Stub Service ---
type stubDockingEventService struct{}

func (s *stubDockingEventService) ListDockingEvents(ctx context.Context, req *docking_event.ListDockingEventsRequest) (*docking_event.ListDockingEventsResponse, error) {
	return &docking_event.ListDockingEventsResponse{
		Count: 1,
		Results: []docking_event.DockingEvent{},
	}, nil
}

func (s *stubDockingEventService) GetDockingEvent(ctx context.Context, req *docking_event.GetDockingEventRequest) (*docking_event.DockingEvent, error) {
	return &docking_event.DockingEvent{}, nil
}

// --- Event Stub Service ---
type stubEventService struct{}

func (s *stubEventService) ListEvents(ctx context.Context, req *event.ListEventsRequest) (*event.ListEventsResponse, error) {
	return &event.ListEventsResponse{
		Count: 1,
		Results: []event.Event{},
	}, nil
}

func (s *stubEventService) GetEvent(ctx context.Context, req *event.GetEventRequest) (*event.Event, error) {
	return &event.Event{}, nil
}

// --- Expedition Stub Service ---
type stubExpeditionService struct{}

func (s *stubExpeditionService) ListExpeditions(ctx context.Context, req *expedition.ListExpeditionsRequest) (*expedition.ListExpeditionsResponse, error) {
	return &expedition.ListExpeditionsResponse{
		Count: 1,
		Results: []expedition.Expedition{},
	}, nil
}

func (s *stubExpeditionService) GetExpedition(ctx context.Context, req *expedition.GetExpeditionRequest) (*expedition.Expedition, error) {
	return &expedition.Expedition{}, nil
}

// --- Landing Stub Service ---
type stubLandingService struct{}

func (s *stubLandingService) ListLandings(ctx context.Context, req *landing.ListLandingsRequest) (*landing.ListLandingsResponse, error) {
	return &landing.ListLandingsResponse{
		Count: 1,
		Results: []landing.Landing{},
	}, nil
}

func (s *stubLandingService) GetLanding(ctx context.Context, req *landing.GetLandingRequest) (*landing.Landing, error) {
	return &landing.Landing{}, nil
}

// --- Launch Stub Service ---
type stubLaunchService struct{}

func (s *stubLaunchService) ListLaunches(ctx context.Context, req *launch.ListLaunchesRequest) (*launch.ListLaunchesResponse, error) {
	return &launch.ListLaunchesResponse{
		Count: 1,
		Results: []launch.Launch{},
	}, nil
}

func (s *stubLaunchService) GetLaunch(ctx context.Context, req *launch.GetLaunchRequest) (*launch.Launch, error) {
	return &launch.Launch{}, nil
}

// --- Launcher Stub Service ---
type stubLauncherService struct{}

func (s *stubLauncherService) ListLaunchers(ctx context.Context, req *launcher.ListLaunchersRequest) (*launcher.ListLaunchersResponse, error) {
	return &launcher.ListLaunchersResponse{
		Count: 1,
		Results: []launcher.Launcher{},
	}, nil
}

func (s *stubLauncherService) GetLauncher(ctx context.Context, req *launcher.GetLauncherRequest) (*launcher.Launcher, error) {
	return &launcher.Launcher{}, nil
}

// --- LauncherConfiguration Stub Service ---
type stubLauncherConfigurationService struct{}

func (s *stubLauncherConfigurationService) ListLauncherConfigurations(ctx context.Context, req *launcher_configuration.ListLauncherConfigurationsRequest) (*launcher_configuration.ListLauncherConfigurationsResponse, error) {
	return &launcher_configuration.ListLauncherConfigurationsResponse{
		Count: 1,
		Results: []launcher_configuration.LauncherConfiguration{},
	}, nil
}

func (s *stubLauncherConfigurationService) GetLauncherConfiguration(ctx context.Context, req *launcher_configuration.GetLauncherConfigurationRequest) (*launcher_configuration.LauncherConfiguration, error) {
	return &launcher_configuration.LauncherConfiguration{}, nil
}

// --- Location Stub Service ---
type stubLocationService struct{}

func (s *stubLocationService) ListLocations(ctx context.Context, req *location.ListLocationsRequest) (*location.ListLocationsResponse, error) {
	return &location.ListLocationsResponse{
		Count: 1,
		Results: []location.Location{},
	}, nil
}

func (s *stubLocationService) GetLocation(ctx context.Context, req *location.GetLocationRequest) (*location.Location, error) {
	return &location.Location{}, nil
}

// --- Pad Stub Service ---
type stubPadService struct{}

func (s *stubPadService) ListPads(ctx context.Context, req *pad.ListPadsRequest) (*pad.ListPadsResponse, error) {
	return &pad.ListPadsResponse{
		Count: 1,
		Results: []pad.Pad{},
	}, nil
}

func (s *stubPadService) GetPad(ctx context.Context, req *pad.GetPadRequest) (*pad.Pad, error) {
	return &pad.Pad{}, nil
}

// --- Payload Stub Service ---
type stubPayloadService struct{}

func (s *stubPayloadService) ListPayloads(ctx context.Context, req *payload.ListPayloadsRequest) (*payload.ListPayloadsResponse, error) {
	return &payload.ListPayloadsResponse{
		Count: 1,
		Results: []payload.Payload{},
	}, nil
}

func (s *stubPayloadService) GetPayload(ctx context.Context, req *payload.GetPayloadRequest) (*payload.Payload, error) {
	return &payload.Payload{}, nil
}

// --- Program Stub Service ---
type stubProgramService struct{}

func (s *stubProgramService) ListPrograms(ctx context.Context, req *program.ListProgramsRequest) (*program.ListProgramsResponse, error) {
	return &program.ListProgramsResponse{
		Count: 1,
		Results: []program.Program{},
	}, nil
}

func (s *stubProgramService) GetProgram(ctx context.Context, req *program.GetProgramRequest) (*program.Program, error) {
	return &program.Program{}, nil
}

// --- SpaceStation Stub Service ---
type stubSpaceStationService struct{}

func (s *stubSpaceStationService) ListSpaceStations(ctx context.Context, req *space_station.ListSpaceStationsRequest) (*space_station.ListSpaceStationsResponse, error) {
	return &space_station.ListSpaceStationsResponse{
		Count: 1,
		Results: []space_station.SpaceStation{},
	}, nil
}

func (s *stubSpaceStationService) GetSpaceStation(ctx context.Context, req *space_station.GetSpaceStationRequest) (*space_station.SpaceStation, error) {
	return &space_station.SpaceStation{}, nil
}

// --- Spacecraft Stub Service ---
type stubSpacecraftService struct{}

func (s *stubSpacecraftService) ListSpacecrafts(ctx context.Context, req *spacecraft.ListSpacecraftsRequest) (*spacecraft.ListSpacecraftsResponse, error) {
	return &spacecraft.ListSpacecraftsResponse{
		Count: 1,
		Results: []spacecraft.Spacecraft{},
	}, nil
}

func (s *stubSpacecraftService) GetSpacecraft(ctx context.Context, req *spacecraft.GetSpacecraftRequest) (*spacecraft.Spacecraft, error) {
	return &spacecraft.Spacecraft{}, nil
}

// --- Spacewalk Stub Service ---
type stubSpacewalkService struct{}

func (s *stubSpacewalkService) ListSpacewalks(ctx context.Context, req *spacewalk.ListSpacewalksRequest) (*spacewalk.ListSpacewalksResponse, error) {
	return &spacewalk.ListSpacewalksResponse{
		Count: 1,
		Results: []spacewalk.Spacewalk{},
	}, nil
}

func (s *stubSpacewalkService) GetSpacewalk(ctx context.Context, req *spacewalk.GetSpacewalkRequest) (*spacewalk.Spacewalk, error) {
	return &spacewalk.Spacewalk{}, nil
}

// --- Update Stub Service ---
type stubUpdateService struct{}

func (s *stubUpdateService) ListUpdates(ctx context.Context, req *update.ListUpdatesRequest) (*update.ListUpdatesResponse, error) {
	return &update.ListUpdatesResponse{
		Count: 1,
		Results: []update.Update{},
	}, nil
}

func (s *stubUpdateService) GetUpdate(ctx context.Context, req *update.GetUpdateRequest) (*update.Update, error) {
	return &update.Update{}, nil
}

func main() {
	mux := http.NewServeMux()

	fmt.Println("Registering ConnectRPC service handlers...")
	mux.Handle(agencyv1connect.NewAgencyServiceHandler(agency.NewAgencyHandler(&stubAgencyService{})))
	mux.Handle(astronautv1connect.NewAstronautServiceHandler(astronaut.NewAstronautHandler(&stubAstronautService{})))
	mux.Handle(celestial_bodyv1connect.NewCelestialBodyServiceHandler(celestial_body.NewCelestialBodyHandler(&stubCelestialBodyService{})))
	mux.Handle(docking_eventv1connect.NewDockingEventServiceHandler(docking_event.NewDockingEventHandler(&stubDockingEventService{})))
	mux.Handle(eventv1connect.NewEventServiceHandler(event.NewEventHandler(&stubEventService{})))
	mux.Handle(expeditionv1connect.NewExpeditionServiceHandler(expedition.NewExpeditionHandler(&stubExpeditionService{})))
	mux.Handle(landingv1connect.NewLandingServiceHandler(landing.NewLandingHandler(&stubLandingService{})))
	mux.Handle(launchv1connect.NewLaunchServiceHandler(launch.NewLaunchHandler(&stubLaunchService{})))
	mux.Handle(launcherv1connect.NewLauncherServiceHandler(launcher.NewLauncherHandler(&stubLauncherService{})))
	mux.Handle(launcher_configurationv1connect.NewLauncherConfigurationServiceHandler(launcher_configuration.NewLauncherConfigurationHandler(&stubLauncherConfigurationService{})))
	mux.Handle(locationv1connect.NewLocationServiceHandler(location.NewLocationHandler(&stubLocationService{})))
	mux.Handle(padv1connect.NewPadServiceHandler(pad.NewPadHandler(&stubPadService{})))
	mux.Handle(payloadv1connect.NewPayloadServiceHandler(payload.NewPayloadHandler(&stubPayloadService{})))
	mux.Handle(programv1connect.NewProgramServiceHandler(program.NewProgramHandler(&stubProgramService{})))
	mux.Handle(space_stationv1connect.NewSpaceStationServiceHandler(space_station.NewSpaceStationHandler(&stubSpaceStationService{})))
	mux.Handle(spacecraftv1connect.NewSpacecraftServiceHandler(spacecraft.NewSpacecraftHandler(&stubSpacecraftService{})))
	mux.Handle(spacewalkv1connect.NewSpacewalkServiceHandler(spacewalk.NewSpacewalkHandler(&stubSpacewalkService{})))
	mux.Handle(updatev1connect.NewUpdateServiceHandler(update.NewUpdateHandler(&stubUpdateService{})))

	// Wrap mux with CORS middleware
	handler := corsMiddleware(mux)

	// Enable HTTP/2 h2c (HTTP/2 cleartext) for ConnectRPC
	h2Handler := h2c.NewHandler(handler, &http2.Server{})

	port := ":8080"
	fmt.Printf("ConnectRPC API server starting on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, h2Handler); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

import json
import os

# Mapping of the 18 primary LL2 API resources
RESOURCES = {
    "agency": {"plural_method": "ListAgencies", "singular": "Agency"},
    "astronaut": {"plural_method": "ListAstronauts", "singular": "Astronaut"},
    "celestial_body": {"plural_method": "ListCelestialBodies", "singular": "CelestialBody"},
    "docking_event": {"plural_method": "ListDockingEvents", "singular": "DockingEvent"},
    "event": {"plural_method": "ListEvents", "singular": "Event"},
    "expedition": {"plural_method": "ListExpeditions", "singular": "Expedition"},
    "landing": {"plural_method": "ListLandings", "singular": "Landing"},
    "launcher": {"plural_method": "ListLaunchers", "singular": "Launcher"},
    "launcher_configuration": {"plural_method": "ListLauncherConfigurations", "singular": "LauncherConfiguration"},
    "launch": {"plural_method": "ListLaunches", "singular": "Launch"},
    "location": {"plural_method": "ListLocations", "singular": "Location"},
    "pad": {"plural_method": "ListPads", "singular": "Pad"},
    "payload": {"plural_method": "ListPayloads", "singular": "Payload"},
    "program": {"plural_method": "ListPrograms", "singular": "Program"},
    "space_station": {"plural_method": "ListSpaceStations", "singular": "SpaceStation"},
    "spacecraft": {"plural_method": "ListSpacecrafts", "singular": "Spacecraft"},
    "spacewalk": {"plural_method": "ListSpacewalks", "singular": "Spacewalk"},
    "update": {"plural_method": "ListUpdates", "singular": "Update"}
}

def camel_case(s):
    return "".join(x.capitalize() for x in s.split("_"))

def main():
    print("Generating cmd/server/main.go...")
    os.makedirs("cmd/server", exist_ok=True)
    
    # Imports
    imports = [
        '"context"',
        '"fmt"',
        '"log"',
        '"net/http"',
    ]
    for feature in sorted(RESOURCES.keys()):
        imports.append(f'{feature} "github.com/pobochiigo/bhole/internal/{feature}"')
        imports.append(f'{feature}v1connect "github.com/pobochiigo/bhole/proto/{feature}/v1/{feature}v1connect"')
        
    main_content = f"""package main

import (
\t{chr(10).join(imports)}
)

// CORS Middleware to support browser-based TypeScript clients
func corsMiddleware(next http.Handler) http.Handler {{
\treturn http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {{
\t\tw.Header().Set("Access-Control-Allow-Origin", "*")
\t\tw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
\t\tw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Connect-Protocol-Version, Connect-Timeout")
\t\t
\t\tif r.Method == "OPTIONS" {{
\t\t\tw.WriteHeader(http.StatusOK)
\t\t\treturn
\t\t}}
\t\t
\t\tnext.ServeHTTP(w, r)
\t}})
}}
"""

    # Generate Dummy Services
    for feature, config in sorted(RESOURCES.items()):
        feature_camel = camel_case(feature)
        plural_method = config["plural_method"]
        
        main_content += f"""
// --- {feature_camel} Stub Service ---
type stub{feature_camel}Service struct{{}}

func (s *stub{feature_camel}Service) {plural_method}(ctx context.Context, req *{feature}.{plural_method}Request) (*{feature}.{plural_method}Response, error) {{
\treturn &{feature}.{plural_method}Response{{
\t\tCount: 1,
\t\tResults: []{feature}.{feature_camel}{{}},
\t}}, nil
}}

func (s *stub{feature_camel}Service) Get{feature_camel}(ctx context.Context, req *{feature}.Get{feature_camel}Request) (*{feature}.{feature_camel}, error) {{
\treturn &{feature}.{feature_camel}{{}}, nil
}}
"""

    # Generate Main function
    main_content += """
func main() {
	mux := http.NewServeMux()

	fmt.Println("Registering ConnectRPC service handlers...")
"""

    for feature, config in sorted(RESOURCES.items()):
        feature_camel = camel_case(feature)
        main_content += f"\tmux.Handle({feature}v1connect.New{feature_camel}ServiceHandler({feature}.New{feature_camel}Handler(&stub{feature_camel}Service{{}})))\n"

    main_content += """
	// Wrap mux with CORS middleware
	handler := corsMiddleware(mux)

	port := ":8080"
	fmt.Printf("ConnectRPC API server starting on http://localhost%s\\n", port)

	srv := &http.Server{
		Addr:    port,
		Handler: handler,
	}
	srv.Protocols = new(http.Protocols)
	srv.Protocols.SetHTTP1(true)
	srv.Protocols.SetUnencryptedHTTP2(true)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}
}
"""

    with open("cmd/server/main.go", "w") as f:
        f.write(main_content)
        
    print("cmd/server/main.go generated successfully!")

if __name__ == "__main__":
    main()

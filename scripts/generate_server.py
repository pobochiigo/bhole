import json
import os
import re

# Mapping of the 18 primary LL2 API resources to their path and configuration
RESOURCES = {
    "agency": {
        "path": "/2.3.0/agencies/", 
        "plural_method": "ListAgencies",
        "singular": "Agency",
        "id_type": "int32",
        "parse_id": "int(protoReq.Id)"
    },
    "astronaut": {
        "path": "/2.3.0/astronauts/", 
        "plural_method": "ListAstronauts",
        "singular": "Astronaut",
        "id_type": "int32",
        "parse_id": "int(protoReq.Id)"
    },
    "celestial_body": {
        "path": "/2.3.0/celestial_bodies/", 
        "plural_method": "ListCelestialBodies",
        "singular": "CelestialBody",
        "id_type": "int32",
        "parse_id": "int(protoReq.Id)"
    },
    "docking_event": {
        "path": "/2.3.0/docking_events/", 
        "plural_method": "ListDockingEvents",
        "singular": "DockingEvent",
        "id_type": "int32",
        "parse_id": "int(protoReq.Id)"
    },
    "event": {
        "path": "/2.3.0/events/", 
        "plural_method": "ListEvents",
        "singular": "Event",
        "id_type": "int32",
        "parse_id": "int(protoReq.Id)"
    },
    "expedition": {
        "path": "/2.3.0/expeditions/", 
        "plural_method": "ListExpeditions",
        "singular": "Expedition",
        "id_type": "int32",
        "parse_id": "int(protoReq.Id)"
    },
    "landing": {
        "path": "/2.3.0/landings/", 
        "plural_method": "ListLandings",
        "singular": "Landing",
        "id_type": "int32",
        "parse_id": "int(protoReq.Id)"
    },
    "launcher": {
        "path": "/2.3.0/launchers/", 
        "plural_method": "ListLaunchers",
        "singular": "Launcher",
        "id_type": "int32",
        "parse_id": "int(protoReq.Id)"
    },
    "launcher_configuration": {
        "path": "/2.3.0/launcher_configurations/", 
        "plural_method": "ListLauncherConfigurations",
        "singular": "LauncherConfiguration",
        "id_type": "int32",
        "parse_id": "int(protoReq.Id)"
    },
    "launch": {
        "path": "/2.3.0/launches/", 
        "plural_method": "ListLaunches",
        "singular": "Launch",
        "id_type": "string",
        "parse_id": "protoReq.Id"
    },
    "location": {
        "path": "/2.3.0/locations/", 
        "plural_method": "ListLocations",
        "singular": "Location",
        "id_type": "int32",
        "parse_id": "int(protoReq.Id)"
    },
    "pad": {
        "path": "/2.3.0/pads/", 
        "plural_method": "ListPads",
        "singular": "Pad",
        "id_type": "int32",
        "parse_id": "int(protoReq.Id)"
    },
    "payload": {
        "path": "/2.3.0/payloads/", 
        "plural_method": "ListPayloads",
        "singular": "Payload",
        "id_type": "int32",
        "parse_id": "int(protoReq.Id)"
    },
    "program": {
        "path": "/2.3.0/programs/", 
        "plural_method": "ListPrograms",
        "singular": "Program",
        "id_type": "int32",
        "parse_id": "int(protoReq.Id)"
    },
    "space_station": {
        "path": "/2.3.0/space_stations/", 
        "plural_method": "ListSpaceStations",
        "singular": "SpaceStation",
        "id_type": "int32",
        "parse_id": "int(protoReq.Id)"
    },
    "spacecraft": {
        "path": "/2.3.0/spacecraft/", 
        "plural_method": "ListSpacecrafts",
        "singular": "Spacecraft",
        "id_type": "int32",
        "parse_id": "int(protoReq.Id)"
    },
    "spacewalk": {
        "path": "/2.3.0/spacewalks/", 
        "plural_method": "ListSpacewalks",
        "singular": "Spacewalk",
        "id_type": "int32",
        "parse_id": "int(protoReq.Id)"
    },
    "update": {
        "path": "/2.3.0/updates/", 
        "plural_method": "ListUpdates",
        "singular": "Update",
        "id_type": "int32",
        "parse_id": "int(protoReq.Id)"
    }
}

# Go keywords that should be escaped if used as field names
GO_KEYWORDS = {"type", "select", "range", "interface", "map", "chan", "go", "struct", "import"}

def camel_case(s):
    return "".join(x.capitalize() for x in s.split("_"))

def go_field_name(name):
    parts = name.split("_")
    field = "".join(p.capitalize() for p in parts)
    if field.lower() in GO_KEYWORDS:
        field += "Val"
    return field

def go_proto_field_name(name):
    if name == "alpha_2_code":
        return "Alpha_2Code"
    if name == "alpha_3_code":
        return "Alpha_3Code"
    parts = name.split("_")
    field = "".join(p.capitalize() for p in parts)
    return field

def get_ref_name(schema):
    if "$ref" in schema:
        return schema["$ref"].split("/")[-1]
    if "allOf" in schema:
        for sub in schema["allOf"]:
            if "$ref" in sub:
                return sub["$ref"].split("/")[-1]
    return None

def resolve_polymorphic_schema(schema_name, schemas):
    if schema_name not in schemas:
        return schema_name
    schema = schemas[schema_name]
    if "oneOf" in schema:
        refs = []
        for x in schema["oneOf"]:
            if "$ref" in x:
                refs.append(x["$ref"].split("/")[-1])
        for r in refs:
            if "Detailed" in r:
                return resolve_polymorphic_schema(r, schemas)
        for r in refs:
            if "Normal" in r:
                return resolve_polymorphic_schema(r, schemas)
        if refs:
            return resolve_polymorphic_schema(refs[-1], schemas)
    return schema_name

def map_schema_name(t_name, main_schema, feature_camel):
    if t_name == main_schema:
        return feature_camel
    if t_name == feature_camel:
        return feature_camel + "Record"
    return t_name

def get_schema_properties(schema_name, schemas):
    if schema_name not in schemas:
        return {}
    schema = schemas[schema_name]
    properties = {}
    if "allOf" in schema:
        for sub in schema["allOf"]:
            if "$ref" in sub:
                ref_name = sub["$ref"].split("/")[-1]
                properties.update(get_schema_properties(ref_name, schemas))
            elif "properties" in sub:
                properties.update(sub["properties"])
    if "properties" in schema:
        properties.update(schema["properties"])
    return properties

def resolve_property_info(prop_name, prop_val, schemas, resource_name):
    ref = get_ref_name(prop_val)
    if ref:
        resolved_ref = resolve_polymorphic_schema(ref, schemas)
        return {"kind": "ref", "type": resolved_ref}
    if prop_val.get("type") == "array" and "items" in prop_val:
        arr_ref = get_ref_name(prop_val["items"])
        if arr_ref:
            resolved_arr_ref = resolve_polymorphic_schema(arr_ref, schemas)
            return {"kind": "array_ref", "type": resolved_arr_ref}
        arr_type = prop_val["items"].get("type", "string")
        return {"kind": "array_primitive", "type": arr_type}
    p_type = prop_val.get("type")
    if p_type in ["string", "integer", "boolean", "number"]:
        return {"kind": "primitive", "type": p_type}
    if p_type == "object" and "properties" in prop_val:
        sub_name = camel_case(resource_name) + camel_case(prop_name)
        schemas[sub_name] = prop_val
        return {"kind": "ref", "type": sub_name}
    return {"kind": "primitive", "type": "string"}

def collect_types(schema_name, schemas, collected):
    if schema_name in collected:
        return
    collected.add(schema_name)
    props = get_schema_properties(schema_name, schemas)
    for p_name, p_val in props.items():
        info = resolve_property_info(p_name, p_val, schemas, schema_name)
        if info["kind"] == "ref":
            collect_types(info["type"], schemas, collected)
        elif info["kind"] == "array_ref":
            collect_types(info["type"], schemas, collected)

def main():
    print("Reading ll2_api.json...")
    with open("ll2_api.json") as f:
        spec = json.load(f)
    schemas = spec["components"]["schemas"]
    paths = spec["paths"]
    
    feature_schemas = {}
    
    for feature, config in RESOURCES.items():
        path_schema = paths[config["path"]]["get"]["responses"]["200"]["content"]["application/json"]["schema"]
        list_schema_name = path_schema["$ref"].split("/")[-1]
        results_items = schemas[list_schema_name]["properties"]["results"]["items"]
        
        main_schema = None
        if "$ref" in results_items:
            main_schema = results_items["$ref"].split("/")[-1]
        elif "oneOf" in results_items:
            refs = [x["$ref"].split("/")[-1] for x in results_items["oneOf"] if "$ref" in x]
            for r in refs:
                if "Detailed" in r:
                    main_schema = r
                    break
            if not main_schema and refs:
                main_schema = refs[-1]
        if not main_schema:
            main_schema = config["singular"]
            
        main_schema = resolve_polymorphic_schema(main_schema, schemas)
        config["main_schema"] = main_schema
        
        feature_types = set()
        collect_types(main_schema, schemas, feature_types)
        feature_schemas[feature] = feature_types

    for feature, config in RESOURCES.items():
        print(f"Generating server files for: {feature}...")
        feature_camel = camel_case(feature)
        main_schema = config["main_schema"]
        feature_types = feature_schemas[feature]
        plural_method = config["plural_method"]
        
        # Make directories
        os.makedirs(f"internal/{feature}", exist_ok=True)
        
        # --- A. Generate endpoint.go ---
        endpoint_content = f"""package {feature}

import (
	"context"

	"com.gitlab/pobochiigo/bhole/client/endpoint"
)

type Endpoints struct {{
	list{plural_method} endpoint.Endpoint[*{plural_method}Request, *{plural_method}Response]
	get{feature_camel}    endpoint.Endpoint[*Get{feature_camel}Request, *{feature_camel}]
}}

func MakeEndpoints(svc Service) Endpoints {{
	return Endpoints{{
		list{plural_method}: makeList{plural_method}Endpoint(svc),
		get{feature_camel}:    makeGet{feature_camel}Endpoint(svc),
	}}
}}

func makeList{plural_method}Endpoint(svc Service) endpoint.Endpoint[*{plural_method}Request, *{plural_method}Response] {{
	return func(ctx context.Context, req *{plural_method}Request) (*{plural_method}Response, error) {{
		return svc.{plural_method}(ctx, req)
	}}
}}

func makeGet{feature_camel}Endpoint(svc Service) endpoint.Endpoint[*Get{feature_camel}Request, *{feature_camel}] {{
	return func(ctx context.Context, req *Get{feature_camel}Request) (*{feature_camel}, error) {{
		return svc.Get{feature_camel}(ctx, req)
	}}
}}
"""
        with open(f"internal/{feature}/endpoint.go", "w") as ef:
            ef.write(endpoint_content)
            
        # --- B. Generate connectrpc_server.go ---
        server_content = f"""package {feature}

import (
	"context"

	"com.gitlab/pobochiigo/bhole/internal/transport"
	{feature}v1 "com.gitlab/pobochiigo/bhole/proto/{feature}/v1"
	v1connect "com.gitlab/pobochiigo/bhole/proto/{feature}/v1/{feature}v1connect"
	"connectrpc.com/connect"
)

type server struct {{
	list{plural_method} transport.Handler[{feature}v1.{plural_method}Request, {feature}v1.{plural_method}Response]
	get{feature_camel}    transport.Handler[{feature}v1.Get{feature_camel}Request, {feature}v1.Get{feature_camel}Response]
}}

func (s *server) {plural_method}(ctx context.Context, req *connect.Request[{feature}v1.{plural_method}Request]) (*connect.Response[{feature}v1.{plural_method}Response], error) {{
	return s.list{plural_method}(ctx, req)
}}

func (s *server) Get{feature_camel}(ctx context.Context, req *connect.Request[{feature}v1.Get{feature_camel}Request]) (*connect.Response[{feature}v1.Get{feature_camel}Response], error) {{
	return s.get{feature_camel}(ctx, req)
}}

func New{feature_camel}Handler(svc Service) v1connect.{feature_camel}ServiceHandler {{
	eps := MakeEndpoints(svc)
	return &server{{
		list{plural_method}: transport.NewConnectServer(
			eps.list{plural_method},
			decode{plural_method}Request,
			encode{plural_method}Response,
		),
		get{feature_camel}: transport.NewConnectServer(
			eps.get{feature_camel},
			decodeGet{feature_camel}Request,
			encodeGet{feature_camel}Response,
		),
	}}
}}

func decode{plural_method}Request(_ context.Context, req *{feature}v1.{plural_method}Request) (*{plural_method}Request, error) {{
	return &{plural_method}Request{{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}}, nil
}}

func encode{plural_method}Response(ctx context.Context, resp *{plural_method}Response) (*{feature}v1.{plural_method}Response, error) {{
	results := make([]*{feature}v1.{feature_camel}, len(resp.Results))
	for i := range resp.Results {{
		results[i] = mapBizToProto{feature_camel}(&resp.Results[i])
	}}
	return &{feature}v1.{plural_method}Response{{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}}, nil
}}

func decodeGet{feature_camel}Request(_ context.Context, req *{feature}v1.Get{feature_camel}Request) (*Get{feature_camel}Request, error) {{
	return &Get{feature_camel}Request{{
		ID:   req.Id,
		Mode: req.Mode,
	}}, nil
}}

func encodeGet{feature_camel}Response(ctx context.Context, resp *{feature_camel}) (*{feature}v1.Get{feature_camel}Response, error) {{
	return &{feature}v1.Get{feature_camel}Response{{
		{feature_camel}: mapBizToProto{feature_camel}(resp),
	}}, nil
}}

"""
        # Generate the biz-to-proto mappings for all dependent types in the feature
        for t_name in sorted(list(feature_types)):
            struct_name = map_schema_name(t_name, main_schema, feature_camel)
            server_content += f"func mapBizToProto{struct_name}(r *{struct_name}) *{feature}v1.{struct_name} {{\n"
            server_content += f"\tif r == nil {{\n\t\treturn nil\n\t}}\n"
            server_content += f"\treturn &{feature}v1.{struct_name}{{\n"
            
            props = get_schema_properties(t_name, schemas)
            for p_name, p_val in sorted(props.items()):
                p_info = resolve_property_info(p_name, p_val, schemas, t_name)
                f_type = map_schema_name(p_info["type"], main_schema, feature_camel)
                    
                g_field = go_field_name(p_name)
                p_proto_go_field = go_proto_field_name(p_name)
                
                if p_info["kind"] == "primitive":
                    server_content += f"\t\t{p_proto_go_field}: r.{g_field},\n"
                elif p_info["kind"] == "array_primitive":
                    server_content += f"\t\t{p_proto_go_field}: r.{g_field},\n"
                elif p_info["kind"] == "ref":
                    server_content += f"\t\t{p_proto_go_field}: mapBizToProto{f_type}(r.{g_field}),\n"
                elif p_info["kind"] == "array_ref":
                    server_content += f"\t\t{p_proto_go_field}: func() []*{feature}v1.{f_type} {{\n"
                    server_content += f"\t\t\tif r.{g_field} == nil {{\n\t\t\t\treturn nil\n\t\t\t}}\n"
                    server_content += f"\t\t\tres := make([]*{feature}v1.{f_type}, len(r.{g_field}))\n"
                    server_content += f"\t\t\tfor i := range r.{g_field} {{\n"
                    server_content += f"\t\t\t\tres[i] = mapBizToProto{f_type}(&r.{g_field}[i])\n"
                    server_content += f"\t\t\t}}\n"
                    server_content += f"\t\t\treturn res\n"
                    server_content += f"\t\t}}(),\n"
                    
            server_content += "\t}\n}\n\n"
            
        with open(f"internal/{feature}/connectrpc_server.go", "w") as sf:
            sf.write(server_content)

    print("Server files generation completed successfully!")

if __name__ == "__main__":
    main()

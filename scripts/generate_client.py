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
    # e.g. celestial_body -> CelestialBody
    return "".join(x.capitalize() for x in s.split("_"))

def snake_case(s):
    # e.g. CelestialBody -> celestial_body
    s1 = re.sub('(.)([A-Z][a-z]+)', r'\1_\2', s)
    return re.sub('([a-z0-9])([A-Z])', r'\1_\2', s1).lower()

def go_field_name(name):
    # e.g. alpha_2_code -> Alpha2Code
    parts = name.split("_")
    field = "".join(p.capitalize() for p in parts)
    # Check if field name clashes with Go keyword
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

def proto_field_name(name):
    # e.g. type -> type_val if it clashes with proto keyword (unlikely but safe)
    clean = name.replace("-", "_").replace(".", "_")
    return clean

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
    
    # Resolve inheritance
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
    # 1. Direct Reference
    ref = get_ref_name(prop_val)
    if ref:
        resolved_ref = resolve_polymorphic_schema(ref, schemas)
        return {"kind": "ref", "type": resolved_ref}
        
    # 2. Array
    if prop_val.get("type") == "array" and "items" in prop_val:
        arr_ref = get_ref_name(prop_val["items"])
        if arr_ref:
            resolved_arr_ref = resolve_polymorphic_schema(arr_ref, schemas)
            return {"kind": "array_ref", "type": resolved_arr_ref}
        arr_type = prop_val["items"].get("type", "string")
        return {"kind": "array_primitive", "type": arr_type}
        
    # 3. Primitive Type
    p_type = prop_val.get("type")
    if p_type in ["string", "integer", "boolean", "number"]:
        return {"kind": "primitive", "type": p_type}
        
    # 4. Inline Object
    if p_type == "object" and "properties" in prop_val:
        sub_name = camel_case(resource_name) + camel_case(prop_name)
        schemas[sub_name] = prop_val
        return {"kind": "ref", "type": sub_name}
        
    # Default fallback to string
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

def map_openapi_to_proto_type(kind, t_name):
    if kind == "primitive":
        if t_name == "integer":
            return "int32"
        if t_name == "number":
            return "float"
        if t_name == "boolean":
            return "bool"
        return "string"
    elif kind == "array_primitive":
        sub = map_openapi_to_proto_type("primitive", t_name)
        return f"repeated {sub}"
    elif kind == "ref":
        return t_name
    elif kind == "array_ref":
        return f"repeated {t_name}"
    return "string"

def map_openapi_to_go_type(kind, t_name):
    if kind == "primitive":
        if t_name == "integer":
            return "int32"
        if t_name == "number":
            return "float32"
        if t_name == "boolean":
            return "bool"
        return "string"
    elif kind == "array_primitive":
        sub = map_openapi_to_go_type("primitive", t_name)
        return f"[]{sub}"
    elif kind == "ref":
        return f"*{t_name}"
    elif kind == "array_ref":
        return f"[]{t_name}"
    return "string"

def map_openapi_to_json_go_type(kind, t_name):
    if kind == "primitive":
        return map_openapi_to_go_type(kind, t_name)
    elif kind == "array_primitive":
        return map_openapi_to_go_type(kind, t_name)
    elif kind == "ref":
        return f"*{t_name}JSON"
    elif kind == "array_ref":
        return f"[]{t_name}JSON"
    return "string"

def main():
    print("Reading ll2_api.json...")
    with open("ll2_api.json") as f:
        spec = json.load(f)
    schemas = spec["components"]["schemas"]
    paths = spec["paths"]
    
    # We will gather all schema details for all 18 features
    feature_schemas = {}
    all_shared_types = set()
    
    for feature, config in RESOURCES.items():
        print(f"Analyzing schema for resource: {feature}...")
        path_schema = paths[config["path"]]["get"]["responses"]["200"]["content"]["application/json"]["schema"]
        list_schema_name = path_schema["$ref"].split("/")[-1]
        
        results_items = schemas[list_schema_name]["properties"]["results"]["items"]
        
        # Resolve the concrete item schema name (Detailed representation)
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
        print(f"  Resolved main schema: {main_schema}")
        config["main_schema"] = main_schema
        
        # Collect all dependent types recursively
        feature_types = set()
        collect_types(main_schema, schemas, feature_types)
        feature_schemas[feature] = feature_types
        all_shared_types.update(feature_types)
        
    print(f"Total collected unique types: {len(all_shared_types)}")
    
    # 1. Generate Proto, Internal Models, Internal Services, and Clients for each feature
    for feature, config in RESOURCES.items():
        print(f"Generating package files for: {feature}...")
        feature_camel = camel_case(feature)
        main_schema = config["main_schema"]
        feature_types = feature_schemas[feature]
        plural_method = config["plural_method"]
        
        # Make directories
        os.makedirs(f"proto/{feature}/v1", exist_ok=True)
        os.makedirs(f"internal/{feature}", exist_ok=True)
        os.makedirs(f"client/{feature}", exist_ok=True)
        
        # --- A. Generate PROTO ---
        proto_content = f"""syntax = "proto3";

package {feature}.v1;

option go_package = "github.com/pobochiigo/bhole/proto/{feature}/v1;{feature}v1";

"""
        # Define messages for all dependent schemas
        for t_name in sorted(list(feature_types)):
            # If it's the main detailed schema, we map it directly to the singular feature name (e.g. LaunchDetailed -> Launch)
            msg_name = map_schema_name(t_name, main_schema, feature_camel)
            proto_content += f"message {msg_name} {{\n"
            
            props = get_schema_properties(t_name, schemas)
            f_num = 1
            for p_name, p_val in sorted(props.items()):
                # If the field is named "id" and it is the main resource, we enforce it matches the path ID type (string/int32)
                p_info = resolve_property_info(p_name, p_val, schemas, t_name)
                
                # Check mapping for dependent refs to their mapped message name
                f_type = map_schema_name(p_info["type"], main_schema, feature_camel)
                
                p_proto_type = map_openapi_to_proto_type(p_info["kind"], f_type)
                
                # Proto 3 optional helper
                optional_prefix = ""
                if p_val.get("nullable") and p_info["kind"] == "primitive":
                    optional_prefix = "optional "
                    
                proto_content += f"  {optional_prefix}{p_proto_type} {proto_field_name(p_name)} = {f_num};\n"
                f_num += 1
            proto_content += "}\n\n"
            
        # Add Service, Request, and Response messages
        proto_content += f"""message {plural_method}Request {{
  int32 limit = 1;
  int32 offset = 2;
  string search = 3;
  string mode = 4;
}}

message {plural_method}Response {{
  int32 count = 1;
  string next = 2;
  string previous = 3;
  repeated {feature_camel} results = 4;
}}

message Get{feature_camel}Request {{
  {config["id_type"]} id = 1;
  string mode = 2;
}}

message Get{feature_camel}Response {{
  {feature_camel} {feature} = 1;
}}

service {feature_camel}Service {{
  rpc {plural_method}({plural_method}Request) returns ({plural_method}Response);
  rpc Get{feature_camel}(Get{feature_camel}Request) returns (Get{feature_camel}Response);
}}
"""
        with open(f"proto/{feature}/v1/{feature}.proto", "w") as pf:
            pf.write(proto_content)
            
        # --- B. Generate INTERNAL MODELS (BUSINESS MODELS) ---
        go_models_content = f"""package {feature}

"""
        for t_name in sorted(list(feature_types)):
            struct_name = map_schema_name(t_name, main_schema, feature_camel)
            go_models_content += f"type {struct_name} struct {{\n"
            
            props = get_schema_properties(t_name, schemas)
            for p_name, p_val in sorted(props.items()):
                p_info = resolve_property_info(p_name, p_val, schemas, t_name)
                f_type = map_schema_name(p_info["type"], main_schema, feature_camel)
                    
                go_type = map_openapi_to_go_type(p_info["kind"], f_type)
                
                # Handle optional float/int pointer for nullable primitives if desired (here we match previous probability *float32 pattern)
                if p_val.get("nullable") and p_info["kind"] == "primitive":
                    go_type = "*" + go_type
                    
                go_models_content += f"\t{go_field_name(p_name)} {go_type}\n"
            go_models_content += "}\n\n"
            
        go_models_content += f"""type {plural_method}Request struct {{
	Limit  int32
	Offset int32
	Search string
	Mode   string
}}

type {plural_method}Response struct {{
	Count    int32
	Next     string
	Previous string
	Results  []{feature_camel}
}}

type Get{feature_camel}Request struct {{
	ID   {config["id_type"]}
	Mode string
}}
"""
        with open(f"internal/{feature}/{feature}.go", "w") as mf:
            mf.write(go_models_content)
            
        # --- C. Generate INTERNAL SERVICE INTERFACE ---
        go_svc_content = f"""package {feature}

import "context"

type Service interface {{
	{plural_method}(ctx context.Context, req *{plural_method}Request) (*{plural_method}Response, error)
	Get{feature_camel}(ctx context.Context, req *Get{feature_camel}Request) (*{feature_camel}, error)
}}
"""
        with open(f"internal/{feature}/service.go", "w") as sf:
            sf.write(go_svc_content)
            
        # --- D. Generate CONNECTRPC CLIENT (GO-KIT GENERICS STYLE) ---
        # Generate endpoint.go
        endpoint_content = f"""package {feature}

import (
	"context"

	"github.com/pobochiigo/bhole/client/endpoint"
	biz{feature} "github.com/pobochiigo/bhole/internal/{feature}"
)

type endpoints struct {{
	list{plural_method} endpoint.Endpoint[*biz{feature}.{plural_method}Request, *biz{feature}.{plural_method}Response]
	get{feature_camel}    endpoint.Endpoint[*biz{feature}.Get{feature_camel}Request, *biz{feature}.{feature_camel}]
}}

func (c *endpoints) {plural_method}(ctx context.Context, req *biz{feature}.{plural_method}Request) (*biz{feature}.{plural_method}Response, error) {{
	return c.list{plural_method}(ctx, req)
}}

func (c *endpoints) Get{feature_camel}(ctx context.Context, req *biz{feature}.Get{feature_camel}Request) (*biz{feature}.{feature_camel}, error) {{
	return c.get{feature_camel}(ctx, req)
}}
"""
        with open(f"client/{feature}/endpoint.go", "w") as ef:
            ef.write(endpoint_content)

        # Generate connectrpc_transport.go
        go_client_content = f"""package {feature}

import (
	"context"

	"github.com/pobochiigo/bhole/client/transport"
	biz{feature} "github.com/pobochiigo/bhole/internal/{feature}"
	{feature}v1 "github.com/pobochiigo/bhole/proto/{feature}/v1"
	v1connect "github.com/pobochiigo/bhole/proto/{feature}/v1/{feature}v1connect"
	"connectrpc.com/connect"
)

func New{feature_camel}Client(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) biz{feature}.Service {{
	connectClient := v1connect.New{feature_camel}ServiceClient(httpClient, baseURL, opts...)

	return &endpoints{{
		list{plural_method}: transport.NewConnectClient(
			connectClient.{plural_method},
			encode{plural_method}Request,
			decode{plural_method}Response,
		),
		get{feature_camel}: transport.NewConnectClient(
			connectClient.Get{feature_camel},
			encodeGet{feature_camel}Request,
			decodeGet{feature_camel}Response,
		),
	}}
}}

func encode{plural_method}Request(_ context.Context, req *biz{feature}.{plural_method}Request) (*{feature}v1.{plural_method}Request, error) {{
	return &{feature}v1.{plural_method}Request{{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}}, nil
}}

func encodeGet{feature_camel}Request(_ context.Context, req *biz{feature}.Get{feature_camel}Request) (*{feature}v1.Get{feature_camel}Request, error) {{
	return &{feature}v1.Get{feature_camel}Request{{
		Id:   req.ID,
		Mode: req.Mode,
	}}, nil
}}

"""
        # Add decoders for each feature
        go_client_content += f"""func decode{plural_method}Response(ctx context.Context, resp *{feature}v1.{plural_method}Response) (*biz{feature}.{plural_method}Response, error) {{
	results := make([]biz{feature}.{feature_camel}, len(resp.Results))
	for i, r := range resp.Results {{
		results[i] = *mapProtoToBiz{feature_camel}(r)
	}}
	return &biz{feature}.{plural_method}Response{{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}}, nil
}}

func decodeGet{feature_camel}Response(ctx context.Context, resp *{feature}v1.Get{feature_camel}Response) (*biz{feature}.{feature_camel}, error) {{
	if resp.{feature_camel} == nil {{
		return nil, nil
	}}
	return mapProtoToBiz{feature_camel}(resp.{feature_camel}), nil
}}

"""
        # Generate the proto-to-biz mappings for all dependent types in the feature
        for t_name in sorted(list(feature_types)):
            struct_name = map_schema_name(t_name, main_schema, feature_camel)
            go_client_content += f"func mapProtoToBiz{struct_name}(r *{feature}v1.{struct_name}) *biz{feature}.{struct_name} {{\n"
            go_client_content += f"\tif r == nil {{\n\t\treturn nil\n\t}}\n"
            go_client_content += f"\treturn &biz{feature}.{struct_name}{{\n"
            
            props = get_schema_properties(t_name, schemas)
            for p_name, p_val in sorted(props.items()):
                p_info = resolve_property_info(p_name, p_val, schemas, t_name)
                f_type = map_schema_name(p_info["type"], main_schema, feature_camel)
                    
                g_field = go_field_name(p_name)
                p_proto_go_field = go_proto_field_name(p_name)
                
                if p_info["kind"] == "primitive":
                    # Check if nullable pointer
                    if p_val.get("nullable"):
                        go_client_content += f"\t\t{g_field}: r.{p_proto_go_field},\n"
                    else:
                        go_client_content += f"\t\t{g_field}: r.{p_proto_go_field},\n"
                elif p_info["kind"] == "array_primitive":
                    go_client_content += f"\t\t{g_field}: r.{p_proto_go_field},\n"
                elif p_info["kind"] == "ref":
                    go_client_content += f"\t\t{g_field}: mapProtoToBiz{f_type}(r.{p_proto_go_field}),\n"
                elif p_info["kind"] == "array_ref":
                    # Map array of ref
                    go_client_content += f"\t\t{g_field}: func() []biz{feature}.{f_type} {{\n"
                    go_client_content += f"\t\t\tif r.{p_proto_go_field} == nil {{\n\t\t\t\treturn nil\n\t\t\t}}\n"
                    go_client_content += f"\t\t\tres := make([]biz{feature}.{f_type}, len(r.{p_proto_go_field}))\n"
                    go_client_content += f"\t\t\tfor i, v := range r.{p_proto_go_field} {{\n"
                    go_client_content += f"\t\t\t\tres[i] = *mapProtoToBiz{f_type}(v)\n"
                    go_client_content += f"\t\t\t}}\n"
                    go_client_content += f"\t\t\treturn res\n"
                    go_client_content += f"\t\t}}(),\n"
                    
            go_client_content += "\t}\n}\n\n"
            
        with open(f"client/{feature}/connectrpc_transport.go", "w") as cf:
            cf.write(go_client_content)
            
        # Clean up legacy connectrpc.go file if it exists
        legacy_file = f"client/{feature}/connectrpc.go"
        if os.path.exists(legacy_file):
            os.remove(legacy_file)

    # 2. Generate client/transport/rest_client.go containing all JSON payload mappings
    print("Generating client/transport/rest_client.go...")
    os.makedirs("client/transport", exist_ok=True)
    
    rest_go_content = """package transport

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
"""
    # Import all proto packages
    for feature in sorted(RESOURCES.keys()):
        rest_go_content += f'\t{feature}v1 "github.com/pobochiigo/bhole/proto/{feature}/v1"\n'
    rest_go_content += ")\n\n"
    
    rest_go_content += """type RESTClient struct {
	client  *http.Client
	baseURL string
}

func NewRESTClient(baseURL string, client *http.Client) *RESTClient {
	if client == nil {
		client = http.DefaultClient
	}
	return &RESTClient{
		client:  client,
		baseURL: strings.TrimSuffix(baseURL, "/"),
	}
}
"""
    
    # Generate JSON structures for ALL unique types resolved globally
    print("Generating shared JSON structs...")
    for t_name in sorted(list(all_shared_types)):
        rest_go_content += f"type {t_name}JSON struct {{\n"
        props = get_schema_properties(t_name, schemas)
        for p_name, p_val in sorted(props.items()):
            p_info = resolve_property_info(p_name, p_val, schemas, t_name)
            json_go_type = map_openapi_to_json_go_type(p_info["kind"], p_info["type"])
            if p_val.get("nullable") and p_info["kind"] == "primitive":
                json_go_type = "*" + json_go_type
            rest_go_content += f"\t{go_field_name(p_name)} {json_go_type} `json:\"{p_name}\"`\n"
        rest_go_content += "}\n\n"
        
    # Generate List response JSON wrapper for each feature
    for feature, config in RESOURCES.items():
        feature_camel = camel_case(feature)
        rest_go_content += f"type {config['plural_method']}ResponseJSON struct {{\n"
        rest_go_content += f"\tCount int32 `json:\"count\"`\n"
        rest_go_content += f"\tNext string `json:\"next\"`\n"
        rest_go_content += f"\tPrevious string `json:\"previous\"`\n"
        rest_go_content += f"\tResults []{config['main_schema']}JSON `json:\"results\"`\n"
        rest_go_content += f"}}\n\n"

    # Generate the switch-case HTTP Do method
    rest_go_content += """func (c *RESTClient) Do(req *http.Request) (*http.Response, error) {
	path := req.URL.Path
	reqContentType := req.Header.Get("Content-Type")

	switch path {
"""
    for feature, config in sorted(RESOURCES.items()):
        feature_camel = camel_case(feature)
        main_schema = config["main_schema"]
        
        # Case List
        rest_go_content += f"""\tcase "/{feature}.v1.{feature_camel}Service/{config['plural_method']}":
		var protoReq {feature}v1.{config['plural_method']}Request
		if err := unmarshalRequest(req, &protoReq); err != nil {{
			return nil, err
		}}

		q := url.Values{{}}
		if protoReq.Limit > 0 {{
			q.Set("limit", strconv.Itoa(int(protoReq.Limit)))
		}}
		if protoReq.Offset > 0 {{
			q.Set("offset", strconv.Itoa(int(protoReq.Offset)))
		}}
		if protoReq.Search != "" {{
			q.Set("search", protoReq.Search)
		}}
		if protoReq.Mode != "" {{
			q.Set("mode", protoReq.Mode)
		}}

		restURL := fmt.Sprintf("%s{config['path']}?%s", c.baseURL, q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {{
			return nil, err
		}}
		defer func() {{{{ _ = restResp.Body.Close() }}}}()

		if restResp.StatusCode != http.StatusOK {{
			return makeErrorResponse(restResp)
		}}

		var jsonResp {config['plural_method']}ResponseJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {{
			return nil, err
		}}

		protoResp := &{feature}v1.{config['plural_method']}Response{{
			Count:    jsonResp.Count,
			Next:     jsonResp.Next,
			Previous: jsonResp.Previous,
		}}
		for _, r := range jsonResp.Results {{
			protoResp.Results = append(protoResp.Results, map{main_schema}JSONToProto_{feature}(&r))
		}}

		return writeResponse(reqContentType, protoResp)

"""
        # Case Get
        rest_go_content += f"""\tcase "/{feature}.v1.{feature_camel}Service/Get{feature_camel}":
		var protoReq {feature}v1.Get{feature_camel}Request
		if err := unmarshalRequest(req, &protoReq); err != nil {{
			return nil, err
		}}

		q := url.Values{{}}
		if protoReq.Mode != "" {{
			q.Set("mode", protoReq.Mode)
		}}

		// Map to REST URL path format with ID
		restURL := fmt.Sprintf("%s{config['path']}%v/?%s", c.baseURL, {config['parse_id']}, q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {{
			return nil, err
		}}
		defer func() {{{{ _ = restResp.Body.Close() }}}}()

		if restResp.StatusCode != http.StatusOK {{
			return makeErrorResponse(restResp)
		}}

		var jsonResp {main_schema}JSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {{
			return nil, err
		}}

		protoResp := &{feature}v1.Get{feature_camel}Response{{
			{feature_camel}: map{main_schema}JSONToProto_{feature}(&jsonResp),
		}}

		return writeResponse(reqContentType, protoResp)

"""

    # Add default case
    rest_go_content += """\tdefault:
		return nil, fmt.Errorf("unsupported connectrpc method path: %s", path)
	}
}

func unmarshalRequest(req *http.Request, msg proto.Message) error {
	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}
	contentType := req.Header.Get("Content-Type")
	if strings.Contains(contentType, "application/json") {
		return protojson.Unmarshal(bodyBytes, msg)
	}
	return proto.Unmarshal(bodyBytes, msg)
}

func writeResponse(reqContentType string, msg proto.Message) (*http.Response, error) {
	var body []byte
	var err error
	var contentType string

	if strings.Contains(reqContentType, "application/json") {
		body, err = protojson.Marshal(msg)
		contentType = "application/json"
	} else {
		body, err = proto.Marshal(msg)
		contentType = "application/proto"
	}

	if err != nil {
		return nil, err
	}

	resp := &http.Response{
		StatusCode:    http.StatusOK,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          io.NopCloser(bytes.NewReader(body)),
		ContentLength: int64(len(body)),
		Header:        make(http.Header),
	}
	resp.Header.Set("Content-Type", contentType)
	return resp, nil
}

func makeErrorResponse(resp *http.Response) (*http.Response, error) {
	bodyBytes, _ := io.ReadAll(resp.Body)
	return nil, fmt.Errorf("REST API returned status %d: %s", resp.StatusCode, string(bodyBytes))
}
"""
    
    # Generate mapper functions for each feature
    print("Generating mapping functions in rest_client.go...")
    for feature, config in RESOURCES.items():
        feature_camel = camel_case(feature)
        feature_types = feature_schemas[feature]
        main_schema = config["main_schema"]
        
        for t_name in sorted(list(feature_types)):
            struct_name = map_schema_name(t_name, main_schema, feature_camel)
            rest_go_content += f"func map{t_name}JSONToProto_{feature}(r *{t_name}JSON) *{feature}v1.{struct_name} {{\n"
            rest_go_content += f"\tif r == nil {{\n\t\treturn nil\n\t}}\n"
            rest_go_content += f"\tl := &{feature}v1.{struct_name}{{\n"
            
            props = get_schema_properties(t_name, schemas)
            for p_name, p_val in sorted(props.items()):
                p_info = resolve_property_info(p_name, p_val, schemas, t_name)
                f_type = map_schema_name(p_info["type"], main_schema, feature_camel)
                    
                g_field = go_field_name(p_name)
                p_proto_go_field = go_proto_field_name(p_name)
                
                if p_info["kind"] == "primitive":
                    if p_val.get("nullable"):
                        rest_go_content += f"\t\t{p_proto_go_field}: r.{g_field},\n"
                    else:
                        rest_go_content += f"\t\t{p_proto_go_field}: r.{g_field},\n"
                elif p_info["kind"] == "array_primitive":
                    rest_go_content += f"\t\t{p_proto_go_field}: r.{g_field},\n"
                elif p_info["kind"] == "ref":
                    rest_go_content += f"\t\t{p_proto_go_field}: map{p_info['type']}JSONToProto_{feature}(r.{g_field}),\n"
                elif p_info["kind"] == "array_ref":
                    rest_go_content += f"\t\t{p_proto_go_field}: func() []*{feature}v1.{f_type} {{\n"
                    rest_go_content += f"\t\t\tif r.{g_field} == nil {{\n\t\t\t\treturn nil\n\t\t\t}}\n"
                    rest_go_content += f"\t\t\tres := make([]*{feature}v1.{f_type}, len(r.{g_field}))\n"
                    rest_go_content += f"\t\t\tfor i, v := range r.{g_field} {{\n"
                    rest_go_content += f"\t\t\t\tres[i] = map{p_info['type']}JSONToProto_{feature}(&v)\n"
                    rest_go_content += f"\t\t\t}}\n"
                    rest_go_content += f"\t\t\treturn res\n"
                    rest_go_content += f"\t\t}}(),\n"
                    
            rest_go_content += "\t}\n\treturn l\n}\n\n"
            
    with open("client/transport/rest_client.go", "w") as rf:
        rf.write(rest_go_content)
        
    print("Code generation completed successfully!")

if __name__ == "__main__":
    main()

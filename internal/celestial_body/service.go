package celestial_body

import "context"

type Service interface {
	ListCelestialBodies(ctx context.Context, req *ListCelestialBodiesRequest) (*ListCelestialBodiesResponse, error)
	GetCelestialBody(ctx context.Context, req *GetCelestialBodyRequest) (*CelestialBody, error)
}

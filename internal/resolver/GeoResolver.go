package resolver

import "github.com/johnsoncwb/grapql-test2/internal/models/jsonPlaceholder"

type GeoResolver struct {
	geo jsonPlaceholder.Geo
}

func (g *GeoResolver) Lat() string {
	return g.geo.Lat
}

func (g *GeoResolver) Lng() string {
	return g.geo.Lng
}

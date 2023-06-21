package resolver

import "github.com/johnsoncwb/grapql-test2/internal/models/jsonPlaceholder"

type AddressResolver struct {
	address jsonPlaceholder.Address
}

func (a *AddressResolver) Street() string {
	return a.address.Street
}

func (a *AddressResolver) Suite() string {
	return a.address.Suite
}
func (a *AddressResolver) City() string {
	return a.address.City
}
func (a *AddressResolver) Zipcode() string {
	return a.address.Zipcode
}

func (a *AddressResolver) Geo() *GeoResolver {
	return &GeoResolver{geo: a.address.Geo}
}

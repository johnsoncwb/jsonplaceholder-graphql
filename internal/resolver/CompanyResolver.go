package resolver

import "github.com/johnsoncwb/grapql-test2/internal/models/jsonPlaceholder"

type CompanyResolver struct {
	company jsonPlaceholder.Company
}

func (c *CompanyResolver) Name() string {
	return c.company.Name
}

func (c *CompanyResolver) CatchPhrase() string {
	return c.company.CatchPhrase
}

func (c *CompanyResolver) Bs() string {
	return c.company.Bs
}

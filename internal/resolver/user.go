package resolver

import (
	"fmt"
	"github.com/graph-gophers/graphql-go"
	"github.com/johnsoncwb/grapql-test2/internal/models/jsonPlaceholder"
)

type UserResolver struct {
	user jsonPlaceholder.User
}

func NewUserResolver(user jsonPlaceholder.User) *UserResolver {
	return &UserResolver{user: user}
}

func (u *UserResolver) ID() graphql.ID {
	id := fmt.Sprintf("%d", u.user.Id)
	return graphql.ID(id)
}

func (u *UserResolver) Name() string {
	return u.user.Name
}

func (u *UserResolver) Username() string {
	return u.user.Username
}

func (u *UserResolver) Email() string {
	return u.user.Email
}

func (u *UserResolver) Phone() string {
	return u.user.Phone
}

func (u *UserResolver) Website() string {
	return u.user.Website
}

func (u *UserResolver) Address() *AddressResolver {
	return &AddressResolver{address: u.user.Address}
}

func (u *UserResolver) Company() *CompanyResolver {
	return &CompanyResolver{company: u.user.Company}
}

package resolver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/graph-gophers/graphql-go"
	"github.com/johnsoncwb/grapql-test2/internal/helpers"
	"github.com/johnsoncwb/grapql-test2/internal/models/jsonPlaceholder"
	"github.com/johnsoncwb/grapql-test2/internal/utils"
	"io"
	"net/http"
	"time"
)

type Resolver struct{}

func (r *Resolver) GetUser(ctx context.Context, args struct{ ID graphql.ID }) (*UserResolver, error) {
	id, err := helpers.GqlIDToUint(args.ID)
	if err != nil {
		return nil, err
	}

	client := utils.NewHttpClient(ctx, 10*time.Second)
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%v", id)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = res.Body.Close()
	}()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var output jsonPlaceholder.User

	err = json.Unmarshal(b, &output)
	if err != nil {
		return nil, err
	}

	s := UserResolver{user: output}

	return &s, nil

}

func (r *Resolver) GetAllUsers(ctx context.Context) (*[]*UserResolver, error) {
	client := utils.NewHttpClient(ctx, 10*time.Second)
	url := "https://jsonplaceholder.typicode.com/users/"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = res.Body.Close()
	}()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var output []jsonPlaceholder.User

	err = json.Unmarshal(b, &output)
	if err != nil {
		return nil, err
	}

	var outputSlice []*UserResolver

	for _, user := range output {

		outputSlice = append(outputSlice, &UserResolver{user: user})

	}

	return &outputSlice, nil
}

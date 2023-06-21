# GraphQL 

this project is a graphql interface for [`jsonplaceholder.typicode.com/users/{ID}`][0] 

build using [`graphql-go`][1] and connected to [`newrelic`][2] (you need create a .env with `NEW_RELIC_LICENSE_KEY` and provide your license key to server)

## Usage

```sh
make server
```
make server - runs the servar locally on port 8000
```sh
make image
```
make image - create an image on docker
```sh
make build
```
make build - build image and run on container using port 8000

-----

This launches the HTTP server on port `8000` of your machine.

Visiting http://localhost:8000 will return a GraphiQL client that you can use to make
requests against the API.

enjoy!



[0]: https://jsonplaceholder.typicode.com/users/1
[1]: https://github.com/graph-gophers/graphql-go
[2]: https://newrelic.com/
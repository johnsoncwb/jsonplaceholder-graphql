package main

import (
	"fmt"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/johnsoncwb/grapql-test2/internal/resolver"
	"github.com/johnsoncwb/grapql-test2/internal/schema"
	"github.com/joho/godotenv"
	"github.com/newrelic/go-agent/v3/integrations/nrgraphgophers"
	"github.com/newrelic/go-agent/v3/newrelic"
	"log"
	"net/http"
	"os"
	"time"
)

var page = []byte(`
<!DOCTYPE html>
<html>
	<head>
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.css" />
		<script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/1.1.0/fetch.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react-dom.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.js"></script>
	</head>
	<body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">
		<div id="graphiql" style="height: 100vh;">Loading...</div>
		<script>
			function graphQLFetcher(graphQLParams) {
				return fetch("/query", {
					method: "post",
					body: JSON.stringify(graphQLParams),
					credentials: "include",
				}).then(function (response) {
					return response.text();
				}).then(function (responseBody) {
					try {
						return JSON.parse(responseBody);
					} catch (error) {
						return responseBody;
					}
				});
			}

			ReactDOM.render(
				React.createElement(GraphiQL, {fetcher: graphQLFetcher}),
				document.getElementById("graphiql")
			);
		</script>
	</body>
</html>
`)

func main() {

	var (
		addr              = ":8000"
		readHeaderTimeout = 1 * time.Second
		writeTimeout      = 10 * time.Second
		idleTimeout       = 90 * time.Second
		maxHeaderBytes    = http.DefaultMaxHeaderBytes
	)

	_ = godotenv.Load()

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("GraphQL App"),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
	)
	if nil != err {
		fmt.Println(err)
	}

	s, err := schema.String()
	if err != nil {
		fmt.Println(err)
	}
	opt := graphql.Tracer(nrgraphgophers.NewTracer())

	schemaInternal := graphql.MustParseSchema(s, &resolver.Resolver{}, opt)

	mux := http.NewServeMux()
	mux.Handle(newrelic.WrapHandle(app, "/query", &relay.Handler{Schema: schemaInternal}))
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(page)
	}))

	srv := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
		MaxHeaderBytes:    maxHeaderBytes,
	}

	log.Printf("Listening for requests on %s", srv.Addr)

	if err = srv.ListenAndServe(); err != nil {
		log.Println("server.ListenAndServe:", err)
	}

}

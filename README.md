# FibSrv - A little Fibonacci number generator

This is a little coding exercise that aims to try to squeeze speed out of basic
golang where possible to return the nth Fibonacci number in response to GET http
requests at `$uri/api/fibonacci/:number` where `:number` is the nth position in
the sequence you'd like to see.

If you enjoy watching metrics, you can also view prometheus metrics (or ideally
scrape them with prometheus) at `$uri/metrics`.  The server will log to stdout.

When running locally, this defaults to a URI of localhost:3000.

## Building

If you like to try things without docker, and presuming you have a recent version
of Go that includes support for modules, you should be able to just run:
`make build` at the root of this repository. This will produce a binary at
`bin/fibsrv`.

## Running tests

The really easy way here is `make test`. If you want to try out the benchmarks,
you can run `cd pkg/fibber && go test -bench=.` to see the wild difference between
calculating `big.Int`s with the go code vs just doing a simple lookup in an array.

## Docker

You can build a docker container with `make docker_build`. This will tag the
image to make it fairly easy to run with something like `docker run -e GIN_MODE=release -p 3000:3000 docker.io/crookedstorm/fibsrv:0.0.1` or you can leave off the environment variable for release mode if you
want the default debug mode.

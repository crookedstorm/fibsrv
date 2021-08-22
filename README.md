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
want the default debug mode. However, you really probably should use...


## Docker Compose!
If you want things to be a bit tidier, you can run `make up` and you'll get a 
docker container on port :3000 locally, prometheus monitoring it on :9000 and
cadvisor reporting to prometheus as well.

## The Wish List

This exercise is time-boxed and done in spare time in general. As a result this
is going to be a pale reflection of a production system. Here are some ways I 
would ramp this up for production if the world really was in need of quickly
obtained Fibonacci numbers from the web.

I'm deliberately not putting backups on this list becuase we don't have a lot of
state to worry about. I'm going to blithely trust that my git server will keep
my code safe enough and the various services will still be availble from Dockerhub
tomorrow. For some applications, this isn't always a perfect assumption, but I'll
go with it here. If I was that worried about that, all the code would be forked,
the containers would be cached in a registry, and I would have my bunker stocked.

### A Helm Chart
It would be quicker and more flexible for Kubernetes deployment to chart things
out, add a horizontal pod autoscaler, be ready with the ingress for whatever
controller our cluster is using and take care of all those little concerns.

It probably wouldn't be too hard to squeeze in an optional TLS cert from
certmanager/Let's Encrypt that's actually valid, all that the docker compose
does and then some. I inherited the care and feeding of the helm setup for
[PAWS](https://github.com/toolforge/paws), which is mostly using a helm chart to
customize and bring together other helm charts. That sort of setup would be
lovely for setting things up given a little time and testing. This also implies
that there's a k8s cluster to plug into.

### Actual Parallelism
As hinted at above, I don't think starting one multithreaded golang process is all
that mighty. Golang's http primitives are remarkably capable, but that doesn't
mean a concerted effort wouldn't require horizontal scaling behind a load balancer.

The dynamic load balancing you can get from autoscaling mechanisms is a thing of
beauty. Kubernetes can do that behind a service regardless of platform if you
don't rely on "load balancer" service types for everything.

### External Monitoring

The telescope was invented for more reasons than viewing the sky. It pays to be
able to tell if the entire system is down, not just one pod. You cannot get an
alert from inside a network that cannot communicate, for instance.

### A Robust and HA Front Proxy

A couple servers in a VRRP setup with keepalived and a configuration they share
via puppet/chef/salt or some such nonsense not only keeps a stable IP for the
outside world, it can also protect system that weren't written specifically for 
the purpose from things like slowloris attacks and provide a quick reference for
logging, DDoS mitigation, etc.

### Remote Logging

You really shouldn't dockerize things without sending the logs somewhere. It just
makes things harder.

### Staging Environment

SLOs are tricky to develop at the best of times, but if you are testing performance
on dramatically different hardware than you might be deploying on, it's just guessing.

I'm developing this on an ARM64 Mac. There isn't much chance of that being the
deployment platform. Given a bit of time, I might want to try it out on my
staging environment and then beat on it a while to decide just what acceptable
looks like.

### Histogram Buckets in Prometheus

I used a quick open source solution for exporting prometheus stats instead of
concocting my own middleware. That's often a great time saver, but if we are working
on performance SLOs, quantiles would be very nice. That would take me a little more
time.

### Grafana (or similar)

Prometheus is nice, but it really doesn't "give" you much without some dashboards
made while you had some time to think about it. Dashboards and visualizations are
a gift to your future self that allow you to see things you didn't know you were
looking for and find the patterns that actually help you decide if that SLA you 
agreed to was a bad idea.

## Some Tests

The benchmarks and tests in the go code looked promising, but I know that
I'd done nothing to stop golang from just running out of threads as it takes load

When I cranked up `rakyll/hey` a little, I got what I was after (even using
the precalculated array of numbers):
```
 hey -n 600 -c 250 http://localhost:3000/api/fibonacci/20

Summary:
  Total:	0.1168 secs
  Slowest:	0.1013 secs
  Fastest:	0.0015 secs
  Average:	0.0549 secs
  Requests/sec:	4282.1769

  Total data:	1392 bytes
  Size/request:	4 bytes

Response time histogram:
  0.001 [1]	|■
  0.011 [44]	|■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.021 [55]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.031 [3]	|■■
  0.041 [34]	|■■■■■■■■■■■■■■■■■■■■
  0.051 [5]	|■■■
  0.061 [20]	|■■■■■■■■■■■■
  0.071 [20]	|■■■■■■■■■■■■
  0.081 [68]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.091 [60]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.101 [38]	|■■■■■■■■■■■■■■■■■■■■■■


Latency distribution:
  10% in 0.0087 secs
  25% in 0.0132 secs
  50% in 0.0698 secs
  75% in 0.0870 secs
  90% in 0.0920 secs
  95% in 0.0970 secs
  99% in 0.1004 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0078 secs, 0.0015 secs, 0.1013 secs
  DNS-lookup:	0.0034 secs, 0.0000 secs, 0.0094 secs
  req write:	0.0002 secs, 0.0000 secs, 0.0029 secs
  resp wait:	0.0393 secs, 0.0014 secs, 0.0803 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0002 secs

Status code distribution:
  [200]	348 responses

Error distribution:
  [4]	Get "http://localhost:3000/api/fibonacci/20": dial tcp [::1]:3000: socket: too many open files
  [3]	Get "http://localhost:3000/api/fibonacci/20": dial tcp: lookup localhost: no such host
  [1]	Get "http://localhost:3000/api/fibonacci/20": read tcp [::1]:57091->[::1]:3000: read: connection reset by peer
  [1]	Get "http://localhost:3000/api/fibonacci/20": read tcp [::1]:57092->[::1]:3000: read: connection reset by peer
  [1]	Get "http://localhost:3000/api/fibonacci/20": read tcp [::1]:57093->[::1]:3000: read: connection reset by peer
  [1]	Get "http://localhost:3000/api/fibonacci/20": read tcp [::1]:57095->[::1]:3000: read: connection reset by peer
<cut off for brevity>
```
It runs fast, but eventually, it gets hurt. Also, prometheus didn't notice any
problems because I wasn't running a proxy or load balancer yet to detect the failures
from.


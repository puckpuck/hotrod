# Hot R.O.D. - Rides on Demand

This is a demo application that consists of several microservices and illustrates
the use of the OpenTracing API. It can be run standalone, but requires a Jaeger compatible 
collector to gather the data like the OpenTelemetry Collector.

This project was originally cloned from [Jaeger Examples](https://github.com/jaegertracing/jaeger/tree/2fdd1cd7454e617148c1f8ec44f87a37bd5d52f4/examples/hotrod)


## Features

* Discover architecture of the whole system via data-driven dependency diagram
* View request timeline & errors, understand how the app works
* Find sources of latency, lack of concurrency
* Highly contextualized logging
* Use baggage propagation to
  * Diagnose inter-request contention (queueing)
  * Attribute time spent in a service
* Use open source libraries with OpenTracing integration to get vendor-neutral instrumentation for free

## Running

### Run HotROD from source
Adjust JAEGER_ENDPOINT environment variable to point to an instance of OpenTelemetry Collector.
```bash
git clone https://github.com/puckpuck/hotrod
cd hotrod 
export JAEGER_ENDPOINT=http://opentelemetry-collector:14268/api/traces
go run ./main.go all
```

### Run HotROD from docker
Adjust JAEGER_ENDPOINT environment variable to point to an instance of OpenTelemetry Collector.
```bash
docker run \
  --env JAEGER_ENDPOINT=http://opentelemetry-collector:14268/api/traces \
  -p8080:8080 \
  puckpuck/hotrod:latest \
  all
```

Then open http://127.0.0.1:8080


<p align="center"><img alt="kind" src="./karmen.png" width="300x" /></p>

# Karmen ![Actions Status](https://github.com/jrcichra/karmen/workflows/Karmen/badge.svg) [![Docker Hub](https://img.shields.io/badge/docker-hub-blue.svg)](https://hub.docker.com/r/jrcichra/) [![Go Report Card](https://goreportcard.com/badge/github.com/jrcichra/karmen)](https://goreportcard.com/report/github.com/jrcichra/karmen)

Centralized Pub/Sub for microservices

## Get Started
### Using Karmen (Server):
1. See [Docker Hub](https://github.com/jrcichra/karmen/releases) for releases
2. See [an example config](./example.yml) to start declaring your workflow
3. Run Karmen as part of your docker-compose.yml. see my [ example docker-compose.yml](./example_docker-compose.yml)
### Clients with examples:
+ [Python](./pythonclient) - example in `karmen.py` when executed as script
+ [Golang](./goclient)     - example in `main.go`
+ Or write your own! Karmen runs on gRPC. See existing implementations for reference
### Powered by
+ [gRPC](https://grpc.io/)
+ [Golang](https://golang.org/)
+ [Docker](https://www.docker.com/)
### Projects using Karmen
1. https://github.com/jrcichra/smartcar

### Similar projects
1. https://netflix.github.io/conductor/

### More docs to come!

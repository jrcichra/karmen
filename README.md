<p align="center"><img alt="kind" src="https://raw.githubusercontent.com/jrcichra/karmen/master/karmen.png" width="300x" /></p>

# Karmen ![Actions Status](https://github.com/jrcichra/karmen/workflows/Karmen/badge.svg) [![Docker Hub](https://img.shields.io/badge/docker-hub-blue.svg)](https://hub.docker.com/r/jrcichra/) [![Go Report Card](https://goreportcard.com/badge/github.com/jrcichra/karmen)](https://goreportcard.com/report/github.com/jrcichra/karmen)

Centralized Pub/Sub for microservices


# 2.0 release
+ I rewrote Karmen from the ground up in July 2021. It is incompatible with version 1.0.
## Enhancements
+ Karmen now runs off of gRPC, which cuts down lots of nasty bugs
+ `if:` as a key under an action is now a reserved word for conditional expressions
+ `parallel` and `serial` blocks should perform how you expect...each block is done serially
+ Context variables - actions can return parameters that are injected into a block-level state. These can be referenced in conditionals with `{hostname-action-variablename}`
    + I'll be converting the dashes to dots once I add that feature to the condition parser I used
    + For each action, the `{hostname-action-pass}` boolean is set automatically so you can conditionally run actions based on the result of previous actions without managing a parameter. Code 200 is defined as a `pass`
+ Action error handling is improved, currently returning HTTP-like codes. I may downgrade this to a boolean

## Get Started
### Using Karmen (Server):
1. See [Docker Hub](https://github.com/jrcichra/karmen/releases) for releases
2. See [an example config](./example.yml) to start declaring your workflow
3. Run Karmen as part of your docker-compose.yml. see my [ example docker-compose.yml](./example_docker-compose.yml)
### Clients with examples:
+ [Python](./pythonclient) - example in `karmen.py` when executed as script
+ [Golang](./goclient)     - example in `main.go`
+ [Rust](./rustclient)     - example in `example/src/main.rs` [Crates.io](https://crates.io/crates/karmen)
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

<p align="center"><img alt="kind" src="./karmen.png" width="300x" /></p>

# Karmen ![Actions Status](https://github.com/jrcichra/karmen/workflows/Karmen/badge.svg) [![Docker Hub](https://img.shields.io/badge/docker-hub-blue.svg)](https://hub.docker.com/r/jrcichra/) [![Go Report Card](https://goreportcard.com/badge/github.com/jrcichra/karmen)](https://goreportcard.com/report/github.com/jrcichra/karmen)

Centralized Pub/Sub for microservices

## Get Started
### Using Karmen:
1. See [Docker Hub](https://github.com/jrcichra/karmen/releases) for releases
2. See [an example config](./example_config.yml) to start declaring your workflow
3. Run Karmen as part of your docker-compose.yml. see my [ example docker-compose.yml](./example_docker-compose.yml)
### Using Karmen's Python Client:
1. Install using `pip install karmen`
2. Usage example:
```python
import karmen

# Function that performs an action and returns a result
def hello(params,result):
    print("Hello, world!")
    result.Pass()
# Spawn a karmen client
k = karmen.Client()
# Register this client with the karmen server (based on hostname)
k.registerContainer()
# Register an event with the karmen server
k.registerEvent("docker_rocks")
# Register an action with the karmen server
k.registerAction("hello", hello)
# Emit an event called docker_rocks - this is declared in config.yml
k.emitEvent("docker_rocks")
```

### Projects using Karmen
1. https://github.com/jrcichra/smartcar

### Similar projects
1. https://netflix.github.io/conductor/

### More docs to come!

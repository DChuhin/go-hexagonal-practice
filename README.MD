# Sample Go Application in Hexagonal Style

This is sample REST CRUD application of hexagonal approach to implement go services

Is supposed to bu run in AWS behind ALB with SSO authentication

## Folder Structure

```
project root
├── cmd
│   └── main.go
├── internal
│   ├── adapters
│   │   ├── input               # input adapters split by input interface (e.g. rest, grpc, message consumers)
│   │   │   ├── rest
│   │   │   ├── kafka
│   │   │   └── grpc
│   │   └── output              # output adapters split by backing services (databases, brokers, message publishers)
│   │       ├── mongodb
│   │       └── kafka
│   ├── application             # business layer inclidint contracts (port) and impl (services)       
│   │   ├── ports       
│   │   │   ├── input           # input ports contract: interfaces and structures     
│   │   │   └── output          # output ports interfaces
│   │   └── service             # input ports implementations, split by domain    
│   ├── boot                    # wires application components (DI): outpud adapters into services, services into input adapters and starts input adapters    
│   ├── docs                    # swagger compiled files    
│   ├── domain                  # domain model. Alternatively can be placed into application/ports of small projects    
│   ├── integration             # domain agnositc code to contain utilities to work with backing services    
│   └── util                    # domain agnositc low level utils    
├── runConfigurations           # Goland shared run configs (build, run, test)
│   └── local.run.xml           # running application locally configs
├── docker-compose.yaml         # compose file used to run locally and for integration tests with testcontainers
├── Dockerfile
├── go.mod
├── local.env                   # .env file used for local runs
└── Makefile
```

## Build

Use makefile

`make build`

## Run locally

### Run mongo db

Run mongo in docker before running service:
`docker compose up`

### Run service

Use shared Goland run configuration if you use EnvFile plugin or create new one with propagated env variables from
`local.env`


[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![Documentation](https://godoc.org/github.com/jmbarzee/bitbox?status.svg)](https://godoc.org/github.com/jmbarzee/bitbox)
[![Code Quality](https://goreportcard.com/badge/github.com/jmbarzee/bitbox)](https://goreportcard.com/report/github.com/jmbarzee/bitbox)
[![Build Status](https://github.com/jmbarzee/bitbox/workflows/build/badge.svg)](https://github.com/jmbarzee/bitbox/actions)
[![Coverage](https://codecov.io/gh/jmbarzee/bitbox/branch/main/graph/badge.svg)](https://codecov.io/gh/jmbarzee/bitbox)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fjmbarzee%2Fbitbox.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fjmbarzee%2Fbitbox?ref=badge_shield)

# BitBox
Rudementary implementation of containerization in linux.


## Server
For available endpoints view the [gRPC deffinition](grpc/bitbox.proto).
```bash
# Build the Server.
go build -o bin/bitbox cmd/server/main.go
# Run it.
./bin/bitbox
```


## CLI Client
```bash
# Build the CLI
go build -o bin/bitboxcli cmd/client/*.go
# Run it.
./bin/bitboxcli start <process>
```

## Development Tools
[gRPC & protoc](https://grpc.io/docs/languages/go/quickstart/) are used by `go generate` to update [bitbox/grpc](grpc/).

[gRPCox](https://github.com/gusaul/grpcox) is a lightweight docker container for easy manual testing.
# GoLang experiment

- Install go: [go.dev](https://go.dev/dl/)
- Install [vscode extensions for go](https://marketplace.visualstudio.com/items?itemName=golang.go)
- For debugging, install dlv: `go install -v github.com/go-delve/delve/cmd/dlv@latest`

## Install dependencies

`./scripts/install_deps.sh`

## Run it

Multiple choices:

1. use .vscode Debug preset in debug options
1. run `go run cmd/b2b-service-pmp/main.go`
1. build a binary an then run it:

- build the binary with: `./scripts/build.sh`
- execute it with `./scripts/run.sh`

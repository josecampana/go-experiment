# GoLang experiment

- Install go: [go.dev](https://go.dev/dl/)
- Install [vscode extensions for go](https://marketplace.visualstudio.com/items?itemName=golang.go)
- For debugging, install dlv: `go install -v github.com/go-delve/delve/cmd/dlv@latest`

## Run it

Multiple choices:

1. use .vscode Debug preset in debug options
1. run `go run src/main.go`
1. build a binary an then run it:

- build the binary with: `go build -o pmp src/main.go`
- execute it with `./pmp`

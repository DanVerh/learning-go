## Development tools
Development tools are accessible with `go` command:
- version: `go version`
- compiler: `go build`
- code formatter: `go fmt`
- dependency manager: `go pkg`
- test runner: `go test`
- scanner: `go vet`

## Project initialization
- go project = go module
- module = source code + dependecies
- every module has go.mod
- go.mod: module name, minimum supported go version and other modules-dependencies
- go mod init = create a module and go.mod file
```
go mod init <module_name>
```
# mcptest

A simple tool that allows to test your MCP servers using MCP protocol by defining YAML files with requests and responses.

MCP cases file is a multi-document YAML, which defines every case as a separated document, like this:

```yaml
case: Initialize
in: {"jsonrpc":"2.0","id":0,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"testing","version":"0.0.1"}}}
out: {"result":{"protocolVersion":"2024-11-05","capabilities":{"resources":{},"tools":{}},"serverInfo":{"name":"example-servers/postgres","version":"0.1.0"}},"jsonrpc":"2.0","id":0}

---
# next case...
```


## Installation

## npm

```bash
npm install -g mcptest
```

## Github Releases

Download prebulit binaries from the [releases](https://github.com/strowk/mcptest/releases) page and put in your PATH

## Build from source

```bash
go get github.com/strowk/mcptest
go install github.com/strowk/mcptest
```

## Usage

```bash
mcptest [flags] run [--] path/to/folder/with/test/scenarios command-to-run-mcp-server [server-args]
```

Example:
```bash
mcptest run testdata go run main.go
mcptest run -v testdata -- npx -y @modelcontextprotocol/server-postgres localhost:5432
```

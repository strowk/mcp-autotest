# mcp-autotest

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
npm install -g mcp-autotest
```

## Github Releases

Download prebulit binaries from the [releases](https://github.com/strowk/mcp-autotest/releases) page and put in your PATH

## Build from source

```bash
go get github.com/strowk/mcp-autotest
go install github.com/strowk/mcp-autotest
```

## Usage

```bash
mcp-autotest [flags] run [--] path/to/folder/with/test/scenarios command-to-run-mcp-server [server-args]
```

Example:
```bash
# start go MCP server and test via stdio transport
mcp-autotest run testdata go run main.go
# start Postgres MCP server and test via stdio transport
mcp-autotest run -v testdata -- npx -y @modelcontextprotocol/server-postgres localhost:5432
# start go MCP server and test via Streamable HTTP transport
mcp-autotest run --url http://localhost:8080/mcp testdata go run main.go
```

### Transports

mcp-autotest by default would use stdio transport, but if you want to use HTTP transport instead, you can use `--url` flag to specify the URL of the MCP server. 
The URL must be in the format `http://host:port/path` or `https://host:port/path`.

Currently from Streamable HTTP only `POST` method is supported, but testing via `GET` is planned for the future.

## Quick Demo

In bash shell run following

```bash

# create folder for test data
mkdir -p testdata

# create test cases file
cat << EOF > testdata/list_tools_test.yaml
# This is a test cases file. It contains a list of test cases in YAML format.
# Each test case has variable number of inputs (keyst starting with 'in') and outputs (keys starting with 'out').
# The test cases are separated by '---' (three dashes) on a new line, making it multi-document YAML file.
# File name must end with '_test.yaml' to be recognized as a test cases file.

case: List tools

# requesting list of tools
in: { "jsonrpc": "2.0", "method": "tools/list", "id": 1 }

# expect one tool in the list
out:
  {
    "jsonrpc": "2.0",
    "id": 1,
    "result":
      {
        "tools":
          [
            {
              "description": "Run a read-only SQL query",
              "inputSchema":
                {
                  "type": "object",
                  "properties": { "sql": { "type": "string" } },
                },
              "name": "query",
            },
          ],
      }
  }
EOF

# Now running autotest
npx mcp-autotest run testdata -- npx -y @modelcontextprotocol/server-postgres localhost:5432
```

The output should simply print one word `PASS`.

Now if you change something in line `"description": "Run a read-only SQL query",`, f.e to `"description": "Run a read-only SQL query 2"`, and run the last command again, you should see the output like this:

```bash
2025/03/31 22:23:39 actual json did not match expectation,
 got: '{"id":1,"jsonrpc":"2.0","result":{"tools":[{"description":"Run a read-only SQL query","inputSchema":{"properties":{"sql":{"type":"string"}},"type":"object"},"name":"query"}]}}'
 diff with expected:
  "result": {
    "tools": {
      "0": {
        "description": {
        ^ value mismatch,
expected string: 'Run a read-only SQL query 2',
     got string: 'Run a read-only SQL query'
FAIL
```


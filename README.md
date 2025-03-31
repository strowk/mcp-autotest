<h4 align="center">Autotest MCP servers in a language-agnostic way</h4>

<h1 align="center">
   <img src="docs/images/logo.png" width="180"/>
   <br/>
   mcp-autotest
</h1>

<p align="center">
    <a href="https://github.com/strowk/mcp-autotest/actions/workflows/test.yaml"><img src="https://github.com/strowk/mcp-autotest/actions/workflows/test.yaml/badge.svg"></a>
	<a href="https://github.com/strowk/mcp-autotest/actions/workflows/golangci-lint.yaml"><img src="https://github.com/strowk/mcp-autotest/actions/workflows/golangci-lint.yaml/badge.svg"/></a>
    <a href="https://goreportcard.com/report/github.com/strowk/mcp-autotest"><img src="https://goreportcard.com/badge/github.com/strowk/mcp-autotest" alt="Go Report Card"></a>
</p>

<p align="center">
  <a href="#installation">Installation</a> ⚙
  <a href="#usage">Usage</a> ⚙
  <a href="#quick-demo">Quick Demo</a> ⚙
  <a href="#dynamic-matching">Dynamic Matching</a>
</p>

A simple tool that allows to test your MCP servers using MCP protocol by defining YAML files with requests and responses.

MCP cases file is a multi-document YAML, which defines every case as a separated document, like this:

```yaml
case: Listing tools
in: { "jsonrpc": "2.0", "method": "tools/list", "id": 1 }
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
go install github.com/strowk/mcp-autotest@latest
```

## Usage

```bash
mcp-autotest [flags] run path/to/tests/folder [--] command-to-run-server [server-args]
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

## Dynamic Matching

Sometimes when you have dynamic values in your responses, you might want to use regular expressions to match them. You can do this by using `!!re` tag in your expected output. For example:

```yaml
case: List tools
in_list_tools: {"jsonrpc":"2.0","method":"tools/list","id":1}
out_list_tools: {"jsonrpc":"2.0","result":{"tools":[ { "name": !!re "list-[a-z]+" } ]},"id":1}
```

This example is somewhat oversimplified, in reality you would probably use `!!re` with timestamps, UUIDs or other dynamic values, rather than tool names.
It does, however, demonstrate that strings with `!!re` tag would be treated as regular expressions and would be matched against actual values using regular expression matching instead of string equality.

See working examples for testing [postgres server](https://github.com/strowk/mcp-autotest/tree/main/examples/postgres).

### Embedded regex

Embedded regular expression is just a regular expression that is embedded in a string. It is used to match a part of the string. For example:

```yaml
case: List tools
in_list_tools: {"jsonrpc":"2.0","method":"tools/list","id":1}
out_list_tools: {"jsonrpc":"2.0","result":{"tools":[ { "name": !!ere "list-/[a-z]+/-tool" } ]},"id":1}
```

Essentially `!!ere` allows you to treat most of the string as just regular string, but have a part (or parts) of it treated as regular expression.
Everything inside slashes `/` would be treated as regular expression, here "list-" and "-tool" are regular strings, but `[a-z]+` is a regular expression that would match any lowercase letters non-empty string, so it would match "list-abc-tool", "list-xyz-tool" and so on.

This approach allows you not to think how to escape regular expression syntax in the rest of your string and only match bits that you need to be dynamic.

### Escaping forward slashes

For embedded regex you might need to escape slashes `/` in places of your string where you want to use them, but not designate regular expression, for example: `"url": !! "https:\\/\\/github.com\\//[a-z]+/"` would match `"url": "https://github.com/strowk"`.
In here `\\/` is used to become `/` and `[a-z]+` is used to match any lowercase letters non-empty string. The reason why there are two backslashes `\\` is because in YAML strings backslash is an escape character, so to have a single backslash in the string you need to escape it with another backslash.
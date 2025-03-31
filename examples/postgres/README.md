# Autotesting Postgres MCP Server Example

Run following command in this folder to see the tests in action:

```bash
npx mcp-autotest run testdata -- npx -y @modelcontextprotocol/server-postgres localhost:5432
```

You can then change something in `testdata/list_tools_test.yaml` and run the command again to see test failing.

This test does not require to actually have Postgres server running, as it only tests for list of tools.


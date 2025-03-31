package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/strowk/foxy-contexts/pkg/foxytest"
)

const description = `mcp-autotest run is a command to test Model Context Protocol servers with MCP Story format.

DESCRIPTION

Tool would use stdio transport by default, or would use 
streamable HTTP transport if --url (shotcut -u) flag is specified
with url to the server, for example http://localhost:8080/mcp
		
Test scenarios are defined in a files ending 
with _test.yaml in the folder with test scenarios.

Example running tests from testdata folder:
mcp-autotest run testdata go run main.go

Yaml file with scenario can contain multiple documents 
separated by '---' line. Each document must contain 'case' 
field with the name of the test case, any number of in* 
fields with input data, and out* fields with expected 
output data.

Example testdata/run_test.yaml:
		
case: List Tools
in: {"jsonrpc": "2.0", "method": "tools/list", "id": 1}
out: {
  "jsonrpc": "2.0", 
  "id": 1,
  "result": { 
    "tools": [
      {
        description": "My tool", 
        inputSchema": {"type": "object"}, 
        name": "my-tool"
      }
    ] 
  }
}

When running server with flags - any arguments starting with dash,
you would need to use '--' to separate flags for mcp-autotest and flags for the server.
For example in following command '--' added to separate flags for mcp-autotest and flags for npx:
mcp-autotest -v run testdata -- npx -y @modelcontextprotocol/server-postgres localhost:5432
`

func main() {
	// main_urfave()
	main_cobra()
}

const (
	verbose = "verbose"
	url     = "url"
)

var (
	Verbose bool
	Url     *string
)

func init() {
	runCommand.Flags().BoolVarP(&Verbose, verbose, "v", false, "verbose output")
	Url = runCommand.Flags().StringP(url, "u", "", "url to the server to test for streamable HTTP transport, if not specified, stdio is used")
	runCommand.SetUsageFunc(func(cmd *cobra.Command) error {
		fmt.Println(`USAGE

  mcp-autotest [flags] run [--] path/to/folder/with/test/scenarios command-to-run-mcp-server [server-args]

FLAGS

  -v, --verbose   verbose output
  -h, --help      help for mcp-autotest
  -u, --url       url to the server to test for streamable HTTP transport, if not specified, stdio is used

EXAMPLES

  mcp-autotest run testdata go run main.go
  mcp-autotest run -v testdata -- npx -y @modelcontextprotocol/server-postgres localhost:5432
  mcp-autotest run --url http://localhost:8080/mcp testdata go run main.go`)
		return nil
	})
	cobraCmd.AddCommand(runCommand)
}

var cobraCmd = &cobra.Command{
	Use: "mcp-autotest",
}

var runCommand = &cobra.Command{
	Use:   "run",
	Short: "functional language-agnostic tests for MCP servers",
	Long:  description,
	Args:  cobra.MinimumNArgs(2),
	// This is for completion
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) == 0 {
			// first arg is path to scenarios
			return nil, cobra.ShellCompDirectiveFilterDirs
		}
		if len(args) == 1 {
			// second arg is command to run server
			return nil, cobra.ShellCompDirectiveDefault
		}
		return nil, cobra.ShellCompDirectiveDefault
	},
	Run: func(cmd *cobra.Command, args []string) {
		t := &foxyT{}

		ts, err := foxytest.Read(args[0])
		if err != nil {
			t.Fatal(err)
		}
		ts.WithExecutable(args[1], args[2:])
		if Url != nil {
			if !strings.HasPrefix(*Url, "http://") && !strings.HasPrefix(*Url, "https://") {
				t.Fatalf("URL should start with 'http://' or 'https://', but was %s", *Url)
			}
			ts.WithTransport(foxytest.NewTestTransportStreamableHTTP(*Url))
		}
		if Verbose {
			t.verbose = true
			ts.WithLogging()
			if Url != nil {
				fmt.Printf("running tests with server at %s\n", *Url)
			} else {
				fmt.Println("running tests with server at stdio")
			}
		}

		ts.Run(t)
		ts.AssertNoErrors(t)
		if !t.hadErrors {
			fmt.Println("PASS")
		} else {
			fmt.Println("FAIL")
			os.Exit(1)
		}
	},
}

func main_cobra() {
	if err := cobraCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type foxyT struct {
	hadErrors bool
	verbose   bool
}

func (t *foxyT) Fatal(args ...any) {
	log.Fatal(args...)
}

func (t *foxyT) Fatalf(format string, args ...any) {
	log.Fatalf(format, args...)
}

func (t *foxyT) Errorf(format string, args ...any) {
	t.hadErrors = true
	fmt.Printf(format, args...)
	fmt.Println()
}

func (t *foxyT) Log(args ...any) {
	if t.verbose {
		fmt.Println(args...)
	}
}

func (t *foxyT) Logf(format string, args ...any) {
	if t.verbose {
		fmt.Printf(format, args...)
		fmt.Println()
	}
}

func (t *foxyT) Run(name string, f func(t foxytest.TestRunner)) bool {
	if t.verbose {
		fmt.Printf("RUN %s\n", name)
	}
	f(t)
	return !t.hadErrors
}

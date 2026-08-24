package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"bs-euro/internal/api"
	"bs-euro/internal/bs"
	"bs-euro/internal/io"
)

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	if len(args) == 0 {
		return serve([]string{})
	}
	switch args[0] {
	case "serve":
		return serve(args[1:])
	case "price":
		return priceCommand(args[1:])
	case "greeks":
		return greeksCommand(args[1:])
	case "example":
		return exampleCommand(args[1:])
	case "check":
		return checkCommand(args[1:])
	case "help", "-h", "--help":
		usage()
		return 0
	default:
		fmt.Fprintf(os.Stderr, "unknown command %q\n", args[0])
		usage()
		return 2
	}
}

func serve(args []string) int {
	fs := flag.NewFlagSet("serve", flag.ContinueOnError)
	addr := fs.String("http", ":8080", "HTTP listen address")
	if err := fs.Parse(args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 2
	}
	if err := api.Listen(*addr); err != nil {
		fmt.Fprintf(os.Stderr, "serve: %v\n", err)
		return 1
	}
	return 0
}

func priceCommand(args []string) int {
	path, table, err := parsePathArgs(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 2
	}
	if path == "" {
		fmt.Fprintln(os.Stderr, "usage: bs-euro price <input.json> [--table]")
		return 2
	}
	if err := io.RunPrice(path, table, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "price: %v\n", err)
		return 1
	}
	return 0
}

func greeksCommand(args []string) int {
	path, table, err := parsePathArgs(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 2
	}
	if path == "" {
		fmt.Fprintln(os.Stderr, "usage: bs-euro greeks <input.json> [--table]")
		return 2
	}
	if err := io.RunGreeks(path, table, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "greeks: %v\n", err)
		return 1
	}
	return 0
}

func exampleCommand(args []string) int {
	path := io.ExamplePath()
	table := len(args) > 0 && args[0] == "--table"
	if err := io.RunPrice(path, table, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "example: %v\n", err)
		return 1
	}
	return 0
}

func checkCommand(args []string) int {
	path := io.ExamplePath()
	if len(args) > 0 && !strings.HasPrefix(args[0], "-") {
		path = args[0]
	}
	in, err := io.LoadOptionFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "check: %v\n", err)
		return 1
	}
	res, err := bs.Compute(in)
	if err != nil {
		fmt.Fprintf(os.Stderr, "check: %v\n", err)
		return 1
	}
	fmt.Fprintf(os.Stdout, "case %q: price=%g flag=%s\n",
		label(in.Label), res.Price, res.Flag)
	return 0
}

func parsePathArgs(args []string) (string, bool, error) {
	path := ""
	table := false
	for _, a := range args {
		switch {
		case a == "--table" || a == "-table":
			table = true
		case strings.HasPrefix(a, "-"):
			return "", false, fmt.Errorf("unknown flag %q", a)
		case path == "":
			path = a
		default:
			return "", false, fmt.Errorf("unexpected argument %q", a)
		}
	}
	return path, table, nil
}

func usage() {
	fmt.Fprintln(os.Stdout, `bs-euro: European Black-Scholes option pricing calculator

Commands:
  serve [--http :8080]       start the HTTP service (default command)
  price <input.json> [--table]
  greeks <input.json> [--table]
  example [--table]          run example/atm-1y.json
  check [input.json]         load a case and print the price verdict`)
}

func label(s string) string {
	if s != "" {
		return s
	}
	return "untitled"
}

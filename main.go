package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const help = `Examples:
command:
  {exec} 123 383764 text
print:
  123B
  374.76KB
  NaN(text)

command: 
  echo 123 | {exec}
print:
  123B

Arguments:
  -h, --help  Print this help
  -p          Precision
`

func main() {
	args := os.Args[1:]

	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			fmt.Print(strings.ReplaceAll(help, "{exec}", filepath.Base(os.Args[0])))
			return
		}
	}

	for i, arg := range args {
		if arg == "-p" {
			if i+1 >= len(args) {
				exitf("missing value for argument -p")
			}

			var err error
			prec, err = strconv.Atoi(args[i+1])
			if prec < 0 || err != nil {
				exitf("invalid value %s for argument -p", args[i+1])
			}

			args = args[i+2:]
			break
		}
	}

	if len(args) > 0 {
		for _, arg := range args {
			parse(arg)
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			parse(scanner.Text())
		}
		if scanner.Err() != nil {
			exitf("error reading stdin: %s", scanner.Err())
		}
	}
}

func exitf(format string, a ...interface{}) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(format, a...))
	os.Exit(-2)
}

var units = [...]string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}

func parse(raw string) {
	scale, _ := new(SizeNum).From("1024")
	size, err := new(SizeNum).From(strings.TrimSpace(raw))
	if err != nil {
		fmt.Printf("NaN(%s)\n", raw)
		return
	}

	var label string
	for _, label = range units {
		if size.Gte(scale) {
			size.Div1024()
		} else {
			break
		}
	}

	fmt.Printf("%s%s\n", size, label)
}

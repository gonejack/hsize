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
  {exec} 123 45678 text
print:
  123B
  44.61KB
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
				fmt.Fprint(os.Stderr, "missing value for argument -p")
				return
			}
			var err error
			prec, err = strconv.Atoi(args[i+1])
			if err == nil {
				args = args[i+2:]
				break
			} else {
				fmt.Fprintf(os.Stderr, "invalid value %s for argument -p", args[i+1])
				return
			}
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
			fmt.Printf("error reading stdin: %s", scanner.Err())
		}
	}
}

var units = [...]string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}
var scale, _ = new(SizeNum).from("1024")

func parse(raw string) {
	size, err := new(SizeNum).from(strings.TrimSpace(raw))
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

package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"
)

const help = `Examples:
command:
  {exec} 123 45678
print:
  123 => 123B
  45678 => 44.61KB

command: 
  echo 123 | {exec}
print:
  123 => 123B
`

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			if arg == "-h" || arg == "--help" {
				fmt.Print(strings.ReplaceAll(help, "{exec}", filepath.Base(os.Args[0])))
				return
			}
		}
		for _, arg := range os.Args[1:] {
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
var scale = big.NewInt(1024)

func parse(raw string) {
	size, ok := new(big.Int).SetString(strings.TrimSpace(raw), 10)
	if !ok {
		fmt.Printf("can not parse \"%s\"\n", raw)
		return
	}

	var value, label string
	var unit = big.NewInt(1)
	var next = big.NewInt(0).Set(scale)
	for _, label = range units {
		if size.Cmp(next) >= 0 {
			unit = unit.Mul(unit, scale)
			next = next.Mul(next, scale)
		} else {
			break
		}
	}

	rat := new(big.Rat).SetFrac(size, unit)
	if rat.IsInt() {
		value = rat.FloatString(0)
	} else {
		value = rat.FloatString(2)
	}
	fmt.Printf("%s => %s%s\n", raw, value, label)
}

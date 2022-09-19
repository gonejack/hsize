package main

import (
	"bufio"
	"fmt"
	"github.com/alecthomas/kong"
	"os"
	"strings"
)

type cmd struct {
	opts struct {
		Args      []string `arg:"" optional:""`
		Precision uint     `short:"p" default:"2" help:"Precision."`
		Five      bool     `short:"5" help:"Give me five."`
		About     bool     `help:"About"`
	}
}

func (c *cmd) run() {
	kong.Parse(&c.opts,
		kong.Name("hsize"),
		kong.Description("hsize 123 383764 <OR> echo 19129219219129119 | hsize"),
		kong.UsageOnError(),
	)
	if c.opts.Five {
		fmt.Print("ヘ( ^o^)ノ＼(^_^ )")
		return
	}
	if c.opts.About {
		fmt.Println("Visit https://github.com/gonejack/hsize")
		return
	}
	if c.opts.Precision > 0 {
		prec = int(c.opts.Precision)
	}
	if len(c.opts.Args) > 0 {
		for _, arg := range c.opts.Args {
			c.parse(arg)
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			c.parse(scanner.Text())
		}
		if scanner.Err() != nil {
			exitf("error reading stdin: %s", scanner.Err())
		}
	}
}
func (c *cmd) parse(raw string) {
	size, err := NewSizeNum().From(strings.TrimSpace(raw))
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

func exitf(format string, a ...interface{}) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(format, a...))
	os.Exit(-2)
}
func main() {
	new(cmd).run()
}

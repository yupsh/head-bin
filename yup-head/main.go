package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/head"
)

const (
	flagLines = "lines"
	flagBytes = "bytes"
	flagQuiet = "quiet"
)

func main() {
	app := &cli.App{
		Name:  "head",
		Usage: "output the first part of files",
		UsageText: `head [OPTIONS] [FILE...]

   Print the first 10 lines of each FILE to standard output.
   With more than one FILE, precede each with a header giving the file name.
   With no FILE, or when FILE is -, read standard input.`,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    flagLines,
				Aliases: []string{"n"},
				Usage:   "print the first NUM lines instead of the first 10",
				Value:   10,
			},
			&cli.IntFlag{
				Name:    flagBytes,
				Aliases: []string{"c"},
				Usage:   "print the first NUM bytes of each file",
			},
			&cli.BoolFlag{
				Name:    flagQuiet,
				Aliases: []string{"q", "silent"},
				Usage:   "never print headers giving file names",
			},
		},
		Action: action,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "head: %v\n", err)
		os.Exit(1)
	}
}

func action(c *cli.Context) error {
	var params []any

	// Add file arguments (or none for stdin)
	for i := 0; i < c.NArg(); i++ {
		params = append(params, yup.File(c.Args().Get(i)))
	}

	// Add flags based on CLI options
	if c.IsSet(flagLines) {
		params = append(params, LineCount(c.Int(flagLines)))
	}
	if c.IsSet(flagBytes) {
		params = append(params, ByteCount(c.Int(flagBytes)))
	}
	if c.Bool(flagQuiet) {
		params = append(params, Quiet)
	}

	// Create and execute the head command
	cmd := Head(params...)
	return yup.Run(cmd)
}

package commands

import (
	"errors"
	"github.com/urfave/cli/v2"
)

var conversionArgs = &struct {
	binary   bool
	octal    bool
	decimal  bool
	hex      bool
	upperHex bool
}{}

var conversionFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:        "binary",
		Aliases:     []string{"b"},
		Usage:       "Convert to binary system",
		Destination: &conversionArgs.binary,
	},
	&cli.BoolFlag{
		Name:        "octal",
		Aliases:     []string{"o"},
		Usage:       "Convert to octal system",
		Destination: &conversionArgs.octal,
	},
	&cli.BoolFlag{
		Name:        "decimal",
		Aliases:     []string{"d"},
		Usage:       "Convert to decimal system",
		Destination: &conversionArgs.decimal,
	},
	&cli.BoolFlag{
		Name:        "hex",
		Usage:       "Convert to hexadecimal system",
		Destination: &conversionArgs.hex,
	},
	&cli.BoolFlag{
		Name:        "Hex",
		Aliases:     []string{"H"},
		Usage:       "Convert to hexadecimal system (Uppercase)",
		Destination: &conversionArgs.upperHex,
	},
}

var Conversion = &cli.Command{
	Name:    "convert",
	Aliases: []string{"c"},
	Usage:   "convert number system",
	UsageText: `hex [command] -[flag] [args...]

hex c -b 1988
hex c -o 1988
hex c -d 1988
hex c -hex 1988
hex c -H 1988
hex c 0b110
hex c 06
hex c 0x1e
`,
	ArgsUsage: "encode [ids...]",
	Action: func(c *cli.Context) error {
		if c.NArg() == 0 {
			return errors.New("none number is present")
		}

		for _, v := range c.Args().Slice() {
			number, err := parseInt(v)
			if err != nil {
				return err
			}

			if conversionArgs.binary {
				printBinary(number)
			} else if conversionArgs.octal {
				printOctal(number)
			} else if conversionArgs.decimal {
				printDecimal(number)
			} else if conversionArgs.hex {
				printHex(number)
			} else {
				printUpperHex(number)
			}
		}

		return nil
	},
	Flags: conversionFlags,
}

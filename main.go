package main

import (
	"github.com/gg-tools/hex/commands"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Name = "Hex"
	app.Compiled = time.Now()
	app.Usage = "Hex - convert between number systems"
	app.UsageText = `hex [command] -[flag] [args...]

hex -b 1988
hex -o 1988
hex -H 1988
hex 0b110
hex 06
hex 0x1e`
	app.Commands = []*cli.Command{commands.Conversion}
	app.Action = commands.Conversion.Action
	app.Flags = commands.Conversion.Flags

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

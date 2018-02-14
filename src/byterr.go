package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "Byterr Go Version"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Johnny Santos",
			Email: "johnnysantos@protonmail.com",
		},
	}
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:    "hex",
			Usage:   "string to hex",
			Aliases: []string{"sh"},
			Action:  stringToHex,
		},
		{
			Name:    "xeh",
			Usage:   "hex to string",
			Aliases: []string{"hs"},
			Action:  hexToString,
		},
	}

	app.RunAndExitOnError()
}

func stringToHex(c *cli.Context) error {
	if c.NArg() > 0 {
		firstArg := c.Args().First()
		encodeStr := hex.EncodeToString([]byte(firstArg))

		fmt.Printf("String: %q => Hex: 0x%s\n", firstArg, encodeStr)
	}
	return nil
}

func hexToString(c *cli.Context) error {
	if c.NArg() > 0 {
		firstArg := strings.TrimLeft(c.Args().First(), "0x")
		encodeStr, err := hex.DecodeString(firstArg)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Hex: 0x%s => String: %q\n", firstArg, encodeStr)
	}
	return nil
}

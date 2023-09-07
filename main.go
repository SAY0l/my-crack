package main

import (
	"os"

	"github.com/sayol/my_crack/cmd"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "My-Crack"
	app.Author = "sayol"
	app.Email = "github@sayol.com"
	app.Version = "1.1"
	app.Usage = "Weak password crack"
	app.Commands = []cli.Command{cmd.Scan}
	app.Flags = append(app.Flags, cmd.Scan.Flags...)
	err := app.Run(os.Args)
	_=err
}

package main

import (
	"fmt"
	"io"
	"os"

	"github.com/alecthomas/kong"
	"github.com/mikfreedman/harvey/command"
)

var version = "devel"

var cli struct {
	Strip   command.Strip `cmd:"" name:"strip"`
	Version VersionFlag   `name:"version" short:"v" help:"Print version information and quit"`
}

type VersionFlag string

func (v VersionFlag) Decode(ctx *kong.DecodeContext) error { return nil }
func (v VersionFlag) IsBool() bool                         { return true }
func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Println(version)
	app.Exit(0)
	return nil
}

func main() {
	ctx := kong.Parse(&cli,
		kong.Name("harvey"),
		kong.Description("Cleanup Markdown"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
	)
	ctx.BindTo(os.Stdout, (*io.Writer)(nil))
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}

package cli

import (
	"fmt"
	"os"
	"path"

	log "github.com/Sirupsen/logrus"
	"github.com/bitrise-io/bitrise/colorstring"
	"github.com/codegangsta/cli"
)

var (
	// IsCIMode ...
	IsCIMode = false
)

func before(c *cli.Context) error {
	// Log level
	level, err := log.ParseLevel(c.String(LogLevelKey))
	if err != nil {
		return err
	}

	if err := os.Setenv(LogLevelEnvKey, level.String()); err != nil {
		log.Fatal("Failed to set log level env:", err)
	}
	log.SetLevel(level)

	// CI
	//  if CI mode indicated make sure we set the related env
	//  so all other tools we use will also get it
	if c.Bool(CIKey) {
		if err := os.Setenv("CI", "true"); err != nil {
			return err
		}
		IsCIMode = true
		log.Info(colorstring.Yellow("bitrise runs in CI mode"))
	}
	return nil
}

func printVersion(c *cli.Context) {
	fmt.Fprintf(c.App.Writer, "%v\n", c.App.Version)
}

// Run ...
func Run() {
	// Parse cl
	cli.VersionPrinter = printVersion

	app := cli.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Usage = "Bitrise Automations Workflow Runner"
	app.Version = "0.9.8"

	app.Author = ""
	app.Email = ""

	app.Before = before

	app.Flags = flags
	app.Commands = commands

	if err := app.Run(os.Args); err != nil {
		log.Fatal("Finished with Error:", err)
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kingpin"
	aw "github.com/deanishe/awgo"
)

var (
	// Icons
	updateAvailable = &aw.Icon{Value: "icons/update-available.png"}

	dateOfBirth = os.Getenv("Date of Birth (DD/MM/YYYY)")

	// Kingpin and script options
	app *kingpin.Application

	// Application commands
	searchCmd *kingpin.CmdClause

	query string

	// Workflow stuff
	wf *aw.Workflow
)

// Mostly sets up kingpin commands
func init() {
	wf = aw.New()

	app = kingpin.New("birthday", "See birthday")

	// Commands using query
	searchCmd = app.Command("search", "See birthday.").Alias("s")

	// Common options
	for _, cmd := range []*kingpin.CmdClause{
		searchCmd,
	} {
		cmd.Flag("query", "Search query.").Short('q').StringVar(&query)
	}
}

func run() {
	var err error

	cmd, err := app.Parse(wf.Args())
	if err != nil {
		wf.FatalError(err)
	}

	switch cmd {
	case searchCmd.FullCommand():
		err = doSearch()
	default:
		err = fmt.Errorf("Uknown command: %s", cmd)
	}

	if err != nil {
		wf.FatalError(err)
	}
}

func main() {
	wf.Run(run)
}

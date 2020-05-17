package runner

import (
	"github.com/urfave/cli/v2"

	command "gitlab.com/alienspaces/holyragingmages/common/cli"
	"gitlab.com/alienspaces/holyragingmages/common/prepare"
	"gitlab.com/alienspaces/holyragingmages/common/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/common/type/preparer"
	"gitlab.com/alienspaces/holyragingmages/service/spell/internal/model"
)

// Runner -
type Runner struct {
	command.Runner
}

// NewRunner -
func NewRunner() *Runner {

	r := Runner{}

	// https://github.com/urfave/cli/blob/master/docs/v2/manual.md
	r.App = &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "test",
				Aliases: []string{"t"},
				Usage:   "Runs the test command",
				Action:  r.TestCommand,
			},
		},
	}

	r.PreparerFunc = r.Preparer
	r.ModellerFunc = r.Modeller

	return &r
}

// TestCommand -
func (rnr *Runner) TestCommand(c *cli.Context) error {

	rnr.Log.Info("** Spell Test Command **")

	return nil
}

// Preparer -
func (rnr *Runner) Preparer() (preparer.Preparer, error) {

	rnr.Log.Info("** Spell Preparer **")

	p, err := prepare.NewPrepare(rnr.Log)
	if err != nil {
		rnr.Log.Warn("Failed new preparer >%v<", err)
		return nil, err
	}

	return p, nil
}

// Modeller -
func (rnr *Runner) Modeller() (modeller.Modeller, error) {

	rnr.Log.Info("** Spell Model **")

	m, err := model.NewModel(rnr.Config, rnr.Log, rnr.Store)
	if err != nil {
		rnr.Log.Warn("Failed new model >%v<", err)
		return nil, err
	}

	return m, nil
}

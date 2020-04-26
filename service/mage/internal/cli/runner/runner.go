package runner

import (
	"gitlab.com/alienspaces/holyragingmages/common/cli"
	"gitlab.com/alienspaces/holyragingmages/common/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/service/mage/internal/model"
)

// Runner -
type Runner struct {
	cli.Runner
}

// NewRunner -
func NewRunner() *Runner {

	r := Runner{}

	r.ModellerFunc = r.Modeller

	return &r
}

// Modeller -
func (rnr *Runner) Modeller() (modeller.Modeller, error) {

	rnr.Log.Info("** Template Model **")

	m, err := model.NewModel(rnr.Config, rnr.Log, rnr.Store)
	if err != nil {
		rnr.Log.Warn("Failed new model >%v<", err)
		return nil, err
	}

	return m, nil
}

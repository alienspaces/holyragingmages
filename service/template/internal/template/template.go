package template

import (
	"gitlab.com/alienspaces/holyragingmages/common/service"
)

// Runner -
type Runner struct{}

// Run -
func (r *Runner) Run(c service.Configurer, l service.Logger, d service.Storer, args map[string]interface{}) error {

	l.Printf("** Running **")

	return nil
}

// Configuration -
func (r *Runner) Configuration() []string {
	return nil
}

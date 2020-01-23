package template

import (
	"gitlab.com/alienspaces/holyragingmages/common/service"
)

// Runner -
type Runner struct {
	service.APIRunner
}

// Run -
func (r *Runner) Run(args map[string]interface{}) error {

	r.Log.Printf("** Running **")

	return nil
}

package runnable

import (
	"gitlab.com/alienspaces/holyragingmages/common/configurer"
	"gitlab.com/alienspaces/holyragingmages/common/logger"
	"gitlab.com/alienspaces/holyragingmages/common/storer"
)

// Runnable -
type Runnable interface {
	Init(c configurer.Configurer, l logger.Logger, s storer.Storer) error
	Run(args map[string]interface{}) error
}

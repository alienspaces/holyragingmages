package runnable

import (
	"gitlab.com/alienspaces/holyragingmages/server/core/type/configurer"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/logger"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/storer"
)

// Runnable -
type Runnable interface {
	Init(c configurer.Configurer, l logger.Logger, s storer.Storer) error
	Run(args map[string]interface{}) error
}

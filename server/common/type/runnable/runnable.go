package runnable

import (
	"gitlab.com/alienspaces/holyragingmages/server/common/type/configurer"
	"gitlab.com/alienspaces/holyragingmages/server/common/type/logger"
	"gitlab.com/alienspaces/holyragingmages/server/common/type/storer"
)

// Runnable -
type Runnable interface {
	Init(c configurer.Configurer, l logger.Logger, s storer.Storer) error
	Run(args map[string]interface{}) error
}

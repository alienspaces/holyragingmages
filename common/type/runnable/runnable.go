package runnable

import (
	"gitlab.com/alienspaces/holyragingmages/common/type/configurer"
	"gitlab.com/alienspaces/holyragingmages/common/type/logger"
	"gitlab.com/alienspaces/holyragingmages/common/type/preparer"
	"gitlab.com/alienspaces/holyragingmages/common/type/storer"
)

// Runnable -
type Runnable interface {
	Init(c configurer.Configurer, l logger.Logger, s storer.Storer, p preparer.Preparer) error
	Run(args map[string]interface{}) error
}

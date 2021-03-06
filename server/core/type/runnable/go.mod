module gitlab.com/alienspaces/holyragingmages/server/core/repository

go 1.15

require (
	gitlab.com/alienspaces/holyragingmages/server/core/type/configurer v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/core/type/logger v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/core/type/preparable v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/core/type/preparer v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/core/type/storer v1.0.0
)

replace (
	gitlab.com/alienspaces/holyragingmages/server/core/config => ../../core/config
	gitlab.com/alienspaces/holyragingmages/server/core/log => ../../core/log
	gitlab.com/alienspaces/holyragingmages/server/core/store => ../../core/store
	gitlab.com/alienspaces/holyragingmages/server/core/type/configurer => ../../../core/type/configurer
	gitlab.com/alienspaces/holyragingmages/server/core/type/logger => ../../../core/type/logger
	gitlab.com/alienspaces/holyragingmages/server/core/type/preparable => ../../../core/type/preparable
	gitlab.com/alienspaces/holyragingmages/server/core/type/preparer => ../../../core/type/preparer
	gitlab.com/alienspaces/holyragingmages/server/core/type/storer => ../../../core/type/storer
)

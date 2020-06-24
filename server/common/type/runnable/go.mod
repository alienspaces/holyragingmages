module gitlab.com/alienspaces/holyragingmages/server/common/repository

go 1.13

require (
	gitlab.com/alienspaces/holyragingmages/server/common/type/configurer v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/common/type/logger v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/common/type/preparable v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/common/type/preparer v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/common/type/storer v1.0.0
)

replace (
	gitlab.com/alienspaces/holyragingmages/server/common/config => ../../common/config
	gitlab.com/alienspaces/holyragingmages/server/common/log => ../../common/log
	gitlab.com/alienspaces/holyragingmages/server/common/store => ../../common/store
	gitlab.com/alienspaces/holyragingmages/server/common/type/configurer => ../../../common/type/configurer
	gitlab.com/alienspaces/holyragingmages/server/common/type/logger => ../../../common/type/logger
	gitlab.com/alienspaces/holyragingmages/server/common/type/preparable => ../../../common/type/preparable
	gitlab.com/alienspaces/holyragingmages/server/common/type/preparer => ../../../common/type/preparer
	gitlab.com/alienspaces/holyragingmages/server/common/type/storer => ../../../common/type/storer
)

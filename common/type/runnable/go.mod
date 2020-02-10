module gitlab.com/alienspaces/holyragingmages/common/repository

go 1.13

require (
	gitlab.com/alienspaces/holyragingmages/common/type/configurer v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/type/logger v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/type/storer v1.0.0
)

replace (
	gitlab.com/alienspaces/holyragingmages/common/config => ../../common/config
	gitlab.com/alienspaces/holyragingmages/common/log => ../../common/log
	gitlab.com/alienspaces/holyragingmages/common/store => ../../common/store
	gitlab.com/alienspaces/holyragingmages/common/type/configurer => ../../../common/type/configurer
	gitlab.com/alienspaces/holyragingmages/common/type/logger => ../../../common/type/logger
	gitlab.com/alienspaces/holyragingmages/common/type/storer => ../../../common/type/storer
)

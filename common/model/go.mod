module gitlab.com/alienspaces/holyragingmages/common/model

go 1.13

require (
	github.com/jmoiron/sqlx v1.2.0

	gitlab.com/alienspaces/holyragingmages/common/configurer v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/logger v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/preparer v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/storer v1.0.0
)

replace (
	gitlab.com/alienspaces/holyragingmages/common/config => ../../common/config
	gitlab.com/alienspaces/holyragingmages/common/configurer => ../../common/configurer
	gitlab.com/alienspaces/holyragingmages/common/log => ../../common/log
	gitlab.com/alienspaces/holyragingmages/common/logger => ../../common/logger
	gitlab.com/alienspaces/holyragingmages/common/preparable => ../../common/preparable
	gitlab.com/alienspaces/holyragingmages/common/preparer => ../../common/preparer
	gitlab.com/alienspaces/holyragingmages/common/store => ../../common/store
	gitlab.com/alienspaces/holyragingmages/common/storer => ../../common/storer
)

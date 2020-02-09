module gitlab.com/alienspaces/holyragingmages/common/repository

go 1.13

require (
	github.com/google/uuid v1.1.1
	github.com/jmoiron/sqlx v1.2.0
	gitlab.com/alienspaces/holyragingmages/common/logger v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/preparer v1.0.0
)

replace (
	gitlab.com/alienspaces/holyragingmages/common/log => ../../common/log
	gitlab.com/alienspaces/holyragingmages/common/logger => ../../common/logger
	gitlab.com/alienspaces/holyragingmages/common/preparable => ../../common/preparable
	gitlab.com/alienspaces/holyragingmages/common/preparer => ../../common/preparer
)

module gitlab.com/alienspaces/holyragingmages/common/repository

go 1.13

require (
	github.com/jmoiron/sqlx v1.2.0
	gitlab.com/alienspaces/holyragingmages/common/type/logger v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/type/preparable v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/type/preparer v1.0.0
)

replace (
	gitlab.com/alienspaces/holyragingmages/common/type/logger => ../../common/type/logger
	gitlab.com/alienspaces/holyragingmages/common/type/preparable => ../../common/type/preparable
	gitlab.com/alienspaces/holyragingmages/common/type/preparer => ../../common/type/preparer
)

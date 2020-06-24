module gitlab.com/alienspaces/holyragingmages/server/common/repository

go 1.13

require (
	github.com/google/uuid v1.1.1
	github.com/jmoiron/sqlx v1.2.0
	gitlab.com/alienspaces/holyragingmages/server/common/type/logger v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/common/type/preparer v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/common/type/repositor v1.0.0
)

replace (
	gitlab.com/alienspaces/holyragingmages/server/common/log => ../../common/log
	gitlab.com/alienspaces/holyragingmages/server/common/type/logger => ../../common/type/logger
	gitlab.com/alienspaces/holyragingmages/server/common/type/preparable => ../../common/type/preparable
	gitlab.com/alienspaces/holyragingmages/server/common/type/preparer => ../../common/type/preparer
	gitlab.com/alienspaces/holyragingmages/server/common/type/repositor => ../../common/type/repositor
)

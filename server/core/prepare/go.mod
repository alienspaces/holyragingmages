module gitlab.com/alienspaces/holyragingmages/server/core/repository

go 1.13

require (
	github.com/jmoiron/sqlx v1.2.0
	gitlab.com/alienspaces/holyragingmages/server/core/type/logger v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/core/type/preparable v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/core/type/preparer v1.0.0
)

replace (
	gitlab.com/alienspaces/holyragingmages/server/core/type/logger => ../../core/type/logger
	gitlab.com/alienspaces/holyragingmages/server/core/type/preparable => ../../core/type/preparable
	gitlab.com/alienspaces/holyragingmages/server/core/type/preparer => ../../core/type/preparer
)
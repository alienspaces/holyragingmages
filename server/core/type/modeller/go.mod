module gitlab.com/alienspaces/holyragingmages/server/core/type/modeller

go 1.15

require (
	github.com/jmoiron/sqlx v1.2.0
	gitlab.com/alienspaces/holyragingmages/server/core/type/preparer v1.0.0
)

replace (
	gitlab.com/alienspaces/holyragingmages/server/core/type/preparable => ../../../core/type/preparable
	gitlab.com/alienspaces/holyragingmages/server/core/type/preparer => ../../../core/type/preparer
)

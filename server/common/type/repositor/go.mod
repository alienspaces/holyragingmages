module gitlab.com/alienspaces/holyragingmages/server/common/type/repositor

go 1.13

require (
	github.com/jmoiron/sqlx v1.2.0
	gitlab.com/alienspaces/holyragingmages/server/common/type/preparer v1.0.0
)

replace (
	gitlab.com/alienspaces/holyragingmages/server/common/type/preparable => ../../../common/type/preparable
	gitlab.com/alienspaces/holyragingmages/server/common/type/preparer => ../../../common/type/preparer
)

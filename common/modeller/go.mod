module gitlab.com/alienspaces/holyragingmages/common/modeller

go 1.13

require (
	github.com/jmoiron/sqlx v1.2.0
	gitlab.com/alienspaces/holyragingmages/common/preparer v1.0.0
)

replace (
	gitlab.com/alienspaces/holyragingmages/common/preparable => ../../common/preparable
	gitlab.com/alienspaces/holyragingmages/common/preparer => ../../common/preparer
)

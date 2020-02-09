module gitlab.com/alienspaces/holyragingmages/common/repository

go 1.13

require (
	github.com/jmoiron/sqlx v1.2.0

	gitlab.com/alienspaces/holyragingmages/common/logger v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/preparable v1.0.0
	google.golang.org/appengine v1.6.5 // indirect
)

replace (
	gitlab.com/alienspaces/holyragingmages/common/logger => ../../common/logger
	gitlab.com/alienspaces/holyragingmages/common/preparable => ../../common/preparable
)

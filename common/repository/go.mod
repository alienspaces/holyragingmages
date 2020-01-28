module gitlab.com/alienspaces/holyragingmages/common/repository

go 1.13

require (
	github.com/google/uuid v1.1.1
	github.com/jmoiron/sqlx v1.2.0
	github.com/lib/pq v1.2.0
	google.golang.org/appengine v1.6.5 // indirect
)

replace (
	gitlab.com/alienspaces/holyragingmages/common/database => ../../common/database
	gitlab.com/alienspaces/holyragingmages/common/env => ../../common/env
	gitlab.com/alienspaces/holyragingmages/common/logger => ../../common/logger
)

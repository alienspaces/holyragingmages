module gitlab.com/alienspaces/holyragingmages/common/service

go 1.13

require (
	github.com/jmoiron/sqlx v1.2.0
	github.com/stretchr/testify v1.4.0

	gitlab.com/alienspaces/holyragingmages/common/database v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/env v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/logger v1.0.0
)

replace (
	gitlab.com/alienspaces/holyragingmages/common/database => ../../common/database
	gitlab.com/alienspaces/holyragingmages/common/env => ../../common/env
	gitlab.com/alienspaces/holyragingmages/common/logger => ../../common/logger
)

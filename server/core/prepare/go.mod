module gitlab.com/alienspaces/holyragingmages/server/core/prepare

go 1.15

require (
	github.com/jmoiron/sqlx v1.2.0
	github.com/stretchr/testify v1.6.1

	gitlab.com/alienspaces/holyragingmages/server/core/config v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/core/log v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/core/repository v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/core/store v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/core/type/configurer v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/core/type/logger v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/core/type/preparable v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/core/type/preparer v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/core/type/storer v1.0.0
)

replace (
	gitlab.com/alienspaces/holyragingmages/server/core/config => ../../core/config
	gitlab.com/alienspaces/holyragingmages/server/core/log => ../../core/log
	gitlab.com/alienspaces/holyragingmages/server/core/prepare => ../../core/prepare
	gitlab.com/alienspaces/holyragingmages/server/core/repository => ../../core/repository
	gitlab.com/alienspaces/holyragingmages/server/core/store => ../../core/store
	gitlab.com/alienspaces/holyragingmages/server/core/type/configurer => ../../core/type/configurer
	gitlab.com/alienspaces/holyragingmages/server/core/type/logger => ../../core/type/logger
	gitlab.com/alienspaces/holyragingmages/server/core/type/preparable => ../../core/type/preparable
	gitlab.com/alienspaces/holyragingmages/server/core/type/preparer => ../../core/type/preparer
	gitlab.com/alienspaces/holyragingmages/server/core/type/repositor => ../../core/type/repositor
	gitlab.com/alienspaces/holyragingmages/server/core/type/storer => ../../core/type/storer
)

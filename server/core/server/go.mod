module gitlab.com/alienspaces/holyragingmages/server/core/server

go 1.13

require (
	github.com/google/uuid v1.1.1
	github.com/jmoiron/sqlx v1.2.0
	github.com/julienschmidt/httprouter v1.3.0
	github.com/rs/cors v1.7.0
	github.com/stretchr/testify v1.4.0
	github.com/xeipuuv/gojsonschema v1.2.0

	gitlab.com/alienspaces/holyragingmages/server/core/config v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/core/log v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/core/prepare v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/core/store v1.0.0

	gitlab.com/alienspaces/holyragingmages/server/core/type/configurer v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/core/type/logger v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/core/type/modeller v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/core/type/payloader v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/core/type/preparable v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/core/type/preparer v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/core/type/runnable v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/core/type/storer v1.0.0
)

replace (
	gitlab.com/alienspaces/holyragingmages/server/core/config => ../../core/config
	gitlab.com/alienspaces/holyragingmages/server/core/log => ../../core/log
	gitlab.com/alienspaces/holyragingmages/server/core/model => ../../core/model
	gitlab.com/alienspaces/holyragingmages/server/core/prepare => ../../core/prepare
	gitlab.com/alienspaces/holyragingmages/server/core/store => ../../core/store
	gitlab.com/alienspaces/holyragingmages/server/core/type/configurer => ../../core/type/configurer
	gitlab.com/alienspaces/holyragingmages/server/core/type/logger => ../../core/type/logger
	gitlab.com/alienspaces/holyragingmages/server/core/type/modeller => ../../core/type/modeller
	gitlab.com/alienspaces/holyragingmages/server/core/type/payloader => ../../core/type/payloader
	gitlab.com/alienspaces/holyragingmages/server/core/type/preparable => ../../core/type/preparable
	gitlab.com/alienspaces/holyragingmages/server/core/type/preparer => ../../core/type/preparer
	gitlab.com/alienspaces/holyragingmages/server/core/type/runnable => ../../core/type/runnable
	gitlab.com/alienspaces/holyragingmages/server/core/type/storer => ../../core/type/storer
)
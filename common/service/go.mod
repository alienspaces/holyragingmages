module gitlab.com/alienspaces/holyragingmages/common/service

go 1.13

require (
	github.com/jmoiron/sqlx v1.2.0
	github.com/julienschmidt/httprouter v1.3.0
	github.com/stretchr/testify v1.4.0
	github.com/xeipuuv/gojsonschema v1.2.0

	gitlab.com/alienspaces/holyragingmages/common/config v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/logger v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/store v1.0.0
)

replace (
	gitlab.com/alienspaces/holyragingmages/common/config => ../../common/config
	gitlab.com/alienspaces/holyragingmages/common/logger => ../../common/logger
	gitlab.com/alienspaces/holyragingmages/common/store => ../../common/store
)

module gitlab.com/alienspaces/holyragingmages/common/service

go 1.13

require (
	github.com/jmoiron/sqlx v1.2.0
	github.com/julienschmidt/httprouter v1.3.0
	github.com/stretchr/testify v1.4.0
	github.com/xeipuuv/gojsonschema v1.2.0

	gitlab.com/alienspaces/holyragingmages/common/config v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/configurer v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/log v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/logger v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/modeller v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/preparable v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/prepare v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/preparer v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/runnable v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/store v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/storer v1.0.0
)

replace (
	gitlab.com/alienspaces/holyragingmages/common/config => ../../common/config
	gitlab.com/alienspaces/holyragingmages/common/configurer => ../../common/configurer
	gitlab.com/alienspaces/holyragingmages/common/log => ../../common/log
	gitlab.com/alienspaces/holyragingmages/common/logger => ../../common/logger
	gitlab.com/alienspaces/holyragingmages/common/model => ../../common/model
	gitlab.com/alienspaces/holyragingmages/common/modeller => ../../common/modeller
	gitlab.com/alienspaces/holyragingmages/common/preparable => ../../common/preparable
	gitlab.com/alienspaces/holyragingmages/common/prepare => ../../common/prepare
	gitlab.com/alienspaces/holyragingmages/common/preparer => ../../common/preparer
	gitlab.com/alienspaces/holyragingmages/common/runnable => ../../common/runnable
	gitlab.com/alienspaces/holyragingmages/common/store => ../../common/store
	gitlab.com/alienspaces/holyragingmages/common/storer => ../../common/storer
)

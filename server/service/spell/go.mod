module gitlab.com/alienspaces/holyragingmages/server/service/spell

go 1.13

require (
	github.com/google/uuid v1.1.1
	github.com/jmoiron/sqlx v1.2.0
	github.com/julienschmidt/httprouter v1.3.0
	github.com/stretchr/testify v1.4.0
	github.com/urfave/cli/v2 v2.2.0

	gitlab.com/alienspaces/holyragingmages/server/common/cli v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/common/config v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/common/harness v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/common/log v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/common/model v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/common/payload v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/common/prepare v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/common/repository v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/common/server v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/common/store v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/common/type/configurer v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/common/type/logger v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/common/type/modeller v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/common/type/payloader v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/common/type/preparer v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/common/type/repositor v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/common/type/runnable v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/common/type/storer v1.0.0
)

replace (
	gitlab.com/alienspaces/holyragingmages/server/common/cli => ../../common/cli
	gitlab.com/alienspaces/holyragingmages/server/common/config => ../../common/config
	gitlab.com/alienspaces/holyragingmages/server/common/harness => ../../common/harness
	gitlab.com/alienspaces/holyragingmages/server/common/log => ../../common/log
	gitlab.com/alienspaces/holyragingmages/server/common/model => ../../common/model
	gitlab.com/alienspaces/holyragingmages/server/common/payload => ../../common/payload
	gitlab.com/alienspaces/holyragingmages/server/common/prepare => ../../common/prepare
	gitlab.com/alienspaces/holyragingmages/server/common/repository => ../../common/repository
	gitlab.com/alienspaces/holyragingmages/server/common/server => ../../common/server
	gitlab.com/alienspaces/holyragingmages/server/common/store => ../../common/store
	gitlab.com/alienspaces/holyragingmages/server/common/type/configurer => ../../common/type/configurer
	gitlab.com/alienspaces/holyragingmages/server/common/type/logger => ../../common/type/logger
	gitlab.com/alienspaces/holyragingmages/server/common/type/modeller => ../../common/type/modeller
	gitlab.com/alienspaces/holyragingmages/server/common/type/payloader => ../../common/type/payloader
	gitlab.com/alienspaces/holyragingmages/server/common/type/preparable => ../../common/type/preparable
	gitlab.com/alienspaces/holyragingmages/server/common/type/preparer => ../../common/type/preparer
	gitlab.com/alienspaces/holyragingmages/server/common/type/repositor => ../../common/type/repositor
	gitlab.com/alienspaces/holyragingmages/server/common/type/runnable => ../../common/type/runnable
	gitlab.com/alienspaces/holyragingmages/server/common/type/storer => ../../common/type/storer
	gitlab.com/alienspaces/holyragingmages/server/service/spell => ../../service/spell
)

module gitlab.com/alienspaces/holyragingmages/server/common/client

go 1.13

require (
	github.com/stretchr/testify v1.4.0

	gitlab.com/alienspaces/holyragingmages/server/common/config v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/common/log v1.0.0

	gitlab.com/alienspaces/holyragingmages/server/common/type/configurer v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/common/type/logger v1.0.0
)

replace (
	gitlab.com/alienspaces/holyragingmages/server/common/config => ../../common/config
	gitlab.com/alienspaces/holyragingmages/server/common/log => ../../common/log
	gitlab.com/alienspaces/holyragingmages/server/common/model => ../../common/model
	gitlab.com/alienspaces/holyragingmages/server/common/prepare => ../../common/prepare
	gitlab.com/alienspaces/holyragingmages/server/common/store => ../../common/store
	gitlab.com/alienspaces/holyragingmages/server/common/type/configurer => ../../common/type/configurer
	gitlab.com/alienspaces/holyragingmages/server/common/type/logger => ../../common/type/logger
	gitlab.com/alienspaces/holyragingmages/server/common/type/modeller => ../../common/type/modeller
	gitlab.com/alienspaces/holyragingmages/server/common/type/payloader => ../../common/type/payloader
	gitlab.com/alienspaces/holyragingmages/server/common/type/preparable => ../../common/type/preparable
	gitlab.com/alienspaces/holyragingmages/server/common/type/preparer => ../../common/type/preparer
	gitlab.com/alienspaces/holyragingmages/server/common/type/runnable => ../../common/type/runnable
	gitlab.com/alienspaces/holyragingmages/server/common/type/storer => ../../common/type/storer
)

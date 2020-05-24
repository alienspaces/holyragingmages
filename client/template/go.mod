module gitlab.com/alienspaces/holyragingmages/client/template

go 1.13

require gitlab.com/alienspaces/holyragingmages/common/client v0.0.0

replace (
	gitlab.com/alienspaces/holyragingmages/common/client => ../../common/client
	gitlab.com/alienspaces/holyragingmages/common/config => ../../common/config
	gitlab.com/alienspaces/holyragingmages/common/log => ../../common/log
	gitlab.com/alienspaces/holyragingmages/common/model => ../../common/model
	gitlab.com/alienspaces/holyragingmages/common/prepare => ../../common/prepare
	gitlab.com/alienspaces/holyragingmages/common/store => ../../common/store
	gitlab.com/alienspaces/holyragingmages/common/type/configurer => ../../common/type/configurer
	gitlab.com/alienspaces/holyragingmages/common/type/logger => ../../common/type/logger
	gitlab.com/alienspaces/holyragingmages/common/type/modeller => ../../common/type/modeller
	gitlab.com/alienspaces/holyragingmages/common/type/payloader => ../../common/type/payloader
	gitlab.com/alienspaces/holyragingmages/common/type/preparable => ../../common/type/preparable
	gitlab.com/alienspaces/holyragingmages/common/type/preparer => ../../common/type/preparer
	gitlab.com/alienspaces/holyragingmages/common/type/runnable => ../../common/type/runnable
	gitlab.com/alienspaces/holyragingmages/common/type/storer => ../../common/type/storer
)

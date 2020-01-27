module gitlab.com/alienspaces/holyragingmages/service/template

go 1.13

require (
	github.com/julienschmidt/httprouter v1.3.0
	github.com/stretchr/testify v1.4.0
	gitlab.com/alienspaces/holyragingmages/common/database v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/env v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/logger v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/service v1.0.0
)

replace (
	gitlab.com/alienspaces/holyragingmages/common/database => ../../common/database
	gitlab.com/alienspaces/holyragingmages/common/env => ../../common/env
	gitlab.com/alienspaces/holyragingmages/common/logger => ../../common/logger
	gitlab.com/alienspaces/holyragingmages/common/service => ../../common/service
)

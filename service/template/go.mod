module gitlab.com/alienspaces/holyragingmages/service/template

go 1.13

require (
	github.com/corpix/uarand v0.1.1 // indirect
	github.com/icrowley/fake v0.0.0-20180203215853-4178557ae428 // indirect
	github.com/jmoiron/sqlx v1.2.0
	github.com/julienschmidt/httprouter v1.3.0
	github.com/pkg/errors v0.9.1 // indirect
	github.com/stretchr/testify v1.4.0
	gitlab.com/alienspaces/holyragingmages/common/config v1.0.0

	gitlab.com/alienspaces/holyragingmages/common/database v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/env v0.0.0-20200128202053-19b68cc811c8
	gitlab.com/alienspaces/holyragingmages/common/logger v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/model v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/repository v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/service v1.0.0
)

replace (
	gitlab.com/alienspaces/holyragingmages/common/config => ../../common/config
	gitlab.com/alienspaces/holyragingmages/common/database => ../../common/database
	gitlab.com/alienspaces/holyragingmages/common/logger => ../../common/logger
	gitlab.com/alienspaces/holyragingmages/common/model => ../../common/model
	gitlab.com/alienspaces/holyragingmages/common/repository => ../../common/repository
	gitlab.com/alienspaces/holyragingmages/common/service => ../../common/service
)

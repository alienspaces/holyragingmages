module gitlab.com/alienspaces/holyragingmages/common/client

go 1.13

require (
	github.com/brianvoe/gofakeit v3.18.0+incompatible
	github.com/google/uuid v1.1.1
	github.com/jmoiron/sqlx v1.2.0
	github.com/julienschmidt/httprouter v1.3.0
	github.com/rs/cors v1.7.0
	github.com/stretchr/testify v1.4.0
	github.com/xeipuuv/gojsonschema v1.2.0

	gitlab.com/alienspaces/holyragingmages/common/config v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/log v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/prepare v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/store v1.0.0

	gitlab.com/alienspaces/holyragingmages/common/type/configurer v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/type/logger v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/type/modeller v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/type/payloader v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/type/preparable v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/type/preparer v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/type/runnable v1.0.0
	gitlab.com/alienspaces/holyragingmages/common/type/storer v1.0.0
)

replace (
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
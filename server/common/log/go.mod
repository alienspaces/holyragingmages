module gitlab.com/alienspaces/holyragingmages/server/common/log

go 1.13

require (
	github.com/rs/zerolog v1.18.0
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.4.0

	gitlab.com/alienspaces/holyragingmages/server/common/config v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/common/type/configurer v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/common/type/logger v1.0.0
)

replace (
	gitlab.com/alienspaces/holyragingmages/server/common/config => ../../common/config
	gitlab.com/alienspaces/holyragingmages/server/common/type/configurer => ../../common/type/configurer
	gitlab.com/alienspaces/holyragingmages/server/common/type/logger => ../../common/type/logger
)

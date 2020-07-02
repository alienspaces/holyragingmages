module gitlab.com/alienspaces/holyragingmages/server/core/log

go 1.13

require (
	github.com/rs/zerolog v1.18.0
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.4.0

	gitlab.com/alienspaces/holyragingmages/server/core/config v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/core/type/configurer v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/core/type/logger v1.0.0
)

replace (
	gitlab.com/alienspaces/holyragingmages/server/core/config => ../../core/config
	gitlab.com/alienspaces/holyragingmages/server/core/type/configurer => ../../core/type/configurer
	gitlab.com/alienspaces/holyragingmages/server/core/type/logger => ../../core/type/logger
)

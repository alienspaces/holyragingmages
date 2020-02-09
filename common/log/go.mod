module gitlab.com/alienspaces/holyragingmages/common/log

go 1.13

require (
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.4.0

	gitlab.com/alienspaces/holyragingmages/common/config v0.0.0-00010101000000-000000000000
	gitlab.com/alienspaces/holyragingmages/common/configurer v0.0.0-00010101000000-000000000000
)

replace (
	gitlab.com/alienspaces/holyragingmages/common/config => ../../common/config
	gitlab.com/alienspaces/holyragingmages/common/configurer => ../../common/configurer
)

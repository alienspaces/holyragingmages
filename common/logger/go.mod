module gitlab.com/alienspaces/holyragingmages/common/logger

go 1.13

require (
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.4.0
	gitlab.com/alienspaces/holyragingmages/common/config v0.0.0-00010101000000-000000000000
	gitlab.com/alienspaces/holyragingmages/common/env v0.0.0-20200128202053-19b68cc811c8
)

replace gitlab.com/alienspaces/holyragingmages/common/config => ../../common/config

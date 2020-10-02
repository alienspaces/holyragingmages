module gitlab.com/alienspaces/holyragingmages/server/core/auth

go 1.15

require (
	github.com/brianvoe/gofakeit v3.18.0+incompatible
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/stretchr/testify v1.6.1
	gitlab.com/alienspaces/holyragingmages/server/core/config v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/core/log v0.0.0-00010101000000-000000000000
	gitlab.com/alienspaces/holyragingmages/server/core/type/configurer v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/core/type/logger v1.0.0
)

replace (
	gitlab.com/alienspaces/holyragingmages/server/core/config => ../../core/config
	gitlab.com/alienspaces/holyragingmages/server/core/log => ../../core/log
	gitlab.com/alienspaces/holyragingmages/server/core/model => ../../core/model
	gitlab.com/alienspaces/holyragingmages/server/core/prepare => ../../core/prepare
	gitlab.com/alienspaces/holyragingmages/server/core/store => ../../core/store
	gitlab.com/alienspaces/holyragingmages/server/core/type/configurer => ../../core/type/configurer
	gitlab.com/alienspaces/holyragingmages/server/core/type/logger => ../../core/type/logger
	gitlab.com/alienspaces/holyragingmages/server/core/type/modeller => ../../core/type/modeller
	gitlab.com/alienspaces/holyragingmages/server/core/type/preparable => ../../core/type/preparable
	gitlab.com/alienspaces/holyragingmages/server/core/type/preparer => ../../core/type/preparer
	gitlab.com/alienspaces/holyragingmages/server/core/type/runnable => ../../core/type/runnable
	gitlab.com/alienspaces/holyragingmages/server/core/type/storer => ../../core/type/storer
)

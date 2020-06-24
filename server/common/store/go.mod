module gitlab.com/alienspaces/holyragingmages/server/common/store

go 1.13

require (
	github.com/jmoiron/sqlx v1.2.0
	github.com/lib/pq v1.2.0
	github.com/stretchr/testify v1.4.0
	gitlab.com/alienspaces/holyragingmages/server/common/config v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/common/log v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/common/type/configurer v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/common/type/logger v1.0.0
	google.golang.org/appengine v1.6.5 // indirect
)

replace (
	gitlab.com/alienspaces/holyragingmages/server/common/config => ../../common/config
	gitlab.com/alienspaces/holyragingmages/server/common/log => ../../common/log
	gitlab.com/alienspaces/holyragingmages/server/common/type/configurer => ../../common/type/configurer
	gitlab.com/alienspaces/holyragingmages/server/common/type/logger => ../../common/type/logger
)

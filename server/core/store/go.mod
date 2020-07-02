module gitlab.com/alienspaces/holyragingmages/server/core/store

go 1.13

require (
	github.com/jmoiron/sqlx v1.2.0
	github.com/lib/pq v1.2.0
	github.com/stretchr/testify v1.4.0
	gitlab.com/alienspaces/holyragingmages/server/core/config v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/core/log v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/core/type/configurer v1.0.0
	gitlab.com/alienspaces/holyragingmages/server/core/type/logger v1.0.0
	google.golang.org/appengine v1.6.5 // indirect
)

replace (
	gitlab.com/alienspaces/holyragingmages/server/core/config => ../../core/config
	gitlab.com/alienspaces/holyragingmages/server/core/log => ../../core/log
	gitlab.com/alienspaces/holyragingmages/server/core/type/configurer => ../../core/type/configurer
	gitlab.com/alienspaces/holyragingmages/server/core/type/logger => ../../core/type/logger
)

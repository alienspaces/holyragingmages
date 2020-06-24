module gitlab.com/alienspaces/holyragingmages/server/common/type/preparer

go 1.13

require (
	github.com/jmoiron/sqlx v1.2.0
	github.com/lib/pq v1.2.0 // indirect

	gitlab.com/alienspaces/holyragingmages/server/common/type/preparable v1.0.0
	google.golang.org/appengine v1.6.5 // indirect
)

replace gitlab.com/alienspaces/holyragingmages/server/common/type/preparable => ../../../common/type/preparable

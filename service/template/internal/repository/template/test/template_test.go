package test

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"gitlab.com/alienspaces/holyragingmages/common/type/logger"
	"gitlab.com/alienspaces/holyragingmages/common/type/preparer"
	"gitlab.com/alienspaces/holyragingmages/common/type/repositor"
	"gitlab.com/alienspaces/holyragingmages/service/template/internal/repository/template"
)

// NewRepositories - Custom repositories for this model
func newRepositories(l logger.Logger, p preparer.Preparer, tx *sqlx.Tx) ([]repositor.Repositor, error) {

	repositoryList := []repositor.Repositor{}

	tr, err := template.NewRepository(l, p, tx)
	if err != nil {
		return nil, err
	}

	repositoryList = append(repositoryList, tr)

	return repositoryList, nil
}

func TestCreateRec(t *testing.T) {

}

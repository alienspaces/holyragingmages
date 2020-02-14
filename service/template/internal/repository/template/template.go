package template

import (
	"github.com/jmoiron/sqlx"

	"gitlab.com/alienspaces/holyragingmages/common/repository"
	"gitlab.com/alienspaces/holyragingmages/common/type/logger"
	"gitlab.com/alienspaces/holyragingmages/common/type/preparer"
	"gitlab.com/alienspaces/holyragingmages/service/template/internal/record"
)

// Repository -
type Repository struct {
	repository.Repository
}

// NewRepository - returns an implementation of a Repo
func NewRepository(l logger.Logger, p preparer.Preparer, tx *sqlx.Tx) (*Repository, error) {

	r := &Repository{
		repository.Repository{
			Log:     l,
			Prepare: p,
			Tx:      tx,
		},
	}

	err := r.Init(p, tx)
	if err != nil {
		l.Warn("Failed new repository >%v<", err)
		return nil, err
	}

	return r, nil
}

// NewRecord -
func (r *Repository) NewRecord() *record.Template {
	return &record.Template{}
}

// NewRecordArray -
func (r *Repository) NewRecordArray() []*record.Template {
	return []*record.Template{}
}

// GetOne -
func (r *Repository) GetOne(id string, forUpdate bool) (*record.Template, error) {
	rec := r.NewRecord()
	if err := r.GetOneRec(id, rec); err != nil {
		r.Log.Warn("Failed statement execution >%v<", err)
		return nil, err
	}
	return rec, nil
}

// GetMany -
func (r *Repository) GetMany(
	params map[string]interface{},
	paramOperators map[string]string,
	forUpdate bool) ([]*record.Template, error) {

	recs := r.NewRecordArray()

	rows, err := r.GetManyRecs(params, paramOperators)
	if err != nil {
		r.Log.Warn("Failed statement execution >%v<", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		rec := r.NewRecord()
		err := rows.StructScan(rec)
		if err != nil {
			r.Log.Warn("Failed executing struct scan >%v<", err)
			return nil, err
		}
		recs = append(recs, rec)
	}

	r.Log.Warn("Fetched >%d< records", len(recs))

	return recs, nil
}

// CreateOne -
func (r *Repository) CreateOne(rec *record.Template) error {

	if rec.ID == "" {
		rec.ID = repository.NewRecordID()
	}
	rec.CreatedAt = repository.NewCreatedAt()

	err := r.CreateOneRec(rec)
	if err != nil {
		rec.CreatedAt = ""
		r.Log.Warn("Failed statement execution >%v<", err)
		return err
	}

	return nil
}

// UpdateOne -
func (r *Repository) UpdateOne(rec *record.Template) error {

	origUpdatedAt := rec.UpdatedAt
	rec.UpdatedAt = repository.NewUpdatedAt()

	err := r.UpdateOneRec(rec)
	if err != nil {
		rec.UpdatedAt = origUpdatedAt
		r.Log.Warn("Failed statement execution >%v<", err)
		return err
	}

	return nil
}

// CreateSQL -
func (r *Repository) CreateSQL() string {
	return `
INSERT INTO template
   (id, created_at)
VALUES
   (:id, :created_at)
RETURNING *
`
}

// UpdateOneSQL -
func (r *Repository) UpdateOneSQL() string {
	return `
UPDATE template SET
   updated_at = :updated_at
WHERE id 		   = :id
AND   deleted_at IS NULL
RETURNING *
`
}

// CreateTestRecord - creates a record for testing
func (r *Repository) CreateTestRecord(rec *record.Template) error {

	return r.CreateOne(rec)
}

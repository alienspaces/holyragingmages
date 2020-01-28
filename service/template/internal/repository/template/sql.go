package template

var getByIDSQL = `
SELECT *
FROM template
WHERE id = $1
AND deleted_at IS NULL
`

var getByParamSQL = `
SELECT *
FROM template
WHERE deleted_at IS NULL
`

var createRecordSQL = `
INSERT INTO template (
	id, client_ref, created_at
) VALUES (
	:id, :client_ref, :created_at
)
RETURNING *
`

var updateRecordSQL = `
UPDATE template SET
	client_ref = :client_ref,
	updated_at = :updated_at
WHERE id = :id
AND deleted_at IS NULL
RETURNING *
`

var deleteRecordSQL = `
UPDATE template SET
	deleted_at = :deleted_at
WHERE id = :id
AND deleted_at IS NULL
`

var removeRecordSQL = `
DELETE FROM template
WHERE id = :id
`

// GetByParamSQL -
func (r *Repository) GetByParamSQL() string {
	return getByParamSQL
}

// GetByIDSQL -
func (r *Repository) GetByIDSQL() string {
	return getByIDSQL
}

// GetCreateSQL -
func (r *Repository) GetCreateSQL() string {
	return createRecordSQL
}

// GetUpdateSQL -
func (r *Repository) GetUpdateSQL() string {
	return updateRecordSQL
}

// GetDeleteSQL -
func (r *Repository) GetDeleteSQL() string {
	return deleteRecordSQL
}

// GetRemoveSQL -
func (r *Repository) GetRemoveSQL() string {
	return removeRecordSQL
}

// OrderBy -
func (r *Repository) OrderBy() string {
	return "created_at desc"
}

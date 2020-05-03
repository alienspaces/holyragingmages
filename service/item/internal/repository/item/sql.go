package item

var createOneSQL = `
INSERT INTO item (
	id,
	name,
	description,
	created_at
) VALUES (
	:id,
	:name,
	:description,
	:created_at
)
RETURNING *
`

var updateOneSQL = `
UPDATE item SET
    name        = :name,
    description = :description,
    updated_at  = :updated_at
WHERE id = :id
AND   deleted_at IS NULL
RETURNING *
`

// CreateOneSQL -
func (r *Repository) CreateOneSQL() string {
	return createOneSQL
}

// UpdateOneSQL -
func (r *Repository) UpdateOneSQL() string {
	return updateOneSQL
}

// OrderBy -
func (r *Repository) OrderBy() string {
	return "created_at desc"
}

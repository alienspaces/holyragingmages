package accountentity

var createOneSQL = `
INSERT INTO account_entity (
	id,
	account_id,
	entity_id,
	created_at
) VALUES (
	:id,
	:account_id,
	:entity_id,
	:created_at
)
RETURNING *
`

var updateOneSQL = `
UPDATE account_entity SET
    id                = :id,
    account_id        = :account_id,
    entity_id         = :entity_id,
    updated_at        = :updated_at
WHERE id 		      = :id
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

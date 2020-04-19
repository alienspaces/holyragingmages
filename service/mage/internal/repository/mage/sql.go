package mage

var createOneSQL = `
INSERT INTO mage (
	id,
	name,
	strength,
	dexterity,
	intelligence,
	experience,
	coin,
	created_at
) VALUES (
	:id,
	:name,
	:strength,
	:dexterity,
	:intelligence,
	:experience,
	:coin,
	:created_at
)
RETURNING *
`

var updateOneSQL = `
UPDATE mage SET
    id           = :id,
    name         = :name,
    strength     = :strength,
    dexterity    = :dexterity,
    intelligence = :intelligence,
    experience   = :experience,
    coin         = :coin,
    updated_at   = :updated_at
WHERE id 		   = :id
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

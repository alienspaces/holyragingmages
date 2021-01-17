package entity

var createOneSQL = `
INSERT INTO entity (
	id,
	entity_type,
	name,
	avatar,
	strength,
	dexterity,
	intelligence,
	attribute_points,
	experience_points,
	coins,
	created_at
) VALUES (
	:id,
	:entity_type,
	:name,
	:avatar,
	:strength,
	:dexterity,
	:intelligence,
	:attribute_points,
	:experience_points,
	:coins,
	:created_at
)
RETURNING *
`

var updateOneSQL = `
UPDATE entity SET
	id                = :id,
	entity_type       = :entity_type,      
    name              = :name,
    avatar            = :avatar,
    strength          = :strength,
    dexterity         = :dexterity,
    intelligence      = :intelligence,
    attribute_points  = :attribute_points,
    experience_points = :experience_points,
    coins             = :coins,
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

package preparable

// Preparable -
type Preparable interface {
	TableName() string
	GetOneSQL() string
	GetManySQL() string
	CreateSQL() string
	UpdateOneSQL() string
	UpdateManySQL() string
	DeleteOneSQL() string
	DeleteManySQL() string
	RemoveOneSQL() string
	RemoveManySQL() string
}

package model

import (
	"database/sql"
	"fmt"

	"gitlab.com/alienspaces/holyragingmages/server/service/entity/internal/record"
)

// GetEntityRecs -
func (m *Model) GetEntityRecs(params map[string]interface{}, operators map[string]string, forUpdate bool) ([]*record.Entity, error) {

	m.Log.Info("Getting mage records params >%s<", params)

	r := m.EntityRepository()

	return r.GetMany(params, operators, forUpdate)
}

// GetEntityRec -
func (m *Model) GetEntityRec(recID string, forUpdate bool) (*record.Entity, error) {

	m.Log.Info("Getting mage rec ID >%s<", recID)

	r := m.EntityRepository()

	// validate UUID
	if m.IsUUID(recID) != true {
		return nil, fmt.Errorf("ID >%s< is not a valid UUID", recID)
	}

	rec, err := r.GetOne(recID, forUpdate)
	if err == sql.ErrNoRows {
		m.Log.Warn("No record found ID >%s<", recID)
		return nil, nil
	}

	return rec, err
}

// CreateEntityRec -
func (m *Model) CreateEntityRec(rec *record.Entity) error {

	m.Log.Info("Creating mage rec >%v<", rec)

	r := m.EntityRepository()

	// Defaults
	rec.AttributePoints = startingAttributePoints

	err := m.ValidateEntityRec(rec)
	if err != nil {
		m.Log.Info("Failed model validation >%v<", err)
		return err
	}

	return r.CreateOne(rec)
}

// UpdateEntityRec -
func (m *Model) UpdateEntityRec(rec *record.Entity) error {

	m.Log.Info("Updating mage rec >%v<", rec)

	r := m.EntityRepository()

	err := m.ValidateEntityRec(rec)
	if err != nil {
		m.Log.Info("Failed model validation >%v<", err)
		return err
	}

	return r.UpdateOne(rec)
}

// DeleteEntityRec -
func (m *Model) DeleteEntityRec(recID string) error {

	m.Log.Info("Deleting mage rec ID >%s<", recID)

	r := m.EntityRepository()

	// validate UUID
	if m.IsUUID(recID) != true {
		return fmt.Errorf("ID >%s< is not a valid UUID", recID)
	}

	err := m.ValidateDeleteEntityRec(recID)
	if err != nil {
		m.Log.Info("Failed model validation >%v<", err)
		return err
	}

	return r.DeleteOne(recID)
}

// RemoveEntityRec -
func (m *Model) RemoveEntityRec(recID string) error {

	m.Log.Info("Removing mage rec ID >%s<", recID)

	r := m.EntityRepository()

	// validate UUID
	if m.IsUUID(recID) != true {
		return fmt.Errorf("ID >%s< is not a valid UUID", recID)
	}

	err := m.ValidateDeleteEntityRec(recID)
	if err != nil {
		m.Log.Info("Failed model validation >%v<", err)
		return err
	}

	return r.RemoveOne(recID)
}

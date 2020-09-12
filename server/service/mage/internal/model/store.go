package model

import (
	"database/sql"
	"fmt"

	"gitlab.com/alienspaces/holyragingmages/server/service/mage/internal/record"
)

// GetMageRecs -
func (m *Model) GetMageRecs(params map[string]interface{}, operators map[string]string, forUpdate bool) ([]*record.Mage, error) {

	m.Log.Info("Getting mage records params >%s<", params)

	r := m.MageRepository()

	return r.GetMany(params, operators, forUpdate)
}

// GetMageRec -
func (m *Model) GetMageRec(recID string, forUpdate bool) (*record.Mage, error) {

	m.Log.Info("Getting mage rec ID >%s<", recID)

	r := m.MageRepository()

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

// CreateMageRec -
func (m *Model) CreateMageRec(rec *record.Mage) error {

	m.Log.Info("Creating mage rec >%v<", rec)

	r := m.MageRepository()

	// Defaults

	err := m.ValidateMageRec(rec)
	if err != nil {
		m.Log.Info("Failed model validation >%v<", err)
		return err
	}

	return r.CreateOne(rec)
}

// UpdateMageRec -
func (m *Model) UpdateMageRec(rec *record.Mage) error {

	m.Log.Info("Updating mage rec >%v<", rec)

	r := m.MageRepository()

	err := m.ValidateMageRec(rec)
	if err != nil {
		m.Log.Info("Failed model validation >%v<", err)
		return err
	}

	return r.UpdateOne(rec)
}

// DeleteMageRec -
func (m *Model) DeleteMageRec(recID string) error {

	m.Log.Info("Deleting mage rec ID >%s<", recID)

	r := m.MageRepository()

	// validate UUID
	if m.IsUUID(recID) != true {
		return fmt.Errorf("ID >%s< is not a valid UUID", recID)
	}

	err := m.ValidateDeleteMageRec(recID)
	if err != nil {
		m.Log.Info("Failed model validation >%v<", err)
		return err
	}

	return r.DeleteOne(recID)
}

// RemoveMageRec -
func (m *Model) RemoveMageRec(recID string) error {

	m.Log.Info("Removing mage rec ID >%s<", recID)

	r := m.MageRepository()

	// validate UUID
	if m.IsUUID(recID) != true {
		return fmt.Errorf("ID >%s< is not a valid UUID", recID)
	}

	err := m.ValidateDeleteMageRec(recID)
	if err != nil {
		m.Log.Info("Failed model validation >%v<", err)
		return err
	}

	return r.RemoveOne(recID)
}

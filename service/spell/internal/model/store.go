package model

import (
	"database/sql"
	"fmt"

	"gitlab.com/alienspaces/holyragingmages/service/spell/internal/record"
)

// GetSpellRecs -
func (m *Model) GetSpellRecs(params map[string]interface{}, operators map[string]string, forUpdate bool) ([]*record.Spell, error) {

	m.Log.Info("Getting spell records params >%s<", params)

	r := m.SpellRepository()

	return r.GetMany(params, operators, forUpdate)
}

// GetSpellRec -
func (m *Model) GetSpellRec(recID string, forUpdate bool) (*record.Spell, error) {

	m.Log.Info("Getting spell rec ID >%s<", recID)

	r := m.SpellRepository()

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

// CreateSpellRec -
func (m *Model) CreateSpellRec(rec *record.Spell) error {

	m.Log.Info("Creating spell rec >%v<", rec)

	r := m.SpellRepository()

	err := m.ValidateSpellRec(rec)
	if err != nil {
		m.Log.Info("Failed model validation >%v<", err)
		return err
	}

	return r.CreateOne(rec)
}

// UpdateSpellRec -
func (m *Model) UpdateSpellRec(rec *record.Spell) error {

	m.Log.Info("Updating spell rec >%v<", rec)

	r := m.SpellRepository()

	err := m.ValidateSpellRec(rec)
	if err != nil {
		m.Log.Info("Failed model validation >%v<", err)
		return err
	}

	return r.UpdateOne(rec)
}

// DeleteSpellRec -
func (m *Model) DeleteSpellRec(recID string) error {

	m.Log.Info("Deleting spell rec ID >%s<", recID)

	r := m.SpellRepository()

	// validate UUID
	if m.IsUUID(recID) != true {
		return fmt.Errorf("ID >%s< is not a valid UUID", recID)
	}

	err := m.ValidateDeleteSpellRec(recID)
	if err != nil {
		m.Log.Info("Failed model validation >%v<", err)
		return err
	}

	return r.DeleteOne(recID)
}

// RemoveSpellRec -
func (m *Model) RemoveSpellRec(recID string) error {

	m.Log.Info("Removing spell rec ID >%s<", recID)

	r := m.SpellRepository()

	// validate UUID
	if m.IsUUID(recID) != true {
		return fmt.Errorf("ID >%s< is not a valid UUID", recID)
	}

	err := m.ValidateDeleteSpellRec(recID)
	if err != nil {
		m.Log.Info("Failed model validation >%v<", err)
		return err
	}

	return r.RemoveOne(recID)
}

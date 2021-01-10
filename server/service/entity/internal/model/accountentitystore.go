package model

import (
	"database/sql"
	"fmt"

	"gitlab.com/alienspaces/holyragingmages/server/service/entity/internal/record"
)

// GetAccountEntityRecs -
func (m *Model) GetAccountEntityRecs(params map[string]interface{}, operators map[string]string, forUpdate bool) ([]*record.AccountEntity, error) {

	m.Log.Info("Getting account entity records params >%s<", params)

	r := m.AccountEntityRepository()

	return r.GetMany(params, operators, forUpdate)
}

// GetAccountEntityRec -
func (m *Model) GetAccountEntityRec(recID string, forUpdate bool) (*record.AccountEntity, error) {

	m.Log.Info("Getting account entity rec ID >%s<", recID)

	r := m.AccountEntityRepository()

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

// CreateAccountEntityRec -
func (m *Model) CreateAccountEntityRec(rec *record.AccountEntity) error {

	m.Log.Info("Creating account entity rec >%#v<", rec)

	r := m.AccountEntityRepository()

	err := m.ValidateAccountEntityRec(rec)
	if err != nil {
		m.Log.Info("Failed model validation >%v<", err)
		return err
	}

	return r.CreateOne(rec)
}

// UpdateAccountEntityRec -
func (m *Model) UpdateAccountEntityRec(rec *record.AccountEntity) error {

	m.Log.Info("Updating account entity rec >%#v<", rec)

	r := m.AccountEntityRepository()

	err := m.ValidateAccountEntityRec(rec)
	if err != nil {
		m.Log.Info("Failed model validation >%v<", err)
		return err
	}

	return r.UpdateOne(rec)
}

// DeleteAccountEntityRec -
func (m *Model) DeleteAccountEntityRec(recID string) error {

	m.Log.Info("Deleting account entity rec ID >%s<", recID)

	r := m.AccountEntityRepository()

	// validate UUID
	if m.IsUUID(recID) != true {
		return fmt.Errorf("ID >%s< is not a valid UUID", recID)
	}

	err := m.ValidateDeleteAccountEntityRec(recID)
	if err != nil {
		m.Log.Info("Failed model validation >%v<", err)
		return err
	}

	return r.DeleteOne(recID)
}

// RemoveAccountEntityRec -
func (m *Model) RemoveAccountEntityRec(recID string) error {

	m.Log.Info("Removing account entity rec ID >%s<", recID)

	r := m.AccountEntityRepository()

	// validate UUID
	if m.IsUUID(recID) != true {
		return fmt.Errorf("ID >%s< is not a valid UUID", recID)
	}

	err := m.ValidateDeleteAccountEntityRec(recID)
	if err != nil {
		m.Log.Info("Failed model validation >%v<", err)
		return err
	}

	return r.RemoveOne(recID)
}

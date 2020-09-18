package model

import (
	"database/sql"
	"fmt"

	"gitlab.com/alienspaces/holyragingmages/server/service/account/internal/record"
)

// GetAccountRecs -
func (m *Model) GetAccountRecs(params map[string]interface{}, operators map[string]string, forUpdate bool) ([]*record.Account, error) {

	m.Log.Info("Getting account records params >%s<", params)

	r := m.AccountRepository()

	return r.GetMany(params, operators, forUpdate)
}

// GetAccountRec -
func (m *Model) GetAccountRec(recID string, forUpdate bool) (*record.Account, error) {

	m.Log.Info("Getting account rec ID >%s<", recID)

	r := m.AccountRepository()

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

// CreateAccountRec -
func (m *Model) CreateAccountRec(rec *record.Account) error {

	m.Log.Info("Creating account rec >%v<", rec)

	r := m.AccountRepository()

	err := m.ValidateAccountRec(rec)
	if err != nil {
		m.Log.Info("Failed model validation >%v<", err)
		return err
	}

	return r.CreateOne(rec)
}

// UpdateAccountRec -
func (m *Model) UpdateAccountRec(rec *record.Account) error {

	m.Log.Info("Updating account rec >%v<", rec)

	r := m.AccountRepository()

	err := m.ValidateAccountRec(rec)
	if err != nil {
		m.Log.Info("Failed model validation >%v<", err)
		return err
	}

	return r.UpdateOne(rec)
}

// DeleteAccountRec -
func (m *Model) DeleteAccountRec(recID string) error {

	m.Log.Info("Deleting account rec ID >%s<", recID)

	r := m.AccountRepository()

	// validate UUID
	if m.IsUUID(recID) != true {
		return fmt.Errorf("ID >%s< is not a valid UUID", recID)
	}

	err := m.ValidateDeleteAccountRec(recID)
	if err != nil {
		m.Log.Info("Failed model validation >%v<", err)
		return err
	}

	return r.DeleteOne(recID)
}

// RemoveAccountRec -
func (m *Model) RemoveAccountRec(recID string) error {

	m.Log.Info("Removing account rec ID >%s<", recID)

	r := m.AccountRepository()

	// validate UUID
	if m.IsUUID(recID) != true {
		return fmt.Errorf("ID >%s< is not a valid UUID", recID)
	}

	err := m.ValidateDeleteAccountRec(recID)
	if err != nil {
		m.Log.Info("Failed model validation >%v<", err)
		return err
	}

	return r.RemoveOne(recID)
}

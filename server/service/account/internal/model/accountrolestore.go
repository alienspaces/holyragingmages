package model

import (
	"database/sql"
	"fmt"

	"gitlab.com/alienspaces/holyragingmages/server/service/account/internal/record"
)

// GetAccountRoleRecs -
func (m *Model) GetAccountRoleRecs(params map[string]interface{}, operators map[string]string, forUpdate bool) ([]*record.AccountRole, error) {

	m.Log.Info("Getting account records params >%s<", params)

	r := m.AccountRoleRepository()

	return r.GetMany(params, operators, forUpdate)
}

// GetAccountRoleRec -
func (m *Model) GetAccountRoleRec(recID string, forUpdate bool) (*record.AccountRole, error) {

	m.Log.Info("Getting account rec ID >%s<", recID)

	r := m.AccountRoleRepository()

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

// CreateAccountRoleRec -
func (m *Model) CreateAccountRoleRec(rec *record.AccountRole) error {

	m.Log.Info("Creating account rec >%v<", rec)

	r := m.AccountRoleRepository()

	err := m.ValidateAccountRoleRec(rec)
	if err != nil {
		m.Log.Info("Failed model validation >%v<", err)
		return err
	}

	return r.CreateOne(rec)
}

// UpdateAccountRoleRec -
func (m *Model) UpdateAccountRoleRec(rec *record.AccountRole) error {

	m.Log.Info("Updating account rec >%v<", rec)

	r := m.AccountRoleRepository()

	err := m.ValidateAccountRoleRec(rec)
	if err != nil {
		m.Log.Info("Failed model validation >%v<", err)
		return err
	}

	return r.UpdateOne(rec)
}

// DeleteAccountRoleRec -
func (m *Model) DeleteAccountRoleRec(recID string) error {

	m.Log.Info("Deleting account rec ID >%s<", recID)

	r := m.AccountRoleRepository()

	// validate UUID
	if m.IsUUID(recID) != true {
		return fmt.Errorf("ID >%s< is not a valid UUID", recID)
	}

	err := m.ValidateDeleteAccountRoleRec(recID)
	if err != nil {
		m.Log.Info("Failed model validation >%v<", err)
		return err
	}

	return r.DeleteOne(recID)
}

// RemoveAccountRoleRec -
func (m *Model) RemoveAccountRoleRec(recID string) error {

	m.Log.Info("Removing account rec ID >%s<", recID)

	r := m.AccountRoleRepository()

	// validate UUID
	if m.IsUUID(recID) != true {
		return fmt.Errorf("ID >%s< is not a valid UUID", recID)
	}

	err := m.ValidateDeleteAccountRoleRec(recID)
	if err != nil {
		m.Log.Info("Failed model validation >%v<", err)
		return err
	}

	return r.RemoveOne(recID)
}

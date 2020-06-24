package model

import (
	"database/sql"
	"fmt"

	"gitlab.com/alienspaces/holyragingmages/server/service/item/internal/record"
)

// GetItemRecs -
func (m *Model) GetItemRecs(params map[string]interface{}, operators map[string]string, forUpdate bool) ([]*record.Item, error) {

	m.Log.Info("Getting item records params >%s<", params)

	r := m.ItemRepository()

	return r.GetMany(params, operators, forUpdate)
}

// GetItemRec -
func (m *Model) GetItemRec(recID string, forUpdate bool) (*record.Item, error) {

	m.Log.Info("Getting item rec ID >%s<", recID)

	r := m.ItemRepository()

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

// CreateItemRec -
func (m *Model) CreateItemRec(rec *record.Item) error {

	m.Log.Info("Creating item rec >%v<", rec)

	r := m.ItemRepository()

	err := m.ValidateItemRec(rec)
	if err != nil {
		m.Log.Info("Failed model validation >%v<", err)
		return err
	}

	return r.CreateOne(rec)
}

// UpdateItemRec -
func (m *Model) UpdateItemRec(rec *record.Item) error {

	m.Log.Info("Updating item rec >%v<", rec)

	r := m.ItemRepository()

	err := m.ValidateItemRec(rec)
	if err != nil {
		m.Log.Info("Failed model validation >%v<", err)
		return err
	}

	return r.UpdateOne(rec)
}

// DeleteItemRec -
func (m *Model) DeleteItemRec(recID string) error {

	m.Log.Info("Deleting item rec ID >%s<", recID)

	r := m.ItemRepository()

	// validate UUID
	if m.IsUUID(recID) != true {
		return fmt.Errorf("ID >%s< is not a valid UUID", recID)
	}

	err := m.ValidateDeleteItemRec(recID)
	if err != nil {
		m.Log.Info("Failed model validation >%v<", err)
		return err
	}

	return r.DeleteOne(recID)
}

// RemoveItemRec -
func (m *Model) RemoveItemRec(recID string) error {

	m.Log.Info("Removing item rec ID >%s<", recID)

	r := m.ItemRepository()

	// validate UUID
	if m.IsUUID(recID) != true {
		return fmt.Errorf("ID >%s< is not a valid UUID", recID)
	}

	err := m.ValidateDeleteItemRec(recID)
	if err != nil {
		m.Log.Info("Failed model validation >%v<", err)
		return err
	}

	return r.RemoveOne(recID)
}

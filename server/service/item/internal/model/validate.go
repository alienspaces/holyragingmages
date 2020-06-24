package model

import (
	"gitlab.com/alienspaces/holyragingmages/server/service/item/internal/record"
)

// ValidateItemRec - validates creating and updating a item record
func (m *Model) ValidateItemRec(rec *record.Item) error {

	return nil
}

// ValidateDeleteItemRec - validates it is okay to delete a item record
func (m *Model) ValidateDeleteItemRec(recID string) error {

	return nil
}

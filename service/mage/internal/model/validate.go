package model

import (
	"fmt"

	"gitlab.com/alienspaces/holyragingmages/service/mage/internal/record"
)

// ValidateMageRec - validates creating and updating a mage record
func (m *Model) ValidateMageRec(rec *record.Mage) error {

	// required fields
	if rec.Name == "" {
		return fmt.Errorf("Name is required")
	}

	return nil
}

// ValidateDeleteMageRec - validates it is okay to delete a mage record
func (m *Model) ValidateDeleteMageRec(recID string) error {

	return nil
}

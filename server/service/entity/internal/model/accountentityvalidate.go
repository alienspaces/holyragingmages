package model

import (
	"fmt"

	"gitlab.com/alienspaces/holyragingmages/server/service/entity/internal/record"
)

// ValidateAccountEntityRec - validates creating and updating a mage record
func (m *Model) ValidateAccountEntityRec(rec *record.AccountEntity) error {

	// required fields
	if rec.AccountID == "" {
		return fmt.Errorf("AccountID is required")
	}
	if rec.EntityID == "" {
		return fmt.Errorf("EntityID is required")
	}

	return nil
}

// ValidateDeleteAccountEntityRec - validates it is okay to delete a mage record
func (m *Model) ValidateDeleteAccountEntityRec(recID string) error {

	return nil
}

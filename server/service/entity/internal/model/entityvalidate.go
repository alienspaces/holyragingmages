package model

import (
	"fmt"

	"gitlab.com/alienspaces/holyragingmages/server/service/entity/internal/record"
)

// ValidateEntityRec - validates creating and updating a mage record
func (m *Model) ValidateEntityRec(rec *record.Entity) error {

	if rec.ID == "" {
		if rec.AttributePoints != startingAttributePoints {
			//
		}
	}

	// required fields
	if rec.EntityType == "" {
		return fmt.Errorf("EntityType is required")
	}
	if rec.Name == "" {
		return fmt.Errorf("Name is required")
	}
	if rec.Avatar == "" {
		return fmt.Errorf("Avatar is required")
	}

	return nil
}

// ValidateDeleteEntityRec - validates it is okay to delete a mage record
func (m *Model) ValidateDeleteEntityRec(recID string) error {

	return nil
}

package model

import (
	"gitlab.com/alienspaces/holyragingmages/server/service/spell/internal/record"
)

// ValidateSpellRec - validates creating and updating a spell record
func (m *Model) ValidateSpellRec(rec *record.Spell) error {

	return nil
}

// ValidateDeleteSpellRec - validates it is okay to delete a spell record
func (m *Model) ValidateDeleteSpellRec(recID string) error {

	return nil
}

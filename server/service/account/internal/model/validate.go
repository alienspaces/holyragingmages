package model

import (
	"gitlab.com/alienspaces/holyragingmages/server/service/account/internal/record"
)

// ValidateAccountRec - validates creating and updating a account record
func (m *Model) ValidateAccountRec(rec *record.Account) error {

	return nil
}

// ValidateDeleteAccountRec - validates it is okay to delete a account record
func (m *Model) ValidateDeleteAccountRec(recID string) error {

	return nil
}

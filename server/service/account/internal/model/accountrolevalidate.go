package model

import (
	"gitlab.com/alienspaces/holyragingmages/server/service/account/internal/record"
)

// ValidateAccountRoleRec - validates creating and updating a account record
func (m *Model) ValidateAccountRoleRec(rec *record.AccountRole) error {

	return nil
}

// ValidateDeleteAccountRoleRec - validates it is okay to delete a account record
func (m *Model) ValidateDeleteAccountRoleRec(recID string) error {

	return nil
}

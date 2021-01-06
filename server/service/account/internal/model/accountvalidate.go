package model

import (
	"fmt"

	"gitlab.com/alienspaces/holyragingmages/server/service/account/internal/record"
)

// ValidateAccountRec - validates creating and updating an account record
func (m *Model) ValidateAccountRec(rec *record.Account) error {

	switch rec.Provider {
	case record.AccountProviderAnonymous:
		// We only require a provider account ID to create an anonymous local account
		if rec.ProviderAccountID == "" {
			msg := "Missing ProviderAccountID, cannot create an anonymous account"
			m.Log.Warn(msg)
			return fmt.Errorf(msg)
		}
	case record.AccountProviderGoogle:
		// We require a provider account ID and email to create a Google local account
		if rec.Email == "" {
			msg := "Missing Email, cannot create a Google account"
			m.Log.Warn(msg)
			return fmt.Errorf(msg)
		}
		if rec.ProviderAccountID == "" {
			msg := "Missing ProviderAccountID, cannot create a Google account"
			m.Log.Warn(msg)
			return fmt.Errorf(msg)
		}
	default:
		// no-op
	}
	return nil
}

// ValidateDeleteAccountRec - validates it is okay to delete a account record
func (m *Model) ValidateDeleteAccountRec(recID string) error {

	return nil
}

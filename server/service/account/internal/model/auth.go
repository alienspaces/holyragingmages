package model

import (
	"fmt"
	"net/http"

	"gitlab.com/alienspaces/holyragingmages/server/service/account/internal/record"
	"google.golang.org/api/oauth2/v1"
)

// AuthData - encapsulates data provided by an authorizer
type AuthData struct {
	Provider          string
	ProviderAccountID string
	ProviderToken     string
	AccountEmail      string
	AccountName       string
}

// TODO: Return JWT instead of account record..

// VerifyProviderToken - verifies an authentication token from a provider and returns a local account record
func (m *Model) VerifyProviderToken(data AuthData) (*record.Account, error) {

	var verifiedAccountID string
	var verifiedAccountEmail string
	var verifiedAccountName string
	var rec *record.Account

	switch data.Provider {
	case record.AccountProviderGoogle:
		tokenInfo, err := m.verifyGoogleToken(data.ProviderToken)
		if err != nil {
			m.Log.Warn("Failed verifyGoogleToken >%v<", err)
			return nil, err
		}

		m.Log.Info("Token info UserId>%s<", tokenInfo.UserId)
		m.Log.Info("Token info Email >%s<", tokenInfo.Email)

		if data.ProviderAccountID == tokenInfo.UserId {
			verifiedAccountID = tokenInfo.UserId
			verifiedAccountEmail = tokenInfo.Email
			// Google token verification does not return an account name
			// so we'll use the account name provided by the client
			verifiedAccountName = data.AccountName
		}

	default:
		// Unsupported
		return nil, fmt.Errorf("Unsupported provider >%s<", data.Provider)
	}

	if verifiedAccountID == "" {
		m.Log.Warn("Failed verifying account")
		return nil, fmt.Errorf("Failed verifying account")
	}

	// Fetch account based on provider and provider account ID
	accountRecs, err := m.GetAccountRecs(
		map[string]interface{}{
			"provider":            data.Provider,
			"provider_account_id": verifiedAccountID,
		},
		nil,
		false,
	)
	if err != nil {
		m.Log.Warn("Failed getting user account >%v<", err)
		return nil, err
	}

	if len(accountRecs) > 1 {
		msg := fmt.Sprintf("Found more than expected account records >%d<", len(accountRecs))
		m.Log.Warn(msg)
		return nil, fmt.Errorf(msg)
	}

	// Use account when a single record is found
	if len(accountRecs) == 1 {
		rec = accountRecs[0]
	}

	// Create account when no records are found
	if len(accountRecs) == 0 {
		m.Log.Info("Failed getting user account, records >%d<", len(accountRecs))
		rec = &record.Account{
			Name:              verifiedAccountName,
			Email:             verifiedAccountEmail,
			Provider:          data.Provider,
			ProviderAccountID: verifiedAccountID,
		}
		err := m.CreateAccountRec(rec)
		if err != nil {
			m.Log.Warn("Failed creating account >%v<", err)
			return nil, err
		}
	}

	return rec, nil
}

// Provider verification methods
func (m *Model) verifyGoogleToken(token string) (*oauth2.Tokeninfo, error) {

	var httpClient = &http.Client{}
	oauth2Service, err := oauth2.New(httpClient)
	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfoCall.IdToken(token)
	tokenInfo, err := tokenInfoCall.Do()
	if err != nil {
		return nil, err
	}
	return tokenInfo, nil
}

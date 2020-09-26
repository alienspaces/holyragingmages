package model

import (
	"context"
	"fmt"

	"google.golang.org/api/oauth2/v1"
	"google.golang.org/api/option"

	"gitlab.com/alienspaces/holyragingmages/server/service/account/internal/record"
)

// AuthData - encapsulates data provided by an authorizer
type AuthData struct {
	Provider          string
	ProviderAccountID string
	ProviderToken     string
	AccountEmail      string
	AccountName       string
}

// VerifiedData -
type VerifiedData struct {
	UserID string
	Email  string
}

// VerifyProviderToken - verifies an authentication token from a provider and returns a local account record
func (m *Model) VerifyProviderToken(data AuthData) (*record.Account, error) {

	var verifiedAccountID string
	var verifiedAccountEmail string
	var verifiedAccountName string
	var rec *record.Account

	// Check required
	if data.Provider == "" {
		msg := "AuthData missing Provider, cannot VerifyProviderToken"
		m.Log.Warn(msg)
		return nil, fmt.Errorf(msg)
	}

	if data.ProviderToken == "" {
		msg := "AuthData missing ProviderToken, cannot VerifyProviderToken"
		m.Log.Warn(msg)
		return nil, fmt.Errorf(msg)
	}

	m.Log.Info("Verifying provider >%s<", data.Provider)
	m.Log.Info("Verifying account  >%s<", data.ProviderAccountID)
	m.Log.Info("Verifying token    >%s<", data.ProviderToken)

	switch data.Provider {
	case record.AccountProviderGoogle:
		verifiedData, err := m.VerifyAuthTokenFunc(record.AccountProviderGoogle, data.ProviderToken)
		if err != nil {
			m.Log.Warn("Failed VerifyAuthTokenFunc >%v<", err)
			return nil, err
		}
		if verifiedData == nil {
			msg := "Failed VerifyAuthTokenFunc, verified data is nil"
			m.Log.Warn(msg)
			return nil, fmt.Errorf(msg)
		}

		m.Log.Info("Token info UserId>%s<", verifiedData.UserID)
		m.Log.Info("Token info Email >%s<", verifiedData.Email)

		if data.ProviderAccountID == verifiedData.UserID {
			verifiedAccountID = verifiedData.UserID
			verifiedAccountEmail = verifiedData.Email
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
func (m *Model) verifyAuthToken(provider, token string) (*VerifiedData, error) {

	verifiedData := &VerifiedData{}

	if provider == record.AccountProviderGoogle {

		// API key
		apiKey := m.Config.Get("APP_SERVER_GOOGLE_API_KEY")

		m.Log.Info("Google API key >%s<", apiKey)

		ctx := context.Background()
		oauth2Service, err := oauth2.NewService(ctx, option.WithAPIKey(apiKey))
		if err != nil {
			m.Log.Warn("Failed new Google oauth2 service >%v<", err)
			return nil, err
		}

		tokenInfoCall := oauth2Service.Tokeninfo()
		tokenInfoCall.AccessToken(token)
		tokenInfo, err := tokenInfoCall.Do()
		if err != nil {
			m.Log.Warn("Google oauth2 token info call >%v<", err)
			return nil, err
		}
		verifiedData.UserID = tokenInfo.UserId
		verifiedData.Email = tokenInfo.Email
	}

	return verifiedData, nil
}

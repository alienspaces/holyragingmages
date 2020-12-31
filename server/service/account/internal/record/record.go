package record

import (
	"gitlab.com/alienspaces/holyragingmages/server/core/repository"
)

// AccountProvider - Valid account providers
const (
	AccountProviderAnonymous string = "anonymous"
	AccountProviderApple     string = "apple"
	AccountProviderFacebook  string = "facebook"
	AccountProviderGithub    string = "github"
	AccountProviderGoogle    string = "google"
	AccountProviderTwitter   string = "twitter"
)

// Account -
type Account struct {
	repository.Record
	Name              string `db:"name"`
	Email             string `db:"email"`
	Provider          string `db:"provider"`
	ProviderAccountID string `db:"provider_account_id"`
}

// AccountRole - Valid roles
const (
	AccountRoleDefault       string = "default"
	AccountRoleAdministrator string = "administrator"
)

// AccountRole -
type AccountRole struct {
	repository.Record
	AccountID string `db:"account_id"`
	Role      string `db:"role"`
}

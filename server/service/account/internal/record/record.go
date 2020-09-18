package record

import (
	"gitlab.com/alienspaces/holyragingmages/server/core/repository"
)

// AccountProvider -
const (
	AccountProviderGoogle   string = "google"
	AccountProviderApple    string = "apple"
	AccountProviderFacebook string = "facebook"
	AccountProviderTwitter  string = "twitter"
	AccountProviderGithub   string = "github"
)

// Account -
type Account struct {
	repository.Record
	Name              string `db:"name"`
	Email             string `db:"email"`
	Provider          string `db:"provider"`
	ProviderAccountID string `db:"provider_account_id"`
}

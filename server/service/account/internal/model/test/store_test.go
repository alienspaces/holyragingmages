package test

// NOTE: model tests are run is the public space so we are
// able to use common setup and teardown tooling for all models

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"gitlab.com/alienspaces/holyragingmages/server/service/account/internal/harness"
	"gitlab.com/alienspaces/holyragingmages/server/service/account/internal/model"
	"gitlab.com/alienspaces/holyragingmages/server/service/account/internal/record"
)

func TestCreateAccountRec(t *testing.T) {

	// harness
	config := harness.DataConfig{}

	h, err := harness.NewTesting(config)
	require.NoError(t, err, "NewTesting returns without error")

	// harness commit data
	h.CommitData = true

	tests := []struct {
		name string
		rec  func() *record.Account
		err  bool
	}{
		{
			name: "Without ID",
			rec: func() *record.Account {
				return &record.Account{
					Name:              "Scary Susan",
					Email:             "scarysusan@example.com",
					Provider:          record.AccountProviderGoogle,
					ProviderAccountID: "abcdefg",
				}
			},
			err: false,
		},
		{
			name: "With ID",
			rec: func() *record.Account {
				rec := &record.Account{
					Name:              "Horrific Harry",
					Email:             "horrificharry@example.com",
					Provider:          record.AccountProviderGoogle,
					ProviderAccountID: "abcdefg",
				}
				id, _ := uuid.NewRandom()
				rec.ID = id.String()
				return rec
			},
			err: false,
		},
	}

	for _, tc := range tests {

		t.Logf("Run test >%s<", tc.name)

		func() {

			// test harness
			err = h.Setup()
			require.NoError(t, err, "Setup returns without error")
			defer func() {
				err = h.RollbackTx()
				require.NoError(t, err, "RollbackTx returns without error")
				err = h.Teardown()
				require.NoError(t, err, "Teardown returns without error")
			}()

			// init tx
			err = h.InitTx(nil)
			require.NoError(t, err, "InitTx returns without error")

			rec := tc.rec()

			err = h.Model.(*model.Model).CreateAccountRec(rec)
			if tc.err == true {
				require.Error(t, err, "CreateAccountRec returns error")
				return
			}
			require.NoError(t, err, "CreateAccountRec returns without error")
			require.NotEmpty(t, rec.CreatedAt, "CreateAccountRec returns record with CreatedAt")
		}()
	}
}

func TestGetAccountRec(t *testing.T) {

	// harness
	config := harness.DataConfig{
		AccountConfig: []harness.AccountConfig{
			{
				Record: record.Account{},
			},
		},
	}

	h, err := harness.NewTesting(config)
	require.NoError(t, err, "NewTesting returns without error")

	// harness commit data
	h.CommitData = true

	tests := []struct {
		name string
		id   func() string
		err  bool
	}{
		{
			name: "With ID",
			id: func() string {
				return h.Data.AccountRecs[0].ID
			},
			err: false,
		},
		{
			name: "Without ID",
			id: func() string {
				return ""
			},
			err: true,
		},
	}

	for _, tc := range tests {

		t.Logf("Run test >%s<", tc.name)

		func() {

			// harness setup
			err = h.Setup()
			require.NoError(t, err, "Setup returns without error")
			defer func() {
				err = h.RollbackTx()
				require.NoError(t, err, "RollbackTx returns without error")
				err = h.Teardown()
				require.NoError(t, err, "Teardown returns without error")
			}()

			// init tx
			err = h.InitTx(nil)
			require.NoError(t, err, "InitTx returns without error")

			rec, err := h.Model.(*model.Model).GetAccountRec(tc.id(), false)
			if tc.err == true {
				require.Error(t, err, "GetAccountRec returns error")
				return
			}
			require.NoError(t, err, "GetAccountRec returns without error")
			require.NotNil(t, rec, "GetAccountRec returns record")
			require.NotEmpty(t, rec.ID, "Record ID is not empty")
		}()
	}
}

func TestUpdateAccountRec(t *testing.T) {

	// harness
	config := harness.DataConfig{
		AccountConfig: []harness.AccountConfig{
			{
				Record: record.Account{},
			},
		},
	}

	h, err := harness.NewTesting(config)

	// harness commit data
	h.CommitData = true

	require.NoError(t, err, "NewTesting returns without error")

	tests := []struct {
		name string
		rec  func() record.Account
		err  bool
	}{
		{
			name: "With ID",
			rec: func() record.Account {
				return h.Data.AccountRecs[0]
			},
			err: false,
		},
		{
			name: "Without ID",
			rec: func() record.Account {
				rec := h.Data.AccountRecs[0]
				rec.ID = ""
				return rec
			},
			err: true,
		},
	}

	for _, tc := range tests {

		t.Logf("Run test >%s<", tc.name)

		func() {

			// harness setup
			err = h.Setup()
			require.NoError(t, err, "Setup returns without error")
			defer func() {
				err = h.RollbackTx()
				require.NoError(t, err, "RollbackTx returns without error")
				err = h.Teardown()
				require.NoError(t, err, "Teardown returns without error")
			}()

			// init tx
			err = h.InitTx(nil)
			require.NoError(t, err, "InitTx returns without error")

			rec := tc.rec()

			err := h.Model.(*model.Model).UpdateAccountRec(&rec)
			if tc.err == true {
				require.Error(t, err, "UpdateAccountRec returns error")
				return
			}
			require.NoError(t, err, "UpdateAccountRec returns without error")
			require.NotEmpty(t, rec.UpdatedAt, "UpdateAccountRec returns record with UpdatedAt")
		}()
	}
}

func TestDeleteAccountRec(t *testing.T) {

	// harness
	config := harness.DataConfig{
		AccountConfig: []harness.AccountConfig{
			{
				Record: record.Account{},
			},
		},
	}

	h, err := harness.NewTesting(config)
	require.NoError(t, err, "NewTesting returns without error")

	// harness commit data
	h.CommitData = true

	tests := []struct {
		name string
		id   func() string
		err  bool
	}{
		{
			name: "With ID",
			id: func() string {
				return h.Data.AccountRecs[0].ID
			},
			err: false,
		},
		{
			name: "Without ID",
			id: func() string {
				return ""
			},
			err: true,
		},
	}

	for _, tc := range tests {

		t.Logf("Run test >%s<", tc.name)

		func() {

			// harness setup
			err = h.Setup()
			require.NoError(t, err, "Setup returns without error")
			defer func() {
				err = h.RollbackTx()
				require.NoError(t, err, "RollbackTx returns without error")
				err = h.Teardown()
				require.NoError(t, err, "Teardown returns without error")
			}()

			// init tx
			err = h.InitTx(nil)
			require.NoError(t, err, "InitTx returns without error")

			err := h.Model.(*model.Model).DeleteAccountRec(tc.id())
			if tc.err == true {
				require.Error(t, err, "DeleteAccountRec returns error")
				return
			}
			require.NoError(t, err, "DeleteAccountRec returns without error")

			rec, err := h.Model.(*model.Model).GetAccountRec(tc.id(), false)
			require.NoError(t, err, "GetAccountRec returns without error")
			require.Nil(t, rec, "GetAccountRec does not return record")
		}()
	}
}

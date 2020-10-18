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

func TestCreateAccountRoleRec(t *testing.T) {

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
		rec  func(data *harness.Data) *record.AccountRole
		err  bool
	}{
		{
			name: "Without ID",
			rec: func(data *harness.Data) *record.AccountRole {
				return &record.AccountRole{
					AccountID: data.AccountRecs[0].ID,
					Role:      record.AccountRoleDefault,
				}
			},
			err: false,
		},
		{
			name: "With ID",
			rec: func(data *harness.Data) *record.AccountRole {
				rec := &record.AccountRole{
					AccountID: data.AccountRecs[0].ID,
					Role:      record.AccountRoleDefault,
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

			// Test harness
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

			rec := tc.rec(h.Data)

			err = h.Model.(*model.Model).CreateAccountRoleRec(rec)
			if tc.err == true {
				require.Error(t, err, "CreateAccountRoleRec returns error")
				return
			}
			require.NoError(t, err, "CreateAccountRoleRec returns without error")
			require.NotEmpty(t, rec.CreatedAt, "CreateAccountRoleRec returns record with CreatedAt")
		}()
	}
}

func TestGetAccountRoleRec(t *testing.T) {

	// harness
	config := harness.DataConfig{
		AccountConfig: []harness.AccountConfig{
			{
				Record: record.Account{},
				AccountRoleConfig: []harness.AccountRoleConfig{
					{
						Record: record.AccountRole{},
					},
				},
			},
		},
	}

	h, err := harness.NewTesting(config)
	require.NoError(t, err, "NewTesting returns without error")

	// harness commit data
	h.CommitData = true

	tests := []struct {
		name string
		id   func(data *harness.Data) string
		err  bool
	}{
		{
			name: "With ID",
			id: func(data *harness.Data) string {
				return data.AccountRoleRecs[0].ID
			},
			err: false,
		},
		{
			name: "Without ID",
			id: func(data *harness.Data) string {
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

			rec, err := h.Model.(*model.Model).GetAccountRoleRec(tc.id(h.Data), false)
			if tc.err == true {
				require.Error(t, err, "GetAccountRoleRec returns error")
				return
			}
			require.NoError(t, err, "GetAccountRoleRec returns without error")
			require.NotNil(t, rec, "GetAccountRoleRec returns record")
			require.NotEmpty(t, rec.ID, "Record ID is not empty")
		}()
	}
}

func TestUpdateAccountRoleRec(t *testing.T) {

	// harness
	config := harness.DataConfig{
		AccountConfig: []harness.AccountConfig{
			{
				Record: record.Account{},
				AccountRoleConfig: []harness.AccountRoleConfig{
					{
						Record: record.AccountRole{},
					},
				},
			},
		},
	}

	h, err := harness.NewTesting(config)

	// harness commit data
	h.CommitData = true

	require.NoError(t, err, "NewTesting returns without error")

	tests := []struct {
		name string
		rec  func(data *harness.Data) record.AccountRole
		err  bool
	}{
		{
			name: "With ID",
			rec: func(data *harness.Data) record.AccountRole {
				return data.AccountRoleRecs[0]
			},
			err: false,
		},
		{
			name: "Without ID",
			rec: func(data *harness.Data) record.AccountRole {
				rec := data.AccountRoleRecs[0]
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

			rec := tc.rec(h.Data)

			err := h.Model.(*model.Model).UpdateAccountRoleRec(&rec)
			if tc.err == true {
				require.Error(t, err, "UpdateAccountRoleRec returns error")
				return
			}
			require.NoError(t, err, "UpdateAccountRoleRec returns without error")
			require.NotEmpty(t, rec.UpdatedAt, "UpdateAccountRoleRec returns record with UpdatedAt")
		}()
	}
}

func TestDeleteAccountRoleRec(t *testing.T) {

	// harness
	config := harness.DataConfig{
		AccountConfig: []harness.AccountConfig{
			{
				Record: record.Account{},
				AccountRoleConfig: []harness.AccountRoleConfig{
					{
						Record: record.AccountRole{},
					},
				},
			},
		},
	}

	h, err := harness.NewTesting(config)
	require.NoError(t, err, "NewTesting returns without error")

	// harness commit data
	h.CommitData = true

	tests := []struct {
		name string
		id   func(data *harness.Data) string
		err  bool
	}{
		{
			name: "With ID",
			id: func(data *harness.Data) string {
				return data.AccountRoleRecs[0].ID
			},
			err: false,
		},
		{
			name: "Without ID",
			id: func(data *harness.Data) string {
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

			err := h.Model.(*model.Model).DeleteAccountRoleRec(tc.id(h.Data))
			if tc.err == true {
				require.Error(t, err, "DeleteAccountRoleRec returns error")
				return
			}
			require.NoError(t, err, "DeleteAccountRoleRec returns without error")

			rec, err := h.Model.(*model.Model).GetAccountRoleRec(tc.id(h.Data), false)
			require.NoError(t, err, "GetAccountRoleRec returns without error")
			require.Nil(t, rec, "GetAccountRoleRec does not return record")
		}()
	}
}

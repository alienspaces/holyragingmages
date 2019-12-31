package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEnv(t *testing.T) {

	tests := map[string]struct {
		dotEnv  bool
		items   []Item
		wantErr bool
	}{
		"NewEnv with items": {
			dotEnv: false,
			items: []Item{
				Item{
					Key:      "HOME",
					Required: true,
				},
			},
			wantErr: false,
		},
		"NewEnv without items": {
			dotEnv:  false,
			items:   nil,
			wantErr: false,
		},
		"NewEnv without dot env": {
			dotEnv:  true,
			items:   nil,
			wantErr: true,
		},
	}

	for tcName, tc := range tests {

		t.Logf("Running test >%s<", tcName)

		e, err := NewEnv(tc.items, tc.dotEnv)
		if tc.wantErr {
			if assert.Error(t, err, "NewEnv returns with error") {
				continue
			}
		}
		if assert.NoError(t, err, "NewEnv returns without error") {
			assert.NotNil(t, e, "NewEnv returns environment object")
		}
	}
}

func TestGet(t *testing.T) {

	tests := map[string]struct {
		items      []Item
		wantErr    bool
		wantValues []string
	}{
		"Get valid environment value": {
			items: []Item{
				Item{
					Key:      "HOME",
					Required: true,
				},
			},
			wantErr: false,
			wantValues: []string{
				os.Getenv("HOME"),
			},
		},
		"Get invalid environment value": {
			items: []Item{
				Item{
					Key:      "WORK",
					Required: true,
				},
			},
			wantErr:    true,
			wantValues: []string{},
		}}

	for tcName, tc := range tests {

		t.Logf("Running test >%s<", tcName)

		e, err := NewEnv(tc.items, false)
		if tc.wantErr {
			if assert.Error(t, err, "NewEnv returns with error") {
				continue
			}
		}
		for idx, item := range tc.items {
			value := e.Get(item.Key)
			assert.Equal(t, tc.wantValues[idx], value, "Get returns expected value")
		}
	}
}

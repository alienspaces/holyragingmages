package template

import (
	"gitlab.com/alienspaces/holyragingmages/service/template/internal/record"
)

// ValidateTemplateRec - validates creating and updating a template record
func (m *Model) ValidateTemplateRec(rec *record.Template) error {

	return nil
}

// ValidateDeleteTemplateRec - validates it is okay to delete a template record
func (m *Model) ValidateDeleteTemplateRec(recID string) error {

	return nil
}

package model

import (
	"gitlab.com/alienspaces/holyragingmages/service/template/internal/record"
)

// GetTemplateRecs -
func (m *Model) GetTemplateRecs(params map[string]interface{}) ([]*record.Template, error) {
	return nil, nil
}

// GetTemplateRec -
func (m *Model) GetTemplateRec(recID string, forUpdate bool) (*record.Template, error) {
	return nil, nil
}

// CreateTemplateRec -
func (m *Model) CreateTemplateRec(rec *record.Template) error {

	m.Log.Info("Creating template rec >%v<", rec)

	r := m.TemplateRepository()

	return r.CreateOne(rec)
}

// UpdateTemplateRec -
func (m *Model) UpdateTemplateRec(rec *record.Template) error {
	return nil
}

// DeleteTemplateRec -
func (m *Model) DeleteTemplateRec(recID string) error {
	return nil
}

package service

// APIRunner -
type APIRunner struct {
	Store  Storer
	Log    Logger
	Config Configurer
}

// Init -
func (r *APIRunner) Init(c Configurer, l Logger, s Storer) error {

	r.Config = c
	r.Log = l
	r.Store = s

	return nil
}

// Run -
func (r *APIRunner) Run(args map[string]interface{}) error {

	return nil
}

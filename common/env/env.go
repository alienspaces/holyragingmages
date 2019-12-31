// Package env provides methods for managing environment variables
package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Item defines a valid environment variable and whether it is required
type Item struct {
	Key      string
	Required bool
}

// Env defines a container of items and corresponding values
type Env struct {
	Items  []*Item
	Values map[string]string
}

// NewEnv creates a new environment object
func NewEnv(items []Item, dotEnv bool) (*Env, error) {

	e := Env{
		Values: make(map[string]string),
	}

	err := e.Init(items, dotEnv)
	if err != nil {
		return nil, fmt.Errorf("NewEnv failed init >%v<", err)
	}

	return &e, nil
}

// Init initialises and checks environment values
func (e *Env) Init(items []Item, dotEnv bool) (err error) {

	dir := os.Getenv("APP_HOME")
	if dir == "" {
		dir, err = os.Getwd()
	}

	if dotEnv {
		envFile := fmt.Sprintf("%s/%s", dir, ".env")
		err = godotenv.Load(envFile)
		if err != nil {
			return err
		}
	}

	for _, item := range items {
		err = e.Add(item)
		if err != nil {
			return err
		}
	}

	return nil
}

// Get returns an environment item value
func (e *Env) Get(key string) (value string) {

	for _, item := range e.Items {
		if item.Key == key {
			value = e.Values[key]
			return value
		}
	}

	return ""
}

// Set an environment item value
func (e *Env) Set(key string, value string) {

	for _, item := range e.Items {
		if item.Key == key {
			e.Values[key] = value
			return
		}
	}
}

// Add will add a new environment item
func (e *Env) Add(item Item) (err error) {

	e.Items = append(e.Items, &item)

	err = e.sourceItem(&item)
	if err != nil {
		return err
	}

	err = e.checkItem(&item)
	if err != nil {
		return err
	}

	return nil
}

// Verify checks whether the provided items have values set
func (e *Env) Verify(items []Item) (err error) {

	for _, item := range items {
		err = e.checkItem(&item)
		if err != nil {
			return err
		}
	}

	return nil
}

// sourceItem - sources and sets an environment item value
func (e *Env) sourceItem(item *Item) error {

	value := os.Getenv(item.Key)
	e.Set(item.Key, value)

	return nil
}

// checkItem - checks an environment item
func (e *Env) checkItem(item *Item) error {

	if item.Required && e.Values[item.Key] == "" {
		return fmt.Errorf("Failed checking env value >%s<", item.Key)
	}

	return nil
}

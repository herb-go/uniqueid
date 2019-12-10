package uniqueid

import (
	"encoding/json"
)

//Option generator option interface.
type Option interface {
	ApplyTo(*Generator) error
}

// OptionConfig option config in map format.
type OptionConfig struct {
	Driver string
	Config func(v interface{}) error `config:", lazyload"`
}

//ApplyTo apply option to file store.
func (o *OptionConfig) ApplyTo(g *Generator) error {
	driver, err := NewDriver(o.Driver, o.Config)
	if err != nil {
		return err
	}
	g.Driver = driver
	return nil
}

//NewOptionConfig create new option config.
func NewOptionConfig() *OptionConfig {
	return &OptionConfig{}
}

var LoadConfig = func(c map[string]interface{}, key string, v interface{}) error {
	i, ok := c[key]
	if !ok {
		return nil
	}
	bs, err := json.Marshal(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(bs, v)
}

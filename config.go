package uniqueid

import (
	"encoding/json"
)

//Option generator option interface.
type Option interface {
	ApplyTo(*Generator) error
}

// OptionConfigMap option config in map format.
type OptionConfigMap struct {
	Driver string
	Config map[string]interface{}
}

//ApplyTo apply option to file store.
func (o *OptionConfigMap) ApplyTo(g *Generator) error {
	driver, err := NewDriver(o.Driver, o.Config, "")
	if err != nil {
		return err
	}
	g.Driver = driver
	return nil
}

//NewOptionConfigMap create new option config.
func NewOptionConfigMap() *OptionConfigMap {
	return &OptionConfigMap{
		Config: map[string]interface{}{},
	}
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

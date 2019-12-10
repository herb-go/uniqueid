package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"github.com/herb-go/uniqueid"
)

//SnowFlake snow flake driver
type SnowFlake struct {
	node *snowflake.Node
}

//NewSnowFlake create new snow flake driver
func NewSnowFlake() *SnowFlake {
	return &SnowFlake{}
}

//GenerateID generate unique id.
//Return  generated id and any error if rasied.
func (s *SnowFlake) GenerateID() (string, error) {
	return s.node.Generate().String(), nil
}

type SnowFlakeConfig struct {
	Node int64
}

//Factory snow flake driver factory
func Factory(loader func(v interface{}) error) (uniqueid.Driver, error) {
	var err error
	s := NewSnowFlake()
	conf := &SnowFlakeConfig{}
	err = loader(conf)
	if err != nil {
		return nil, err
	}
	s.node, err = snowflake.NewNode(conf.Node)
	if err != nil {
		return nil, err
	}
	return s, nil

}
func init() {
	uniqueid.Register("snowflake", Factory)
}

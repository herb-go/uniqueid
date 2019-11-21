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

//Factory snow flake driver factory
func Factory(conf map[string]interface{}, prefix string) (uniqueid.Driver, error) {
	var err error
	s := NewSnowFlake()
	var node int64
	uniqueid.LoadConfig(conf, prefix+"Node", &node)
	s.node, err = snowflake.NewNode(node)
	if err != nil {
		return nil, err
	}
	return s, nil

}
func init() {
	uniqueid.Register("snowflake", Factory)
}

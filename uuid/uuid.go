package uuid

import (
	"github.com/herb-go/uniqueid"
	uuid "github.com/satori/go.uuid"
)

//UUID uuid driver
type UUID struct {
	creator func() (uuid.UUID, error)
}

//NewUUID create new uuid driver
func NewUUID() *UUID {
	return &UUID{}
}

//V1 generate unique id by uuid version1.
//Return  generated id and any error if rasied.
func V1() (string, error) {
	uid, err := uuid.NewV1()
	if err != nil {
		return "", err
	}
	return uid.String(), nil
}

//GenerateID generate unique id.
//Return  generated id and any error if rasied.
func (u *UUID) GenerateID() (string, error) {
	uid, err := u.creator()
	if err != nil {
		return "", err
	}
	return uid.String(), nil
}

type UUIDConfig struct {
	Version int
}

//Factory uuid driver factory
func Factory(loader func(v interface{}) error) (uniqueid.Driver, error) {
	i := NewUUID()
	conf := &UUIDConfig{}

	err := loader(conf)
	if err != nil {
		return nil, err
	}
	switch conf.Version {
	case 4:
		i.creator = uuid.NewV4
	default:
		i.creator = uuid.NewV1
	}
	return i, nil

}

func init() {
	uniqueid.Register("uuid", Factory)
}
